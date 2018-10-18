package fixer

//Currency structure for a type of currency
type Currency struct {
	Acr string //currency acronyms i.e JPY, USD
	Sym string //symbols for the currency e.g $, ¥, ₡, ₱, ₩ etc.
	Ok  bool   //supported or not
}

//Source defines datasources
type Source struct {
	ty, url string
}

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

//EuroCenterBankXML Europen Bank XMl data structure
type EuroCenterBankXML struct {
	Subject string `xml:"subject"`
	Sender  struct {
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube CubeParent `xml:"Cube"`
}

//EUData xml data structure
type EUData struct {
	Timestamp string `json:"timestamp"`
	ImageURL  string `json:"img"`
	History   []struct {
		From  string  `json:"from"`
		To    string  `json:"to"`
		Value float32 `json:"value"`
	} `json:"history"`
}

var (
	_offline = true
	_version = "0.0.1"

	_sources = []Source{
		{"xml", "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"},
		{"api", "http://free.currencyconverterapi.com/api/v5/convert?q=JPY_PHP&compact=y"},
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
		ALL: Currency{"ALL", "Lek", false},  //Albania Lek
		AFN: Currency{"AFN", "؋", false},    //Afghanistan Afghani
		ARS: Currency{"ARS", "$", false},    //Argentina Peso
		AWG: Currency{"AWG", "ƒ", false},    //Aruba Guilder
		AUD: Currency{"AUD", "$", false},    //Australia Dollar
		AZN: Currency{"AZN", "₼", false},    //Azerbaijan Manat
		BSD: Currency{"BSD", "$", false},    //Bahamas Dollar
		BBD: Currency{"BBD", "$", false},    //Barbados Dollar
		BYN: Currency{"BYN", "Br", false},   //Belarus Ruble
		BZD: Currency{"BZD", "BZ$", false},  //Belize Dollar
		BMD: Currency{"BMD", "$", false},    //Bermuda Dollar
		BOB: Currency{"BOB", "$b", false},   //Bolivia Bolíviano
		BAM: Currency{"BAM", "KM", false},   //Bosnia and Herzegovina Convertible Marka
		BWP: Currency{"BWP", "P", false},    //Botswana Pula
		BGN: Currency{"BGN", "лв", false},   //Bulgaria Lev
		BRL: Currency{"BRL", "R$", false},   //Brazil Real
		BND: Currency{"BND", "$", false},    //Brunei Darussalam Dollar
		KHR: Currency{"KHR", "៛", false},    //Cambodia Riel
		CAD: Currency{"CAD", "$", false},    //Canada Dollar
		KYD: Currency{"KYD", "$", false},    //Cayman Islands Dollar
		CLP: Currency{"CLP", "$", false},    //Chile Peso
		CNY: Currency{"CNY", "¥", false},    //China Yuan Renminbi
		COP: Currency{"COP", "$", false},    //Colombia Peso
		CRC: Currency{"CRC", "₡", false},    //Costa Rica Colon
		HRK: Currency{"HRK", "kn", false},   //Croatia Kuna
		CUP: Currency{"CUP", "₱", false},    //Cuba Peso
		CZK: Currency{"CZK", "Kč", false},   //Czech Republic Koruna
		DKK: Currency{"DKK", "kr", false},   //Denmark Krone
		DOP: Currency{"DOP", "RD$", false},  //Dominican Republic Peso
		XCD: Currency{"XCD", "$", false},    //East Caribbean Dollar
		EGP: Currency{"EGP", "£", false},    //Egypt Pound
		SVC: Currency{"SVC", "$", false},    //El Salvador Colon
		EUR: Currency{"EUR", "€", false},    //Euro Member Countries
		FKP: Currency{"FKP", "£", false},    //Falkland Islands (Malvinas) Pound
		FJD: Currency{"FJD", "$", false},    //Fiji Dollar
		GHS: Currency{"GHS", "¢", false},    //Ghana Cedi
		GIP: Currency{"GIP", "£", false},    //Gibraltar Pound
		GTQ: Currency{"GTQ", "Q", false},    //Guatemala Quetzal
		GGP: Currency{"GGP", "£", false},    //Guernsey Pound
		GYD: Currency{"GYD", "$", false},    //Guyana Dollar
		HNL: Currency{"HNL", "L", false},    //Honduras Lempira
		HKD: Currency{"HKD", "$", false},    //Hong Kong Dollar
		HUF: Currency{"HUF", "Ft", false},   //Hungary Forint
		ISK: Currency{"ISK", "kr", false},   //Iceland Krona
		INR: Currency{"INR", "", false},     //India Rupee
		IDR: Currency{"IDR", "Rp", false},   //Indonesia Rupiah
		IRR: Currency{"IRR", "﷼", false},    //Iran Rial
		IMP: Currency{"IMP", "£", false},    //Isle of Man Pound
		ILS: Currency{"ILS", "₪", false},    //Israel Shekel
		JMD: Currency{"JMD", "J$", false},   //Jamaica Dollar
		JPY: Currency{"JPY", "¥", false},    //Japan Yen
		JEP: Currency{"JEP", "£", false},    //Jersey Pound
		KZT: Currency{"KZT", "лв", false},   //Kazakhstan Tenge
		KPW: Currency{"KPW", "₩", false},    //Korea (North) Won
		KRW: Currency{"KRW", "₩", false},    //Korea (South) Won
		KGS: Currency{"KGS", "лв", false},   //Kyrgyzstan Som
		LAK: Currency{"LAK", "₭", false},    //Laos Kip
		LBP: Currency{"LBP", "£", false},    //Lebanon Pound
		LRD: Currency{"LRD", "$", false},    //Liberia Dollar
		MKD: Currency{"MKD", "ден", false},  //Macedonia Denar
		MYR: Currency{"MYR", "RM", false},   //Malaysia Ringgit
		MUR: Currency{"MUR", "₨", false},    //Mauritius Rupee
		MXN: Currency{"MXN", "$", false},    //Mexico Peso
		MNT: Currency{"MNT", "₮", false},    //Mongolia Tughrik
		MZN: Currency{"MZN", "MT", false},   //Mozambique Metical
		NAD: Currency{"NAD", "$", false},    //Namibia Dollar
		NPR: Currency{"NPR", "₨", false},    //Nepal Rupee
		ANG: Currency{"ANG", "ƒ", false},    //Netherlands Antilles Guilder
		NZD: Currency{"NZD", "$", false},    //New Zealand Dollar
		NIO: Currency{"NIO", "C$", false},   //Nicaragua Cordoba
		NGN: Currency{"NGN", "₦", false},    //Nigeria Naira
		NOK: Currency{"NOK", "kr", false},   //Norway Krone
		OMR: Currency{"OMR", "﷼", false},    //Oman Rial
		PKR: Currency{"PKR", "₨", false},    //Pakistan Rupee
		PAB: Currency{"PAB", "B/.", false},  //Panama Balboa
		PYG: Currency{"PYG", "Gs", false},   //Paraguay Guarani
		PEN: Currency{"PEN", "S/.", false},  //Peru Sol
		PHP: Currency{"PHP", "₱", false},    //Philippines Piso
		PLN: Currency{"PLN", "zł", false},   //Poland Zloty
		QAR: Currency{"QAR", "﷼", false},    //Qatar Riyal
		RON: Currency{"RON", "lei", false},  //Romania Leu
		RUB: Currency{"RUB", "₽", false},    //Russia Ruble
		SHP: Currency{"SHP", "£", false},    //Saint Helena Pound
		SAR: Currency{"SAR", "﷼", false},    //Saudi Arabia Riyal
		RSD: Currency{"RSD", "Дин.", false}, //Serbia Dinar
		SCR: Currency{"SCR", "₨", false},    //Seychelles Rupee
		SGD: Currency{"SGD", "$", false},    //Singapore Dollar
		SBD: Currency{"SBD", "$", false},    //Solomon Islands Dollar
		SOS: Currency{"SOS", "S", false},    //Somalia Shilling
		ZAR: Currency{"ZAR", "R", false},    //South Africa Rand
		LKR: Currency{"LKR", "₨", false},    //Sri Lanka Rupee
		SEK: Currency{"SEK", "kr", false},   //Sweden Krona
		CHF: Currency{"CHF", "CHF", false},  //Switzerland Franc
		SRD: Currency{"SRD", "$", false},    //Suriname Dollar
		SYP: Currency{"SYP", "£", false},    //Syria Pound
		TWD: Currency{"TWD", "NT$", false},  //Taiwan New Dollar
		THB: Currency{"THB", "฿", false},    //Thailand Baht
		TTD: Currency{"TTD", "TT$", false},  //Trinidad and Tobago Dollar
		TRY: Currency{"TRY", "", false},     //Turkey Lira
		TVD: Currency{"TVD", "$", false},    //Tuvalu Dollar
		UAH: Currency{"UAH", "₴", false},    //Ukraine Hryvnia
		GBP: Currency{"GBP", "£", false},    //United Kingdom Pound
		USD: Currency{"USD", "$", false},    //United States Dollar
		UYU: Currency{"UYU", "$U", false},   //Uruguay Peso
		UZS: Currency{"UZS", "лв", false},   //Uzbekistan Som
		VEF: Currency{"VEF", "Bs", false},   //Venezuela Bolívar
		VND: Currency{"VND", "₫", false},    //Viet Nam Dong
		YER: Currency{"YER", "﷼", false},    //Yemen Rial
		ZWD: Currency{"ZWD", "Z$", false},   //Zimbabwe Dollar
	}
)
