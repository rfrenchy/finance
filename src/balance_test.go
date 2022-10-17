package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalanceTotalCurrentAssets(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.TotalCurrentAssets())
	assert.NotZero(t, bs.Y2019.TotalCurrentAssets())
	assert.NotZero(t, bs.Y2020.TotalCurrentAssets())
	assert.NotZero(t, bs.Y2021.TotalCurrentAssets())
}

func TestBalanceTotalCurrentLiabilities(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.TotalCurrentLiabilities())
	assert.NotZero(t, bs.Y2019.TotalCurrentLiabilities())
	assert.NotZero(t, bs.Y2020.TotalCurrentLiabilities())
	assert.NotZero(t, bs.Y2021.TotalCurrentLiabilities())
}

func TestBalanceCurrentRatio(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.CurrentRatio())
	assert.NotZero(t, bs.Y2019.CurrentRatio())
	assert.NotZero(t, bs.Y2020.CurrentRatio())
	assert.NotZero(t, bs.Y2021.CurrentRatio())
}

func TestBalanceDebtToShareholderEquityRatio(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.DebtToShareholderEquityRatio())
	assert.NotZero(t, bs.Y2019.DebtToShareholderEquityRatio())
	assert.NotZero(t, bs.Y2020.DebtToShareholderEquityRatio())
	assert.NotZero(t, bs.Y2021.DebtToShareholderEquityRatio())
}

func TestBalanceShortTermDebt(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.ShortTermDebt())
	assert.NotZero(t, bs.Y2019.ShortTermDebt())
	assert.NotZero(t, bs.Y2020.ShortTermDebt())
	assert.NotZero(t, bs.Y2021.ShortTermDebt())
}

func TestBalanceLongTermDebt(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.LongTermDebt())
	assert.NotZero(t, bs.Y2019.LongTermDebt())
	assert.NotZero(t, bs.Y2020.LongTermDebt())
	assert.NotZero(t, bs.Y2021.LongTermDebt())
}

func TestBalanceTotalAssets(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	ybs, _ := m.GetBalanceSheet("")
	bs := NewBalanceSheet(ybs)

	// Act / Assert
	assert.NotZero(t, bs.Y2018.TotalAssets())
	assert.NotZero(t, bs.Y2019.TotalAssets())
	assert.NotZero(t, bs.Y2020.TotalAssets())
	assert.NotZero(t, bs.Y2021.TotalAssets())
}
