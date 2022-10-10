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
	TotalRevenue                 int64
	CostOfRevenue                int64
	SellingGeneralAdministrative int64
	InterestExpense              int64
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
	y.CostOfRevenue = x.CostOfRevenue.Raw
	y.TotalRevenue = x.TotalRevenue.Raw
	y.SellingGeneralAdministrative = x.SellingGeneralAdministrative.Raw
	y.InterestExpense = x.InterestExpense.Raw

	d, err := time.Parse("2006-01-02", x.EndDate.Fmt)
	y.Year = d.Year()

	if err != nil {
		return nil
	}

	return y
}

func (I *YearIncomeStatement) GrossProfit() int64 {
	return I.TotalRevenue - I.CostOfRevenue
}

func (I *YearIncomeStatement) GrossProfitMargin() float64 {
	gp := I.GrossProfit()

	gpm := float64(gp) / float64(I.TotalRevenue)

	return gpm
}

func (I *YearIncomeStatement) SellingGeneralAdministrativeMargin() float64 {
	gp := I.GrossProfit()

	sgam := float64(I.SellingGeneralAdministrative) / float64(gp)	

	return sgam
}

func (I *YearIncomeStatement) InterestExpenseMargin() float64 {
	gp := I.GrossProfit()

	iem := float64(I.InterestExpense) / float64(gp)	

	return iem
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
