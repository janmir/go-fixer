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

//Fixer an instance of fixer api
type Fixer struct {
}

func init() {
	pwd, err := util.GetCurrDir()
	util.Catch(err)

	dbFile := filepath.Join(pwd, _offDB)
	util.Logger("DB file: ", dbFile)

	//database: online->dynamoDB, offline->bolt
	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	util.Catch(err)

	defer db.Close()
}
