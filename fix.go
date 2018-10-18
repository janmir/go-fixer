package fixer

import (
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
func (f Fixer) Convert(from, to Currency, out *interface{}) {

}

//Fetch makes a Get request to the source
func (f Fixer) Fetch() {

}

//Close close the things that are open
func (f Fixer) Close() {
	switch {
	case !_offline:
	default:
		f.db.(*bolt.DB).Close()
	}
}
