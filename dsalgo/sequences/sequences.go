package sequences

func Fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		ret := b
		t := b
		b = a
		a += t
		return ret
	}
}
