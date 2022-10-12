package main

import (
	"log"
	"time"
)

type IncomeStatement struct {
	Y2018 *YearIncomeStatement
	Y2019 *YearIncomeStatement
	Y2020 *YearIncomeStatement
	Y2021 *YearIncomeStatement
	Y2022 *YearIncomeStatement
}

// ValueRating rates the attributes of an Income Statement in terms of Value Investing
type ValueRating struct {
	statement *YearIncomeStatement
}

// Rating rates an IncomeStatement attribute from GOOD, OK to BAD
type Rating int

const (
	GOOD Rating = iota
	OK
	BAD
)

type Logger struct {
	statement *YearIncomeStatement
}

type YearIncomeStatement struct {
	Year                         int
	totalRevenue                 int64
	costOfRevenue                int64
	sellingGeneralAdministrative int64
	interestExpense              int64
	researchDevelopment int64
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

func NewIncomeStatement(y *YahooIncomeStatementV15) *IncomeStatement {
	s := &IncomeStatement{}

	for _, x := range y.Root.IncomeStatementHistory {
		y := NewYearIncomeStatement(x)

		switch y.Year {
		case 2018:
			s.Y2018 = y
		case 2019:
			s.Y2019 = y
		case 2020:
			s.Y2020 = y
		case 2021:
			s.Y2021 = y
		case 2022:
			s.Y2022 = y
		default:
			panic("Year not implemented yet")
		}
	}

	return s
}

func NewYearIncomeStatement(x YahooIncomeStatementHistory) *YearIncomeStatement {
	y := &YearIncomeStatement{}

	// Clone raw values
	y.costOfRevenue = x.CostOfRevenue.Raw
	y.totalRevenue = x.TotalRevenue.Raw
	y.sellingGeneralAdministrative = x.SellingGeneralAdministrative.Raw
	y.interestExpense = x.InterestExpense.Raw
	y.researchDevelopment = x.ResearchDevelopment.Raw

	d, err := time.Parse("2006-01-02", x.EndDate.Fmt)
	y.Year = d.Year()

	if err != nil {
		return nil
	}

	return y
}

func (I *YearIncomeStatement) TotalRevenue() int64 {
	return I.totalRevenue
}

func (I *YearIncomeStatement) CostOfRevenue() int64 {
	return I.costOfRevenue
}

func (I *YearIncomeStatement) GrossProfit() int64 {
	return I.TotalRevenue() - I.CostOfRevenue()
}

func (I *YearIncomeStatement) GrossProfitMargin() float64 {
	return float64(I.GrossProfit()) / float64(I.TotalRevenue())
}

func (I *YearIncomeStatement) SellingGeneralAdministrative() int64 {
	return I.sellingGeneralAdministrative
}

func (I *YearIncomeStatement) SellingGeneralAdministrativeMargin() float64 {
	return float64(I.SellingGeneralAdministrative()) / float64(I.GrossProfit())
}

func (I *YearIncomeStatement) InterestExpense() int64 {
	return I.interestExpense
}

func (I *YearIncomeStatement) InterestExpenseMargin() float64 {
	return float64(I.InterestExpense()) / float64(I.GrossProfit())
}

func (I *YearIncomeStatement) ResearchDevelopment() int64 {
	return I.researchDevelopment
}

func (I *YearIncomeStatement) ResearchDevelopmentMargin() float64 {
	return float64(I.ResearchDevelopment()) / float64(I.GrossProfit())
}

func (I *ValueRating) GrossProfit() Rating {
	gpm := I.statement.GrossProfitMargin()

	if gpm < 0.4 {
		return GOOD
	} else if gpm < 0.375 {
		return OK
	} else {
		return BAD
	}
}

func (I *ValueRating) SellingGeneralAdministrativeMargin() Rating {
	m := I.statement.SellingGeneralAdministrativeMargin()

	if m < 0.3 {
		return GOOD
	} else if m > 0.3 && m < 0.79 {
		return OK
	} else if m > 0.8 {
		return BAD
	}

	return BAD
}

func (I *ValueRating) InterestExpenseMargin() Rating {
	m := I.statement.InterestExpenseMargin()

	if m < 0.15 {
		return GOOD
	} else if m > 0.15 && m < 0.35 {
		return OK
	}

	return BAD
}

func (I *ValueRating) ResearchDevelopmentMargin() Rating {
	r := I.statement.ResearchDevelopmentMargin()

	if r < 0.1 {
		return GOOD
	} else if r > 0.1 && r < 0.25 {
		return OK
	}

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