package main

func main() {
	println("START")

	m := YahooMockClient{}

	y, err := m.GetIncomeStatement("AAPC")

	if err != nil {
		panic(err)
	}

	s, err := m.GetStockInfo("AAPC")

	if err != nil {
		panic(err)
	}

	println("SharesOutstanding")
	println(s.Root.SharesOutstanding)

	x := NewIncomeStatement(y, s)

	println(x.Y2018.PerShareEarnings())

	// r := ValueRating{x.Y2018}

	// println(r.statement.GrossProfit())
	// println(r.SellingGeneralAdministrative())
	// println(r.InterestExpenseMargin())

	// println(r.statement.ResearchDevelopmentMargin())
	// println(r.statement.IncomeBeforeTax())
	// println(r.statement.IncomeTaxExpense())
	// println(r.statement.NetEarnings())

	println("END")
}
