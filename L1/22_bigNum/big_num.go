package bignum

import (
    "fmt"
    "math/big"
)

func BigSum(a *big.Int, b *big.Int) {
	sum := big.NewInt(0).Add(a, b)
	fmt.Printf("Sum %s + %s = %s\n",a ,b , sum)
}

func BigMul(a *big.Int, b *big.Int) {
	mul := big.NewInt(0).Mul(a, b)
	fmt.Printf("Multiplex %s * %s = %s\n",a, b, mul)
}

func BigSub(a *big.Int, b *big.Int) {
	sub := big.NewInt(0).Sub(a, b)
	fmt.Printf("Sub %s - %s = %s\n",a, b, sub)
}

func BigDiv(a *big.Int, b *big.Int) {
    div := big.NewInt(0)
    if b.Cmp(big.NewInt(0)) != 0 { // Проверка, что b не равно 0
        div.Div(a, b)
    } else {
        fmt.Println("Divided by zero")
    }
	fmt.Printf("Div %s / %s = %s\n",a,b, div)
}