package fixer

//Source defines datasources
type Source struct {
	ty, url string
}

var (
	_offline = true
	_version = "0.0.1"

	_sources = []Source{
		{"xml", "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"},
		{"api", "http://free.currencyconverterapi.com/api/v5/convert?q=JPY_PHP&compact=y"},
	}
)
