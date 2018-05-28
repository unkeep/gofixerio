package fixerio

import (
	"testing"
	"time"
)

const baseUrl = host + "/" + apiPath

func TestDefaultParameters(t *testing.T) {
	expected := "https://" + baseUrl + "/latest"
	actual := New().GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestUnsecure(t *testing.T) {
	expected := "http://" + baseUrl + "/latest"

	f := New()
	f.Secure(false)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestBase(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=USD"

	f := New()
	f.Base(USD)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSymbols(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=GBP&symbols=EUR%2CUSD%2CAUD"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR, USD, AUD)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSingleSymbol(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=GBP&symbols=EUR"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestEmptySymbols(t *testing.T) {
	expected := "https://" + baseUrl + "/latest"

	f := New()
	f.Symbols()
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestHistorical(t *testing.T) {
	expected := "https://" + baseUrl + "/2016-06-09"

	f := New()
	f.Historical(time.Date(2016, time.June, 9, 0, 0, 0, 0, time.UTC))
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestApiKey(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?access_key=fake_api_key"

	f := New()
	f.ApiKey("fake_api_key")
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestAllParameters(t *testing.T) {
	expected := "http://" + baseUrl + "/latest?access_key=fake_api_key&base=USD&symbols=EUR%2CGBP"

	f := New()
	f.ApiKey("fake_api_key")
	f.Base(USD)
	f.Symbols(EUR, GBP)
	f.Secure(false)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
