package variadicfunc

//variadic func
func Variadicfunc(num ...int) int { // ellipsiss
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}

//limitations
//cannot have multiple variadic parameters
