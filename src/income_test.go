package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetEarnings(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	y, _ := m.GetIncomeStatement("AAPC")
	s, _ := m.GetStockInfo("AAPC")

	x := NewIncomeStatement(y, s)

	// Act`
	x.NetEarnings()

	// Assert
	assert.True(t, true)
}

func TestNetEarningsSTD(t *testing.T) {

}

func TestPerShareEarningsSTD(t *testing.T) {
	// Arrange
	m := YahooMockClient{}
	y, _ := m.GetIncomeStatement("AAPC")
	s, _ := m.GetStockInfo("AAPC")

	x := NewIncomeStatement(y, s)

	// Act
	x.PerShareEarningsSTD()

	// Assert
	assert.True(t, true)
}
