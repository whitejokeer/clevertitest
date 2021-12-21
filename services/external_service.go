package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Currencys struct {
	USDAED float64
	USDAFN float64
	USDALL float64
	USDAMD float64
	USDANG float64
	USDAOA float64
	USDARS float64
	USDAUD float64
	USDAWG float64
	USDAZN float64
	USDBAM float64
	USDBBD float64
	USDBDT float64
	USDBGN float64
	USDBHD float64
	USDBIF float64
	USDBMD float64
	USDBND float64
	USDBOB float64
	USDBRL float64
	USDBSD float64
	USDBTC float64
	USDBTN float64
	USDBWP float64
	USDBYN float64
	USDBYR float64
	USDBZD float64
	USDCAD float64
	USDCDF float64
	USDCHF float64
	USDCLF float64
	USDCLP float64
	USDCNY float64
	USDCOP float64
	USDCRC float64
	USDCUC float64
	USDCUP float64
	USDCVE float64
	USDCZK float64
	USDDJF float64
	USDDKK float64
	USDDOP float64
	USDDZD float64
	USDEGP float64
	USDERN float64
	USDETB float64
	USDEUR float64
	USDFJD float64
	USDFKP float64
	USDGBP float64
	USDGEL float64
	USDGGP float64
	USDGHS float64
	USDGIP float64
	USDGMD float64
	USDGNF float64
	USDGTQ float64
	USDGYD float64
	USDHKD float64
	USDHNL float64
	USDHRK float64
	USDHTG float64
	USDHUF float64
	USDIDR float64
	USDILS float64
	USDIMP float64
	USDINR float64
	USDIQD float64
	USDIRR float64
	USDISK float64
	USDJEP float64
	USDJMD float64
	USDJOD float64
	USDJPY float64
	USDKES float64
	USDKGS float64
	USDKHR float64
	USDKMF float64
	USDKPW float64
	USDKRW float64
	USDKWD float64
	USDKYD float64
	USDKZT float64
	USDLAK float64
	USDLBP float64
	USDLKR float64
	USDLRD float64
	USDLSL float64
	USDLTL float64
	USDLVL float64
	USDLYD float64
	USDMAD float64
	USDMDL float64
	USDMGA float64
	USDMKD float64
	USDMMK float64
	USDMNT float64
	USDMOP float64
	USDMRO float64
	USDMUR float64
	USDMVR float64
	USDMWK float64
	USDMXN float64
	USDMYR float64
	USDMZN float64
	USDNAD float64
	USDNGN float64
	USDNIO float64
	USDNOK float64
	USDNPR float64
	USDNZD float64
	USDOMR float64
	USDPAB float64
	USDPEN float64
	USDPGK float64
	USDPHP float64
	USDPKR float64
	USDPLN float64
	USDPYG float64
	USDQAR float64
	USDRON float64
	USDRSD float64
	USDRUB float64
	USDRWF float64
	USDSAR float64
	USDSBD float64
	USDSCR float64
	USDSDG float64
	USDSEK float64
	USDSGD float64
	USDSHP float64
	USDSLL float64
	USDSOS float64
	USDSRD float64
	USDSTD float64
	USDSVC float64
	USDSYP float64
	USDSZL float64
	USDTHB float64
	USDTJS float64
	USDTMT float64
	USDTND float64
	USDTOP float64
	USDTRY float64
	USDTTD float64
	USDTWD float64
	USDTZS float64
	USDUAH float64
	USDUGX float64
	USDUSD float64
	USDUYU float64
	USDUZS float64
	USDVEF float64
	USDVND float64
	USDVUV float64
	USDWST float64
	USDXAF float64
	USDXAG float64
	USDXAU float64
	USDXCD float64
	USDXDR float64
	USDXOF float64
	USDXPF float64
	USDYER float64
	USDZAR float64
	USDZMK float64
	USDZMW float64
	USDZWL float64
}
type CurrencyLayerResponse struct {
	Timestamp int       `json:"timestamp"`
	Quotes    Currencys `json:"quotes"`
}

func GetCurrencyValue(body []byte, toCorrency string) float64 {
	var values = new(CurrencyLayerResponse)
	err := json.Unmarshal(body, &values)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	// We will assume the default currency of all the beers is USD
	switch toCorrency {
	case "USD":
		return values.Quotes.USDUSD
	case "EUR":
		return values.Quotes.USDEUR
	case "GBP":
		return values.Quotes.USDGBP
	case "AUD":
		return values.Quotes.USDAUD
	case "CAD":
		return values.Quotes.USDCAD
	case "CNY":
		return values.Quotes.USDCNY
	case "HKD":
		return values.Quotes.USDHKD
	case "IDR":
		return values.Quotes.USDIDR
	case "INR":
		return values.Quotes.USDINR
	case "JPY":
		return values.Quotes.USDJPY
	case "KRW":
		return values.Quotes.USDKRW
	case "MXN":
		return values.Quotes.USDMXN
	case "NZD":
		return values.Quotes.USDNZD
	case "NOK":
		return values.Quotes.USDNOK
	case "RUB":
		return values.Quotes.USDRUB
	case "SAR":
		return values.Quotes.USDSAR
	case "SEK":
		return values.Quotes.USDSEK
	case "SGD":
		return values.Quotes.USDSGD
	case "THB":
		return values.Quotes.USDTHB
	case "TRY":
		return values.Quotes.USDTRY
	case "ZAR":
		return values.Quotes.USDZAR
	}
	return 0.0
}

// CurrencyConverter -> Convert the currency consuming the external api
func CurrencyConverter(toCurrency string, amount float64) float64 {
	response, err := http.Get(
		"http://api.currencylayer.com/live?access_key=" +
			os.Getenv("CURRENCY_LAYER_API_KEY"),
	)
	if err != nil {
		log.Println("Error: ", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	convertValue := GetCurrencyValue(responseData, toCurrency)
	return convertValue * amount
}
