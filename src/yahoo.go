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

type YahooMockClient struct{}

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

type YahooStockInfo struct {
	Root struct {
		SharesOutstanding int64 `json:"sharesOutstanding"`
	} `json:"data"`
}

type YahooBalanceSheetV1 struct {
	Root struct {
		TotalLiabilities int `json:"totalLiabilities"`
	} `json:"data"`
}

type YahooCashFlowV1 struct {
	Root struct {
	} `json:"data"`
}

type YahooCashFlow struct {
	Investments                           int
	ChangeToLiabilities                   int
	TotalCashFlowsFromInvestingActivities int
	NetBorrowings                         int
	TotalCashFromFinancingActivities      int
	ChangeToOperatingActivities           int
	IssuanceOfStock                       int
	NetIncome                             int
	ChangeInCash                          int
	RepurchaseOfStock                     int
	Depreciation                          int
}

func NewYahooAPIClient() *YahooAPIClient {
	return &YahooAPIClient{
		Key:    os.Getenv("RAPID_API_YAHOO_KEY"),
		Host:   "yahoo-finance15.p.rapidapi.com",
		Origin: "https://yahoo-finance15.p.rapidapi.com/api/yahoo",
	}
}

func (m *YahooMockClient) GetIncomeStatement(code string) (*YahooIncomeStatementV15, error) {
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

func (y *YahooAPIClient) GetIncomeStatement(code string) (*YahooIncomeStatementV15, error) {
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

func (y *YahooAPIClient) GetStockInfo(code string) (*YahooStockInfo, error) {
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
	var x *YahooStockInfo
	err = json.Unmarshal(body, &x)

	return x, err
}

func (m *YahooMockClient) GetStockInfo(code string) (*YahooStockInfo, error) {
	ic, err := os.Open("./stock.json")

	if err != nil {
		return nil, err
	}

	defer ic.Close()

	b, err := ioutil.ReadAll(ic)

	if err != nil {
		return nil, err
	}

	var x *YahooStockInfo
	err = json.Unmarshal(b, &x)

	return x, err
}

func (y *YahooAPIClient) GetBalanceSheet(code string) (*YahooBalanceSheetV1, error) {
	url := fmt.Sprintf("%s/balance-sheet", y.Origin)

	payload := strings.NewReader("symbol=" + code)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
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

	var x *YahooBalanceSheetV1
	err = json.Unmarshal(body, &x)

	return x, err
}

func (y *YahooMockClient) GetBalanceSheet(code string) (*YahooBalanceSheetV1, error) {
	ic, err := os.Open("./balance.json")

	if err != nil {
		return nil, err
	}

	defer ic.Close()

	b, err := ioutil.ReadAll(ic)

	if err != nil {
		return nil, err
	}

	var x *YahooBalanceSheetV1
	err = json.Unmarshal(b, &x)

	return x, err
}

func (y *YahooAPIClient) GetCashFlow(code string) (*YahooCashFlowV1, error) {
	url := fmt.Sprintf("%s/cashflow", y.Origin)

	payload := strings.NewReader("symbol=" + code)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
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

	var x *YahooCashFlowV1
	err = json.Unmarshal(body, &x)

	return x, err
}
