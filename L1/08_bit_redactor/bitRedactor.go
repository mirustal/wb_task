package bitredactor



func BitRedactor(n int64, pos uint8, bit bool) int64 {
	n ^= 1 << pos
    return n
}