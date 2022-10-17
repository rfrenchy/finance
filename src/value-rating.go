package main

// Rating rates an IncomeStatement attribute from GOOD, OK to BAD
type Rating int

const (
	GOOD Rating = iota
	OK
	BAD
)

// ValueRating rates the Finances of a business in terms of Value Investing
type ValueRating struct {
	income  *YearIncomeStatement
	balance *YearBalanceSheet
}

type LegitimacyRating struct {
	income *YearIncomeStatement
}

func (I *ValueRating) GrossProfit() Rating {
	gpm := I.income.GrossProfitMargin()

	if gpm < 0.4 {
		return GOOD
	} else if gpm < 0.375 {
		return OK
	} else {
		return BAD
	}
}

func (I *ValueRating) SellingGeneralAdministrativeMargin() Rating {
	m := I.income.SellingGeneralAdministrativeMargin()

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
	m := I.income.InterestExpenseMargin()

	if m < 0.15 {
		return GOOD
	} else if m > 0.15 && m < 0.35 {
		return OK
	}

	return BAD
}

func (I *ValueRating) ResearchDevelopmentMargin() Rating {
	r := I.income.ResearchDevelopmentMargin()

	if r < 0.1 {
		return GOOD
	} else if r > 0.1 && r < 0.25 {
		return OK
	}

	return BAD
}

// IncomeTaxExpense
func (I *LegitimacyRating) IncomeTaxExpense() Rating {
	// t := I.statement.IncomeTaxExpense()

	// Compare against income taxes paid
	// Review how much tax they should have to pay in UK/Country??? 19% in the UK???

	return BAD
}

// CurrentRatio
func (I *ValueRating) CurrentRatio() Rating {
	if I.balance.CurrentRatio() > 1 {
		return GOOD
	}

	return BAD
}

// DebtToShareholderEquityRatio
func (I *ValueRating) DebtToShareholderEquityRatio() Rating {
	if I.balance.DebtToShareholderEquityRatio() <= 0.8 {
		return GOOD
	}

	return BAD
}

// ShortVsLongTermDebt
func (I *ValueRating) ShortVsLongTermDebt() Rating {
	if I.balance.ShortTermDebt() < I.balance.LongTermDebt() {
		return GOOD
	}

	return BAD
}
