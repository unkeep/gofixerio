/*
	Package Fixerio provides a simple interface to the
	fixer.io API, a service for currency exchange rates.
*/
package fixerio

import (
	"encoding/json"
	"errors"
	"net/http"
	urlLib "net/url"
	"strings"
	"time"
)

// Holds the request parameters.
type Request struct {
	base     string
	protocol string
	apiKey   string
	date     string
	symbols  []string
}

// JSON response object.
type Response struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates rates  `json:"rates"`
}

type rates map[string]float32

const host = "data.fixer.io"
const apiPath = "api"

// Initializes fixerio.
func New() *Request {
	return &Request{
		base:     "",
		protocol: "https",
		apiKey:   "",
		date:     "",
		symbols:  make([]string, 0),
	}
}

// Sets base currency.
func (f *Request) Base(currency string) {
	f.base = currency
}

// Make the connection secure or not by setting the
// secure argument to true or false.
func (f *Request) Secure(secure bool) {
	if secure {
		f.protocol = "https"
	} else {
		f.protocol = "http"
	}
}

// List of currencies that should be returned.
func (f *Request) Symbols(currencies ...string) {
	f.symbols = currencies
}

// Specify a date in the past to retrieve historical records.
func (f *Request) Historical(date time.Time) {
	f.date = date.Format("2006-01-02")
}

// Specify a unique key assigned to each API account used to authenticate with the API.
func (f *Request) ApiKey(apiKey string) {
	f.apiKey = apiKey
}

// Retrieve the exchange rates.
func (f *Request) GetRates() (rates, error) {
	url := f.GetUrl()
	response, err := f.makeRequest(url)

	if err != nil {
		return rates{}, err
	}

	return response, nil
}

// Formats the URL correctly for the API Request.
func (f *Request) GetUrl() string {
	url := urlLib.URL{
		Scheme: f.protocol,
		Host:   host,
		Path:   apiPath,
	}

	if f.date == "" {
		url.Path += "/latest"
	} else {
		url.Path += "/" + f.date
	}

	args := urlLib.Values{}

	if f.apiKey != "" {
		args.Set("access_key", f.apiKey)
	}

	if f.base != "" {
		args.Set("base", string(f.base))
	}

	if len(f.symbols) > 0 {
		args.Set("symbols", strings.Join(f.symbols, ","))
	}

	url.RawQuery = args.Encode()

	return url.String()
}

func (f *Request) makeRequest(url string) (rates, error) {
	var response Response
	body, err := http.Get(url)

	if err != nil {
		return rates{}, errors.New("Couldn't connect to server")
	}

	defer body.Body.Close()

	err = json.NewDecoder(body.Body).Decode(&response)

	if err != nil {
		return rates{}, errors.New("Couldn't parse Response")
	}

	return response.Rates, nil
}
