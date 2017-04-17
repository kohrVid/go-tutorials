package main

import (
	"fmt"
	"math/big"
)

func main() {
	places := 1000 //handleCommandLine(1000)
	scaledPi := fmt.Sprint(π(places))
	fmt.Printf("3.%s\n", scaledPi[1:])
}

func π(places int) *big.Int {
	digits := big.NewInt(int64(places))
	unity := big.NewInt(0)
	ten := big.NewInt(10)
	exponent := big.NewInt(0)
	unity.Exp(ten, exponent.Add(digits, ten), nil)
	pi := big.NewInt(4)
	left := arccot(big.NewInt(5), unity)
	left.Mul(left, big.NewInt(4))
	right := arccot(big.NewInt(239), unity)
	left.Sub(left, right)
	pi.Mul(pi, left)
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}

func arccot(x, unity *big.Int) *big.Int {
	sum := big.NewInt(0)
	sum.Div(unity, x)
	xpower := big.NewInt(0)
	xpower.Div(unity, x)
	n := big.NewInt(3)
	sign := big.NewInt(-1)
	zero := big.NewInt(0)
	square := big.NewInt(0)
	square.Mul(x, x)
	for {
		xpower.Div(xpower, square)
		term := big.NewInt(0)
		term.Div(xpower, n)
		if term.Cmp(zero) == 0 {
			break
		}
		addend := big.NewInt(0)
		sum.Add(sum, addend.Mul(sign, term))
		sign.Neg(sign)
		n.Add(n, big.NewInt(2))
	}
	return sum
}
