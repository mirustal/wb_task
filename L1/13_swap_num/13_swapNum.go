package swapnum

func SwapNum(a *int,b *int){
	*a, *b = *b, *a
}