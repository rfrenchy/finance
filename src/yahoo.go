package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type YahooIncomeClient interface {
	Get(code string) (*YahooIncomeStatementV15, error)
}

type YahooStocksClient interface {
	Get(code string) int
}

type YahooAPIClient struct {
	Key    string
	Host   string
	Origin string
}

type MockYahooClient struct{}

type YahooIncomeStatementHistory struct {
	TotalRevenue                 YahooIncomeStatementItem `json:"totalRevenue"`
	CostOfRevenue                YahooIncomeStatementItem `json:"costOfRevenue"`
	GrossProfit                  YahooIncomeStatementItem `json:"grossProfit"`
	SellingGeneralAdministrative YahooIncomeStatementItem `json:"sellingGeneralAdministrative"`
	InterestExpense              YahooIncomeStatementItem `json:"interestExpense"`
	EndDate                      YahooIncomeStatementItem `json:"endDate"`
	ResearchDevelopment          YahooIncomeStatementItem `json:"researchDevelopment"`
	IncomeBeforeTax              YahooIncomeStatementItem `json:"incomeBeforeTax"`
	IncomeTaxExpense             YahooIncomeStatementItem `json:"incomeTaxExpense"`
	NetEarnings                  YahooIncomeStatementItem `json:"netIncome"`
}

type YahooStockStatement struct {
	SharesOutstanding int64 `json:"sharesOutstanding"`
}

type YahooIncomeStatementItem struct {
	Raw     int64  `json:"raw"`
	Fmt     string `json:"fmt"`
	LongFmt string `json:"longFmt,omitempty"`
}

type YahooIncomeStatementV15 struct {
	Root struct {
		IncomeStatementHistory []YahooIncomeStatementHistory `json:"incomeStatementHistory"`
	} `json:"incomeStatementHistory"`
}

type YahooStockV97 struct {
	Root struct {
		YahooStockStatement YahooStockStatement `json:"data"`
	}
}

func NewYahooAPIClient() *YahooAPIClient {
	return &YahooAPIClient{
		Key:    os.Getenv("RAPID_API_YAHOO_KEY"),
		Host:   "yahoo-finance15.p.rapidapi.com",
		Origin: "https://yahoo-finance15.p.rapidapi.com/api/yahoo",
	}
}

func (m *MockYahooClient) Get(code string) (*YahooIncomeStatementV15, error) {
	ic, err := os.Open("./income.json")

	if err != nil {
		return nil, err
	}

	defer ic.Close()

	b, err := ioutil.ReadAll(ic)

	if err != nil {
		return nil, err
	}

	var x *YahooIncomeStatementV15
	err = json.Unmarshal(b, &x)

	return x, err
}

func (y *YahooAPIClient) Get(code string) (*YahooIncomeStatementV15, error) {
	url := fmt.Sprintf("%s/qu/quote/%s/income-statement", y.Origin, code)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", y.Key)
	req.Header.Add("X-RapidAPI-Host", y.Host)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var x *YahooIncomeStatementV15
	err = json.Unmarshal(body, &x)

	return x, nil
}

func (y *YahooAPIClient) GetStockInfo(code string) (*YahooStockV97, error) {
	payload := strings.NewReader("symbol=" + code)

	req, _ := http.NewRequest("POST", "https://yahoo-finance97.p.rapidapi.com/stock-info", payload)
	req.Header.Add("X-RapidAPI-Key", y.Key)
	req.Header.Add("X-RapidAPI-Host", y.Host)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	var x *YahooStockV97
	err = json.Unmarshal(body, &x)

	return x, err
}
