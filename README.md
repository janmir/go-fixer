# :bento: Overview
:money_with_wings: __Go-Fixer__ is a simple foreign exchange api library for Golang.

# :see_no_evil: Features
- [x] Multi-sources
- [x] Local store
- [] Trend graph


# :card_file_box: Data Sources
- [Currency Converter](http://free.currencyconverterapi.com/api/v5/convert?q=JPY_PHP&compact=y)
- [European Bank Daily Conversions](http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml)
- [European Bank Conversions](http://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html)
- [European Bank 90-Days Conversions](https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml)

# :package: Storage
- [Bolt Key/Value Store Github](https://github.com/boltdb/bolt)
- [Bolt Key/Value Store Blog](https://npf.io/2014/07/intro-to-boltdb-painless-performant-persistence/)

# :page_facing_up: Snippets
Acquiring Currency List
> Listing struct fields `source: https://www.xe.com/symbols.php`
```js
    //e.g JPY Currency //Japanese Yen
    document.querySelectorAll("table.currencySymblTable tr").forEach(el => {
        var name = el.querySelector("td:first-child").innerText
        var acr = el.querySelector("td:nth-child(2)").innerText
        var sym = el.querySelector("td:nth-child(4)").innerText
        console.log(acr + " Currency" + " //"+ name)
    })
```
> Listing struct values
```js
    //e.g JPY: Currency{"", ""}, //Japanese Yen
    document.querySelectorAll("table.currencySymblTable tr").forEach(el => {
        var name = el.querySelector("td:first-child").innerText
        var acr = el.querySelector("td:nth-child(2)").innerText
        var sym = el.querySelector("td:nth-child(4)").innerText
        console.log(acr + ": Currency{\""+acr+"\",\""+sym+"\"} //" + name )
    })
```