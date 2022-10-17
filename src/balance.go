package main

import "strings"

type BalanceSheet struct {
	Y2018 *YearBalanceSheet
	Y2019 *YearBalanceSheet
	Y2020 *YearBalanceSheet
	Y2021 *YearBalanceSheet
	// Y2022
}

type YearBalanceSheet struct {
	totalCurrentAssets      int64
	totalCurrentLiabilities int64
	totalLiabilities        int64
	totalShareholdersEquity int64
	shortTermDebt           int64
	longTermDebt            int64
	totalAssets             int64
}

// NewBalanceSheet creates a BalanceSheet from Yahoo API data
func NewBalanceSheet(ybs *YahooBalanceSheetV1) *BalanceSheet {
	bs := BalanceSheet{
		Y2018: &YearBalanceSheet{},
		Y2019: &YearBalanceSheet{},
		Y2020: &YearBalanceSheet{},
		Y2021: &YearBalanceSheet{},
	}

	for _, item := range ybs.Root {
		name := strings.ToUpper(strings.ReplaceAll(item.Name, " ", ""))

		switch name {
		case "TOTALCURRENTASSETS":
			bs.Y2018.totalCurrentAssets = item.Y2018
			bs.Y2019.totalCurrentAssets = item.Y2019
			bs.Y2020.totalCurrentAssets = item.Y2020
			bs.Y2021.totalCurrentAssets = item.Y2021
		case "TOTALCURRENTLIABILITIES":
			bs.Y2018.totalCurrentLiabilities = item.Y2018
			bs.Y2019.totalCurrentLiabilities = item.Y2019
			bs.Y2020.totalCurrentLiabilities = item.Y2020
			bs.Y2021.totalCurrentLiabilities = item.Y2021
		case "TOTALLIAB":
			bs.Y2018.totalLiabilities = item.Y2018
			bs.Y2019.totalLiabilities = item.Y2019
			bs.Y2020.totalLiabilities = item.Y2020
			bs.Y2021.totalLiabilities = item.Y2021
		case "TOTALSTOCKHOLDEREQUITY":
			bs.Y2018.totalShareholdersEquity = item.Y2018
			bs.Y2019.totalShareholdersEquity = item.Y2019
			bs.Y2020.totalShareholdersEquity = item.Y2020
			bs.Y2021.totalShareholdersEquity = item.Y2021
		case "SHORTLONGTERMDEBT":
			bs.Y2018.shortTermDebt = item.Y2018
			bs.Y2019.shortTermDebt = item.Y2019
			bs.Y2020.shortTermDebt = item.Y2020
			bs.Y2021.shortTermDebt = item.Y2021
		case "LONGTERMDEBT":
			bs.Y2018.longTermDebt = item.Y2018
			bs.Y2019.longTermDebt = item.Y2019
			bs.Y2020.longTermDebt = item.Y2020
			bs.Y2021.longTermDebt = item.Y2021
		case "TOTALASSETS":
			bs.Y2018.totalAssets = item.Y2018
			bs.Y2019.totalAssets = item.Y2019
			bs.Y2020.totalAssets = item.Y2020
			bs.Y2021.totalAssets = item.Y2021
		}

	}

	return &bs
}

// TotalAssets
func (b *YearBalanceSheet) TotalAssets() int64 {
	return b.totalAssets
}

// TotalCurrentAssets
func (b *YearBalanceSheet) TotalCurrentAssets() int64 {
	return b.totalCurrentAssets
}

// TotalCurrentLiabilities
func (b *YearBalanceSheet) TotalCurrentLiabilities() int64 {
	return b.totalCurrentLiabilities
}

// TotalLiabilities
func (b *YearBalanceSheet) TotalLiabilities() int64 {
	return b.totalLiabilities
}

// TotalShareholdersEquity (Straight from Yahoo) aka BookValue
func (b *YearBalanceSheet) TotalShareholdersEquity() int64 {
	return b.totalShareholdersEquity
}

// ShareholdersEquity (TotalAssets - TotalLiabilities) aka BookValue
func (b *YearBalanceSheet) ShareholdersEquity() int64 {
	return b.TotalAssets() - b.TotalLiabilities()
}

// ShortTermDebt
func (b *YearBalanceSheet) ShortTermDebt() int64 {
	return b.shortTermDebt
}

// LongTermDebt
func (b *YearBalanceSheet) LongTermDebt() int64 {
	return b.longTermDebt
}

// Current Ratio (TotalCurrentAssets / TotalCurrentLiabilities)
func (b *YearBalanceSheet) CurrentRatio() float32 {
	return float32(b.TotalCurrentAssets()) / float32(b.TotalCurrentLiabilities())
}

// DebtToShareholderEquityRatio (TotalLiabilities / ShareHoldersEquity)
func (b *YearBalanceSheet) DebtToShareholderEquityRatio() float32 {
	return float32(b.TotalLiabilities()) / float32(b.TotalShareholdersEquity())
}
