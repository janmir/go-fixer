package fixer

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/beevik/ntp"
	"github.com/boltdb/bolt"
	"github.com/janmir/go-util"
)

const (
	_offDB = "data.db"
)

var (
	pwd string
)

//Fixer an instance of fixer api
type Fixer struct {
	local time.Time
	db    interface{} //*bolt.DB
}

func init() {
	//Database Storage:
	// online->dynamoDB
	// offline->bolt
	switch {
	case !_offline:
	default:
		var err error
		pwd, err = util.GetCurrDir()
		util.Catch(err)
	}
}

//Make creates a new instance of Fixer
func Make() Fixer {
	var (
		local time.Time
		err   error
	)
	local, err = ntp.Time("time.apple.com")
	if err != nil {
		local = time.Now().Local()
		util.Green("Using localtime time(fallback)")
	} else {
		util.Green("Using network time")
	}

	fix := Fixer{
		local: local,
	}

	switch {
	case !_offline:
		fix.db = nil
	default:
		dbFile := filepath.Join(pwd, _offDB)
		util.Logger("DB file: ", dbFile)

		db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
		util.Catch(err)

		//Attach
		fix.db = db
	}

	return fix
}

//Convert converts the value to a different currency
func (f Fixer) Convert(from, to Currency, out interface{}, opt ...interface{}) error {
	defer util.TimeTrack(time.Now(), "Conversion")

	//check if output variable is a pointer
	util.IsInterfaceAPointer(out)

	//check currency availability
	if !(from.Ok && to.Ok) {
		util.Catch(fmt.Errorf("Currency conversion from ｢%s｣ to ｢%s｣ not yet supported", from.Acr, to.Acr))
	}

	//get conversions
	cache := true
	if len(opt) > 0 {
		cache = opt[0].(bool)
	}
	forex, err := f.Fetch(from, to, cache)
	if err != nil {
		return err
	}

	switch out.(type) {
	case *int:
		*(out.(*int)) = int(forex.exc)
	case *float64:
		*(out.(*float64)) = float64(forex.exc)
	case *float32:
		*(out.(*float32)) = float32(forex.exc)
	case *string:
		*(out.(*string)) = fmt.Sprintf("%s%0.4f", forex.Sym, forex.exc)
	default:
		util.Catch(fmt.Errorf("Type unknown/unsupported ｢%T｣", out))
	}

	return nil
}

//Fetch makes a Get request to the source
func (f Fixer) Fetch(from, to Currency, cache bool) (Currency, error) {
	var (
		errr  error
		forex = to
	)

	//check db first
	// key: FROM_TO, e.g PHP_JPY
	// value: "0.12121"
	yesterday := f.local.AddDate(0, 0, -1)
	bucketKey := yesterday.Format("2006-01-02")
	util.Logger("Local:", bucketKey)

	// store some data
	fromTo := fmt.Sprintf("%s_%s", from.Acr, to.Acr)
	found := false
	value := 0.0
	err := f.db.(*bolt.DB).View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketKey))
		if bucket != nil {
			val := bucket.Get([]byte(fromTo))
			if val != nil {
				fl, err := strconv.ParseFloat(string(val), 32)
				util.Catch(err)

				value = fl
				found = true
			}
		}
		return nil
	})
	util.Catch(err)
	util.Logger("Found:", found, value)

	if found {
		util.Magenta("Using cached data")
		forex.exc = float32(value)
		return forex, nil
	}

	//get new data
loopy:
	for _, v := range _sources {
		switch v.typ {
		case "xml":
			util.Cyan("Via XML Fetch")

			//Get the xml
			xmlB, err := f.Get(v.url)
			if err != nil {
				util.Red(fmt.Sprintf("%s, unable to get xml data.", err.Error()))
				continue
			}

			typee := f.getXMLType(xmlB)

			switch typee.(type) {
			case EuroCenterBankRootXML:
				xmlD := EuroCenterBankRootXML{}
				err := xml.Unmarshal(xmlB, &xmlD)
				if err != nil {
					util.Red(fmt.Sprintf("%s, unable to unmarshal xml data.", err.Error()))
					continue
				}

				//Handle EuroBack data here
				forex.exc = xmlD.Calculate(from, to)
				break loopy
			default:
			}
		case "api":
			util.Cyan("Via API Fetch")

			//Get via api
			url := fmt.Sprintf(v.url, from.Acr, to.Acr)

			apiB, err := f.Get(url)
			if err != nil {
				util.Red(fmt.Sprintf("%s, unable to get api data.", err.Error()))
				continue
			}

			var raw map[string]*json.RawMessage
			err = json.Unmarshal(apiB, &raw)
			if err != nil {
				util.Red(fmt.Sprintf("%s, unable to unmarshal api data.", err.Error()))
				continue
			}

			apiB = []byte(*raw[fromTo])
			apiD := CurrencyConverterAPI{}
			err = json.Unmarshal(apiB, &apiD)
			if err != nil {
				util.Red(fmt.Sprintf("%s, unable to unmarshal api data.", err.Error()))
				continue
			}

			//Handle API
			forex.exc = apiD.Val
			break loopy
		default:
			errr = errors.New("Unable to perform conversion, all sources failed")
		}
	}

	//Cache the dates result in database;
	if errr == nil && cache {
		util.Magenta("Storing cached xml/api data")
		err = f.db.(*bolt.DB).Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucketKey))
			util.Catch(err)

			err = bucket.Put([]byte(fromTo), []byte(fmt.Sprintf("%f", forex.exc)))
			if err != nil {
				return err
			}

			return nil
		})
		util.Catch(err)
	}

	return forex, errr
}

//getXMLType return type of xml data
func (f Fixer) getXMLType([]byte) interface{} {
	return EuroCenterBankRootXML{}
}

//Get makes a get request to the server
func (f Fixer) Get(url string) ([]byte, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := client.Do(req)
	if err == nil && res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code was %d", res.StatusCode)
	}
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	//get data body
	bbb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return bbb, nil
}

//Close close the things that are open
func (f Fixer) Close() {
	switch {
	case !_offline:
	default:
		f.db.(*bolt.DB).Close()
	}
}
