package fixer

import (
	"testing"

	util "github.com/janmir/go-util"
)

func TestFetch(t *testing.T) {
	f := Make()
	f.Fetch(Currencies.PHP, Currencies.USD)
	f.Close()
}

func TestConvert(t *testing.T) {
	f := Make()
	defer f.Close()

	valString := ""
	f.Convert(Currencies.PHP, Currencies.USD, &valString)
	util.Logger("PHP->USD:", valString)

	valString = ""
	f.Convert(Currencies.PHP, Currencies.JPY, &valString)
	util.Logger("PHP->JPY:", valString)

	valString = ""
	f.Convert(Currencies.JPY, Currencies.PHP, &valString)
	util.Logger("JPY->PHP:", valString)

	valFloat := float32(0.0)
	f.Convert(Currencies.JPY, Currencies.USD, &valFloat)
	util.Logger("JPY->USD:", valFloat)
}
