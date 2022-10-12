package main

func main() {
	println("START")

	m := MockYahooClient{}

	y, err := m.Get("AAPC")

	if err != nil {
		panic(err)
	}

	// NewIncomeStatement(y)
	x := NewIncomeStatement(y)

	r := ValueRating{x.Y2018}

	println(r.statement.GrossProfit())
	// println(r.SellingGeneralAdministrative())
	// println(r.InterestExpenseMargin())

	println(r.statement.ResearchDevelopmentMargin())
	println(r.statement.IncomeBeforeTax())
	println(r.statement.IncomeTaxExpense())

	println("END")
}
