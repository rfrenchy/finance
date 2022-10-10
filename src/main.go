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

	// println(x.GrossProfit())
	// println(r.SellingGeneralAdministrative())
	println(r.InterestExpenseMargin())

	println("END")
}
