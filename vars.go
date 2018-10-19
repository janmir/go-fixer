package fixer

import (
	"errors"

	util "github.com/janmir/go-util"
)

//Currency structure for a type of currency
type Currency struct {
	Acr string  //currency acronyms i.e JPY, USD
	Sym string  //symbols for the currency e.g $, ¥, ₡, ₱, ₩ etc.
	Ok  bool    //supported or not
	exc float32 //the exchange value
}

//Source defines datasources
type Source struct {
	typ, url string
}

/******************************
	European Bank XML Data
*******************************/

//CubeParent data
type CubeParent struct {
	Cube CubeTime `xml:"Cube"`
}

//CubeTime data
type CubeTime struct {
	Time string `xml:"time,attr"`
	Cube []Cube `xml:"Cube"`
}

//Cube data
type Cube struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

//EuroCenterBankRootXML Europen Bank XMl data structure
type EuroCenterBankRootXML struct {
	Subject string `xml:"subject"`
	Sender  struct {
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube CubeParent `xml:"Cube"`
}

/******************************
	Currency Converter API
*******************************/

//CurrencyConverterAPI ...
type CurrencyConverterAPI struct {
	Val float32 `json:"val"`
}

var (
	_offline = true
	_version = "0.0.1"

	_sources = []Source{
		{"xml", "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"},
		{"api", "http://free.currencyconverterapi.com/api/v5/convert?q=%s_%s&compact=y"},
	}

	//Currencies list of all supported currencies
	Currencies = struct {
		ALL Currency //Albania Lek
		AFN Currency //Afghanistan Afghani
		ARS Currency //Argentina Peso
		AWG Currency //Aruba Guilder
		AUD Currency //Australia Dollar
		AZN Currency //Azerbaijan Manat
		BSD Currency //Bahamas Dollar
		BBD Currency //Barbados Dollar
		BYN Currency //Belarus Ruble
		BZD Currency //Belize Dollar
		BMD Currency //Bermuda Dollar
		BOB Currency //Bolivia Bolíviano
		BAM Currency //Bosnia and Herzegovina Convertible Marka
		BWP Currency //Botswana Pula
		BGN Currency //Bulgaria Lev
		BRL Currency //Brazil Real
		BND Currency //Brunei Darussalam Dollar
		KHR Currency //Cambodia Riel
		CAD Currency //Canada Dollar
		KYD Currency //Cayman Islands Dollar
		CLP Currency //Chile Peso
		CNY Currency //China Yuan Renminbi
		COP Currency //Colombia Peso
		CRC Currency //Costa Rica Colon
		HRK Currency //Croatia Kuna
		CUP Currency //Cuba Peso
		CZK Currency //Czech Republic Koruna
		DKK Currency //Denmark Krone
		DOP Currency //Dominican Republic Peso
		XCD Currency //East Caribbean Dollar
		EGP Currency //Egypt Pound
		SVC Currency //El Salvador Colon
		EUR Currency //Euro Member Countries
		FKP Currency //Falkland Islands (Malvinas) Pound
		FJD Currency //Fiji Dollar
		GHS Currency //Ghana Cedi
		GIP Currency //Gibraltar Pound
		GTQ Currency //Guatemala Quetzal
		GGP Currency //Guernsey Pound
		GYD Currency //Guyana Dollar
		HNL Currency //Honduras Lempira
		HKD Currency //Hong Kong Dollar
		HUF Currency //Hungary Forint
		ISK Currency //Iceland Krona
		INR Currency //India Rupee
		IDR Currency //Indonesia Rupiah
		IRR Currency //Iran Rial
		IMP Currency //Isle of Man Pound
		ILS Currency //Israel Shekel
		JMD Currency //Jamaica Dollar
		JPY Currency //Japan Yen
		JEP Currency //Jersey Pound
		KZT Currency //Kazakhstan Tenge
		KPW Currency //Korea (North) Won
		KRW Currency //Korea (South) Won
		KGS Currency //Kyrgyzstan Som
		LAK Currency //Laos Kip
		LBP Currency //Lebanon Pound
		LRD Currency //Liberia Dollar
		MKD Currency //Macedonia Denar
		MYR Currency //Malaysia Ringgit
		MUR Currency //Mauritius Rupee
		MXN Currency //Mexico Peso
		MNT Currency //Mongolia Tughrik
		MZN Currency //Mozambique Metical
		NAD Currency //Namibia Dollar
		NPR Currency //Nepal Rupee
		ANG Currency //Netherlands Antilles Guilder
		NZD Currency //New Zealand Dollar
		NIO Currency //Nicaragua Cordoba
		NGN Currency //Nigeria Naira
		NOK Currency //Norway Krone
		OMR Currency //Oman Rial
		PKR Currency //Pakistan Rupee
		PAB Currency //Panama Balboa
		PYG Currency //Paraguay Guarani
		PEN Currency //Peru Sol
		PHP Currency //Philippines Piso
		PLN Currency //Poland Zloty
		QAR Currency //Qatar Riyal
		RON Currency //Romania Leu
		RUB Currency //Russia Ruble
		SHP Currency //Saint Helena Pound
		SAR Currency //Saudi Arabia Riyal
		RSD Currency //Serbia Dinar
		SCR Currency //Seychelles Rupee
		SGD Currency //Singapore Dollar
		SBD Currency //Solomon Islands Dollar
		SOS Currency //Somalia Shilling
		ZAR Currency //South Africa Rand
		LKR Currency //Sri Lanka Rupee
		SEK Currency //Sweden Krona
		CHF Currency //Switzerland Franc
		SRD Currency //Suriname Dollar
		SYP Currency //Syria Pound
		TWD Currency //Taiwan New Dollar
		THB Currency //Thailand Baht
		TTD Currency //Trinidad and Tobago Dollar
		TRY Currency //Turkey Lira
		TVD Currency //Tuvalu Dollar
		UAH Currency //Ukraine Hryvnia
		GBP Currency //United Kingdom Pound
		USD Currency //United States Dollar
		UYU Currency //Uruguay Peso
		UZS Currency //Uzbekistan Som
		VEF Currency //Venezuela Bolívar
		VND Currency //Viet Nam Dong
		YER Currency //Yemen Rial
		ZWD Currency //Zimbabwe Dollar
	}{
		ALL: Currency{"ALL", "Lek", false, 0.0},  //Albania Lek
		AFN: Currency{"AFN", "؋", false, 0.0},    //Afghanistan Afghani
		ARS: Currency{"ARS", "$", false, 0.0},    //Argentina Peso
		AWG: Currency{"AWG", "ƒ", false, 0.0},    //Aruba Guilder
		AUD: Currency{"AUD", "$", true, 0.0},     //Australia Dollar
		AZN: Currency{"AZN", "₼", false, 0.0},    //Azerbaijan Manat
		BSD: Currency{"BSD", "$", false, 0.0},    //Bahamas Dollar
		BBD: Currency{"BBD", "$", false, 0.0},    //Barbados Dollar
		BYN: Currency{"BYN", "Br", false, 0.0},   //Belarus Ruble
		BZD: Currency{"BZD", "BZ$", false, 0.0},  //Belize Dollar
		BMD: Currency{"BMD", "$", false, 0.0},    //Bermuda Dollar
		BOB: Currency{"BOB", "$b", false, 0.0},   //Bolivia Bolíviano
		BAM: Currency{"BAM", "KM", false, 0.0},   //Bosnia and Herzegovina Convertible Marka
		BWP: Currency{"BWP", "P", false, 0.0},    //Botswana Pula
		BGN: Currency{"BGN", "лв", true, 0.0},    //Bulgaria Lev
		BRL: Currency{"BRL", "R$", true, 0.0},    //Brazil Real
		BND: Currency{"BND", "$", false, 0.0},    //Brunei Darussalam Dollar
		KHR: Currency{"KHR", "៛", false, 0.0},    //Cambodia Riel
		CAD: Currency{"CAD", "$", true, 0.0},     //Canada Dollar
		KYD: Currency{"KYD", "$", false, 0.0},    //Cayman Islands Dollar
		CLP: Currency{"CLP", "$", false, 0.0},    //Chile Peso
		CNY: Currency{"CNY", "¥", true, 0.0},     //China Yuan Renminbi
		COP: Currency{"COP", "$", false, 0.0},    //Colombia Peso
		CRC: Currency{"CRC", "₡", false, 0.0},    //Costa Rica Colon
		HRK: Currency{"HRK", "kn", true, 0.0},    //Croatia Kuna
		CUP: Currency{"CUP", "₱", false, 0.0},    //Cuba Peso
		CZK: Currency{"CZK", "Kč", true, 0.0},    //Czech Republic Koruna
		DKK: Currency{"DKK", "kr", true, 0.0},    //Denmark Krone
		DOP: Currency{"DOP", "RD$", false, 0.0},  //Dominican Republic Peso
		XCD: Currency{"XCD", "$", false, 0.0},    //East Caribbean Dollar
		EGP: Currency{"EGP", "£", false, 0.0},    //Egypt Pound
		SVC: Currency{"SVC", "$", false, 0.0},    //El Salvador Colon
		EUR: Currency{"EUR", "€", true, 0.0},     //Euro Member Countries
		FKP: Currency{"FKP", "£", false, 0.0},    //Falkland Islands (Malvinas) Pound
		FJD: Currency{"FJD", "$", false, 0.0},    //Fiji Dollar
		GHS: Currency{"GHS", "¢", false, 0.0},    //Ghana Cedi
		GIP: Currency{"GIP", "£", false, 0.0},    //Gibraltar Pound
		GTQ: Currency{"GTQ", "Q", false, 0.0},    //Guatemala Quetzal
		GGP: Currency{"GGP", "£", false, 0.0},    //Guernsey Pound
		GYD: Currency{"GYD", "$", false, 0.0},    //Guyana Dollar
		HNL: Currency{"HNL", "L", false, 0.0},    //Honduras Lempira
		HKD: Currency{"HKD", "$", true, 0.0},     //Hong Kong Dollar
		HUF: Currency{"HUF", "Ft", true, 0.0},    //Hungary Forint
		ISK: Currency{"ISK", "kr", true, 0.0},    //Iceland Krona
		INR: Currency{"INR", "", true, 0.0},      //India Rupee
		IDR: Currency{"IDR", "Rp", true, 0.0},    //Indonesia Rupiah
		IRR: Currency{"IRR", "﷼", false, 0.0},    //Iran Rial
		IMP: Currency{"IMP", "£", false, 0.0},    //Isle of Man Pound
		ILS: Currency{"ILS", "₪", true, 0.0},     //Israel Shekel
		JMD: Currency{"JMD", "J$", false, 0.0},   //Jamaica Dollar
		JPY: Currency{"JPY", "¥", true, 0.0},     //Japan Yen
		JEP: Currency{"JEP", "£", false, 0.0},    //Jersey Pound
		KZT: Currency{"KZT", "лв", false, 0.0},   //Kazakhstan Tenge
		KPW: Currency{"KPW", "₩", false, 0.0},    //Korea (North) Won
		KRW: Currency{"KRW", "₩", true, 0.0},     //Korea (South) Won
		KGS: Currency{"KGS", "лв", false, 0.0},   //Kyrgyzstan Som
		LAK: Currency{"LAK", "₭", false, 0.0},    //Laos Kip
		LBP: Currency{"LBP", "£", false, 0.0},    //Lebanon Pound
		LRD: Currency{"LRD", "$", false, 0.0},    //Liberia Dollar
		MKD: Currency{"MKD", "ден", false, 0.0},  //Macedonia Denar
		MYR: Currency{"MYR", "RM", true, 0.0},    //Malaysia Ringgit
		MUR: Currency{"MUR", "₨", false, 0.0},    //Mauritius Rupee
		MXN: Currency{"MXN", "$", true, 0.0},     //Mexico Peso
		MNT: Currency{"MNT", "₮", false, 0.0},    //Mongolia Tughrik
		MZN: Currency{"MZN", "MT", false, 0.0},   //Mozambique Metical
		NAD: Currency{"NAD", "$", false, 0.0},    //Namibia Dollar
		NPR: Currency{"NPR", "₨", false, 0.0},    //Nepal Rupee
		ANG: Currency{"ANG", "ƒ", false, 0.0},    //Netherlands Antilles Guilder
		NZD: Currency{"NZD", "$", true, 0.0},     //New Zealand Dollar
		NIO: Currency{"NIO", "C$", false, 0.0},   //Nicaragua Cordoba
		NGN: Currency{"NGN", "₦", false, 0.0},    //Nigeria Naira
		NOK: Currency{"NOK", "kr", true, 0.0},    //Norway Krone
		OMR: Currency{"OMR", "﷼", false, 0.0},    //Oman Rial
		PKR: Currency{"PKR", "₨", false, 0.0},    //Pakistan Rupee
		PAB: Currency{"PAB", "B/.", false, 0.0},  //Panama Balboa
		PYG: Currency{"PYG", "Gs", false, 0.0},   //Paraguay Guarani
		PEN: Currency{"PEN", "S/.", false, 0.0},  //Peru Sol
		PHP: Currency{"PHP", "₱", true, 0.0},     //Philippines Piso
		PLN: Currency{"PLN", "zł", true, 0.0},    //Poland Zloty
		QAR: Currency{"QAR", "﷼", false, 0.0},    //Qatar Riyal
		RON: Currency{"RON", "lei", true, 0.0},   //Romania Leu
		RUB: Currency{"RUB", "₽", true, 0.0},     //Russia Ruble
		SHP: Currency{"SHP", "£", false, 0.0},    //Saint Helena Pound
		SAR: Currency{"SAR", "﷼", false, 0.0},    //Saudi Arabia Riyal
		RSD: Currency{"RSD", "Дин.", false, 0.0}, //Serbia Dinar
		SCR: Currency{"SCR", "₨", false, 0.0},    //Seychelles Rupee
		SGD: Currency{"SGD", "$", true, 0.0},     //Singapore Dollar
		SBD: Currency{"SBD", "$", false, 0.0},    //Solomon Islands Dollar
		SOS: Currency{"SOS", "S", false, 0.0},    //Somalia Shilling
		ZAR: Currency{"ZAR", "R", true, 0.0},     //South Africa Rand
		LKR: Currency{"LKR", "₨", false, 0.0},    //Sri Lanka Rupee
		SEK: Currency{"SEK", "kr", true, 0.0},    //Sweden Krona
		CHF: Currency{"CHF", "CHF", true, 0.0},   //Switzerland Franc
		SRD: Currency{"SRD", "$", false, 0.0},    //Suriname Dollar
		SYP: Currency{"SYP", "£", false, 0.0},    //Syria Pound
		TWD: Currency{"TWD", "NT$", false, 0.0},  //Taiwan New Dollar
		THB: Currency{"THB", "฿", true, 0.0},     //Thailand Baht
		TTD: Currency{"TTD", "TT$", false, 0.0},  //Trinidad and Tobago Dollar
		TRY: Currency{"TRY", "", true, 0.0},      //Turkey Lira
		TVD: Currency{"TVD", "$", false, 0.0},    //Tuvalu Dollar
		UAH: Currency{"UAH", "₴", false, 0.0},    //Ukraine Hryvnia
		GBP: Currency{"GBP", "£", true, 0.0},     //United Kingdom Pound
		USD: Currency{"USD", "$", true, 0.0},     //United States Dollar
		UYU: Currency{"UYU", "$U", false, 0.0},   //Uruguay Peso
		UZS: Currency{"UZS", "лв", false, 0.0},   //Uzbekistan Som
		VEF: Currency{"VEF", "Bs", false, 0.0},   //Venezuela Bolívar
		VND: Currency{"VND", "₫", false, 0.0},    //Viet Nam Dong
		YER: Currency{"YER", "﷼", false, 0.0},    //Yemen Rial
		ZWD: Currency{"ZWD", "Z$", false, 0.0},   //Zimbabwe Dollar
	}
)

/******************************
	Calculators
*******************************/

//Calculate calculates the exchange value
func (data EuroCenterBankRootXML) Calculate(from, to Currency) float32 {
	list := data.Cube.Cube.Cube

	//find the two exchanges
	var v1, v2, v3 float32
	for _, v := range list {
		if v.Currency == from.Acr {
			v1 = v.Rate
		} else if v.Currency == to.Acr {
			v2 = v.Rate
		}
	}

	if v1 <= 0.0 && v2 <= 0.0 {
		util.Catch(errors.New("No exchange rate values found"))
	}

	//calculate let EUR=1.0 [base/from * to]
	EUR := float32(1.0)
	v3 = EUR / v1 * v2

	// util.Logger("Values:", v1, v2, v3)
	return v3
}
