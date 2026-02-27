package libraries

import (
	"math/big"
	"testing"
)

func ethToken() ICurrency {
	return NewCurrency(1, "0xETH0000000000000000000000000000000000", 18, "ETH", "Ether")
}

func usdcToken() ICurrency {
	return NewCurrency(1, "0xUSDC00000000000000000000000000000000", 6, "USDC", "US Dollar Coin")
}

func daiToken() ICurrency {
	return NewCurrency(1, "0xDAI0000000000000000000000000000000000", 18, "DAI", "DAI Stablecoin")
}


func TestComputePriceImpact(t *testing.T) {
	ETH := ethToken()
	USDC := usdcToken()

	// midPrice ETH/USDC = 2000
	midPrice := NewPrice(ETH, USDC, big.NewInt(1e18), big.NewInt(2000e6))

	// input: 1 ETH
	input := newCurrencyAmount(ETH, big.NewInt(1e18), big.NewInt(1))

	// actual output: 1990 USDC
	output := newCurrencyAmount(USDC, big.NewInt(1990e6), big.NewInt(1))

	impact, err := ComputePriceImpact(midPrice, input, output)
	if err != nil {
		t.Fatal(err)
	}

	if impact.ToFixed(4) != "0.5000" {
		t.Fatal("expected 0.5000%")
	}

}