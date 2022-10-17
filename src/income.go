package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/leekchan/accounting"
)

type IncomeStatement struct {
	Y2018 *YearIncomeStatement
	Y2019 *YearIncomeStatement
	Y2020 *YearIncomeStatement
	Y2021 *YearIncomeStatement
	// Y2022 *YearIncomeStatement
}

type Logger struct {
	statement *YearIncomeStatement
}

type YearIncomeStatement struct {
	Year                         int
	totalRevenue                 int64
	costOfRevenue                int64
	sellingGeneralAdministrative int64
	interestExpense              int64
	researchDevelopment          int64
	incomeBeforeTax              int64
	incomeTaxExpense             int64
	netEarnings                  int64
	sharesOutstanding            int64
}

type IncomeMargins[T any] interface {
	GrossProfitMargin() T
	SellingGeneralAdministrativeMargin() T
	InterestExpenseMargin() T
}

type IncomeAttributes[T any] interface {
	GrossProfit() T
}

type IncomeOptions struct {
	logging bool
}

func NewIncomeStatement(y *YahooIncomeStatementV15, ysi *YahooStockInfo) *IncomeStatement {
	s := &IncomeStatement{}

	for _, x := range y.Root.IncomeStatementHistory {
		y := NewYearIncomeStatement(x, ysi)

		switch y.Year {
		case 2018:
			s.Y2018 = y
		case 2019:
			s.Y2019 = y
		case 2020:
			s.Y2020 = y
		case 2021:
			s.Y2021 = y
		// case 2022:
		// 	s.Y2022 = y
		default:
			panic("Year not implemented yet")
		}
	}

	return s
}

func NewYearIncomeStatement(yish YahooIncomeStatementHistory, ysi *YahooStockInfo) *YearIncomeStatement {
	y := &YearIncomeStatement{}

	// Clone raw values
	y.costOfRevenue = yish.CostOfRevenue.Raw
	y.totalRevenue = yish.TotalRevenue.Raw
	y.sellingGeneralAdministrative = yish.SellingGeneralAdministrative.Raw
	y.interestExpense = yish.InterestExpense.Raw
	y.researchDevelopment = yish.ResearchDevelopment.Raw
	y.incomeBeforeTax = yish.IncomeBeforeTax.Raw
	y.incomeTaxExpense = yish.IncomeTaxExpense.Raw
	y.netEarnings = yish.NetEarnings.Raw

	y.sharesOutstanding = ysi.Root.SharesOutstanding

	d, err := time.Parse("2006-01-02", yish.EndDate.Fmt)
	y.Year = d.Year()

	if err != nil {
		return nil
	}

	return y
}

// TotalRevenue
func (I *YearIncomeStatement) TotalRevenue() int64 {
	return I.totalRevenue
}

// CostOfRevenue
func (I *YearIncomeStatement) CostOfRevenue() int64 {
	return I.costOfRevenue
}

// GrossProfit
func (I *YearIncomeStatement) GrossProfit() int64 {
	return I.TotalRevenue() - I.CostOfRevenue()
}

// GrossProfitMargin
func (I *YearIncomeStatement) GrossProfitMargin() float64 {
	return float64(I.GrossProfit()) / float64(I.TotalRevenue())
}

// SellingGeneralAdministrative
func (I *YearIncomeStatement) SellingGeneralAdministrative() int64 {
	return I.sellingGeneralAdministrative
}

// SellingGeneralAdministrativeMargin
func (I *YearIncomeStatement) SellingGeneralAdministrativeMargin() float64 {
	return float64(I.SellingGeneralAdministrative()) / float64(I.GrossProfit())
}

// InterestExpense
func (I *YearIncomeStatement) InterestExpense() int64 {
	return I.interestExpense
}

// InterestExpenseMargin
func (I *YearIncomeStatement) InterestExpenseMargin() float64 {
	return float64(I.InterestExpense()) / float64(I.GrossProfit())
}

// ResearchDevelopment
func (I *YearIncomeStatement) ResearchDevelopment() int64 {
	return I.researchDevelopment
}

// ResearchDevelopmentMargin
func (I *YearIncomeStatement) ResearchDevelopmentMargin() float64 {
	return float64(I.ResearchDevelopment()) / float64(I.GrossProfit())
}

// IncomeBeforeTax
func (I *YearIncomeStatement) IncomeBeforeTax() int64 {
	return I.incomeBeforeTax
}

// IncomeTaxExpense
func (I *YearIncomeStatement) IncomeTaxExpense() int64 {
	return I.incomeTaxExpense
}

// NetEarnings calculates (Gross Profit - Expenses - Taxes)
func (I *YearIncomeStatement) NetEarnings() int64 {
	return I.netEarnings
}

// SharesOutstanding returns the total amount of available shares
func (I *YearIncomeStatement) SharesOutstanding() int64 {
	return I.sharesOutstanding
}

// PerShareEarnings returns the total earnings per share (NetEarnings / SharesOutstanding)
func (I *YearIncomeStatement) PerShareEarnings() float64 {
	return float64(I.NetEarnings()) / float64(I.SharesOutstanding())
}

func (I *IncomeStatement) PerShareEarningsMean() float64 {
	x := []float64{I.Y2018.PerShareEarnings(), I.Y2019.PerShareEarnings(), I.Y2020.PerShareEarnings(), I.Y2021.PerShareEarnings()}

	var sum float64 = 0
	for _, e := range x {
		sum += e
	}

	return sum / float64(len(x))
}

func (I *IncomeStatement) PerShareEarningsSTD() float64 {
	mean := I.PerShareEarningsMean()

	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	fmt.Println("2018 PerShareEarnings", ac.FormatMoney(I.Y2018.PerShareEarnings()))
	fmt.Println("2019 PerShareEarnings", ac.FormatMoney(I.Y2019.PerShareEarnings()))
	fmt.Println("2020 PerShareEarnings", ac.FormatMoney(I.Y2020.PerShareEarnings()))
	fmt.Println("2021 PerShareEarnings", ac.FormatMoney(I.Y2021.PerShareEarnings()))

	// deviations
	d2018 := float64(I.Y2018.PerShareEarnings()) - mean
	d2019 := float64(I.Y2019.PerShareEarnings()) - mean
	d2020 := float64(I.Y2020.PerShareEarnings()) - mean
	d2021 := float64(I.Y2021.PerShareEarnings()) - mean

	// fmt.Println("2018 PerShareEarnings", ac.FormatMoney(I.Y2018.PerShareEarnings()))
	// fmt.Println("2019 PerShareEarnings", ac.FormatMoney(I.Y2019.PerShareEarnings()))
	// fmt.Println("2020 PerShareEarnings", ac.FormatMoney(I.Y2020.PerShareEarnings()))
	// fmt.Println("2021 PerShareEarnings", ac.FormatMoney(I.Y2021.PerShareEarnings()))

	sample := 4 - 1

	variance := (math.Pow(d2018, 2) + math.Pow(d2019, 2) + math.Pow(d2020, 2) + math.Pow(d2021, 2)) / float64(sample)
	fmt.Println("Variance", ac.FormatMoney(variance))

	// standard deviation
	std := math.Sqrt(variance)
	fmt.Println("std", ac.FormatMoney(std))

	return std
}

func (I *IncomeStatement) NetEarningsMean() float64 {
	x := []int64{I.Y2018.NetEarnings(), I.Y2019.NetEarnings(), I.Y2020.NetEarnings(), I.Y2021.NetEarnings()}

	// fmt.Println("2018 NetEarnings", ac.FormatMoney(I.Y2018.NetEarnings()))
	// fmt.Println("2019 NetEarnings", ac.FormatMoney(I.Y2019.NetEarnings()))
	// fmt.Println("2020 NetEarnings", ac.FormatMoney(I.Y2020.NetEarnings()))
	// fmt.Println("2021 NetEarnings", ac.FormatMoney(I.Y2021.NetEarnings()))

	var sum float64 = 0
	for _, e := range x {
		sum += float64(e)
	}

	return sum / float64(len(x))
}

// NetEarningsSTD calculate the Standard Deviation of NetEarnings over x years
func (I *IncomeStatement) NetEarningsSTD() float64 {
	// ac := accounting.Accounting{Symbol: "$", Precision: 2 }

	mean := I.NetEarningsMean()
	// fmt.Println("Mean", ac.FormatMoney(mean))

	// deviations
	d2018 := float64(I.Y2018.NetEarnings()) - mean
	d2019 := float64(I.Y2019.NetEarnings()) - mean
	d2020 := float64(I.Y2020.NetEarnings()) - mean
	d2021 := float64(I.Y2021.NetEarnings()) - mean

	sample := 4 - 1

	variance := (math.Pow(d2018, 2) + math.Pow(d2019, 2) + math.Pow(d2020, 2) + math.Pow(d2021, 2)) / float64(sample)
	// fmt.Println("Variance", ac.FormatMoney(variance))

	// standard deviation
	std := math.Sqrt(variance)
	// fmt.Println("std", ac.FormatMoney(std))

	return std
}

func (I *IncomeStatement) NetEarnings() Rating {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	std := I.NetEarningsSTD()

	fmt.Println("std", ac.FormatMoney(std))

	// have to be within standard deviation?
	// and not below it consistently?
	// standard deviation within x
	// also compare to competitors ??

	return BAD
}

func (I *IncomeStatement) PerShareEarnings() Rating {

	I.Y2018.PerShareEarnings()
	I.Y2019.PerShareEarnings()
	I.Y2020.PerShareEarnings()
	I.Y2021.PerShareEarnings()
	// I.Y2022.PerShareEarnings()

	// compare all, look for upward trend

	// also compare to competitors ??

	return BAD
}

func (I *Logger) GrossProfit() {
	log.Println("GrossProfit", I.statement.GrossProfit())
}

func (I *Logger) GrossProfitMargin() {
	log.Printf("GrossProfitMargin (TotalRevenue / Gross Profit)%f\n", I.statement.GrossProfitMargin())
}

func (I *Logger) SellingGeneralAdministrativeMargin() {
	log.Printf("SellingGeneralAdministrativeMargin (SellingGeneralAdministrative / Gross Profit)%f\n", I.statement.SellingGeneralAdministrativeMargin())
}

func (I *Logger) InterestExpenseMargin() {
	log.Printf("InterestExpenseMargin (Interest Expense / Gross Profit)%f\n", I.statement.InterestExpenseMargin())
}

func (I *Logger) ResearchDevelopmentMargin() {
	log.Printf("ResearchDevelopmentMargin ( Research Development / Gross Profit)%f\n", I.statement.ResearchDevelopmentMargin())
}

func (I *Logger) IncomeBeforeTax() {
	log.Println("IncomeBeforeTax", I.statement.IncomeBeforeTax())
}

func (I *Logger) NetEarnings() {
	log.Println("NetEarnings", I.statement.NetEarnings())
}

func (I *Logger) SharesOutstanding() {
	log.Println("SharesOutstanding", I.statement.SharesOutstanding())
}

func (I *Logger) PerShareEarnings() {
	log.Printf("PerShareEarnings ( Net Earnings / Shares Outstanding)%f\n", I.statement.PerShareEarnings())
}
