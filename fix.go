package fixer

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

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
	db interface{}
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
	fix := Fixer{}
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
//
func (f Fixer) Convert(from, to Currency, out interface{}) error {
	defer util.TimeTrack(time.Now(), "Conversion")

	//check if output variable is a pointer
	util.IsInterfaceAPointer(out)

	//check currency availability
	if !(from.Ok && to.Ok) {
		util.Catch(fmt.Errorf("Currency conversion from ｢%s｣ to ｢%s｣ not yet supported", from.Acr, to.Acr))
	}

	//get conversions
	forex, err := f.Fetch(from, to)
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
		*(out.(*string)) = fmt.Sprintf("%s %0.5f", forex.Sym, forex.exc)
	default:
		util.Catch(fmt.Errorf("Type unknown/unsupported ｢%T｣", out))
	}

	return nil
}

//Fetch makes a Get request to the source
func (f Fixer) Fetch(from, to Currency) (Currency, error) {
	var (
		errr  error
		forex = Currency{}
	)

	//check db first
	// key: date
	// value: {curr:"", acr:"", sym: ""}

	//get new data
	for _, v := range _sources {
		switch v.typ {
		case "xml":
			//Get the xml
			xmlB := f.Get(v.url)

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
				forex = to
				forex.exc = xmlD.Calculate(from, to)

				return forex, nil
			default:
			}
		case "api":
			//Get via api
			url := fmt.Sprintf(v.url, from.Acr, to.Acr)

			apiB := f.Get(url)
			_ = apiB

			//Handle API
			forex = to
			forex.exc = 0.0

			return forex, nil
		default:
			errr = errors.New("Unable to perform conversion, all sources failed")
		}
	}

	return forex, errr
}

//getXMLType return type of xml data
func (f Fixer) getXMLType([]byte) interface{} {
	return EuroCenterBankRootXML{}
}

//Get makes a get request to the server
func (f Fixer) Get(url string) []byte {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	util.Catch(err)

	res, err := client.Do(req)
	util.HTTPCatch(res, err)

	//get data body
	bbb, err := ioutil.ReadAll(res.Body)
	util.Catch(err)
	res.Body.Close()

	return bbb
}

//Close close the things that are open
func (f Fixer) Close() {
	switch {
	case !_offline:
	default:
		f.db.(*bolt.DB).Close()
	}
}
