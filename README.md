# go-fixer

# Data Sources
- [Currency Converter](http://free.currencyconverterapi.com/api/v5/convert?q=JPY_PHP&compact=y)
- [European Bank Daily Conversions](http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml)
- [European Bank Conversions](http://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html)
- [European Bank 90-Days Conversions](https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml)

# Acquiring Currency List
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