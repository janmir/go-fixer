package main

import (
	fixer "github.com/janmir/go-fixer"
	util "github.com/janmir/go-util"
)

func main() {
	f := fixer.Make()
	defer f.Close()

	valString := ""
	f.Convert(fixer.Currencies.JPY, fixer.Currencies.PHP, &valString)
	util.Logger("JPY->PHP:", valString)
}
