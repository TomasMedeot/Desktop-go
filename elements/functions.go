package functions

func Calc(num_1 int, num_2 int, calc string) (rest int) {
	switch calc {
	case "mul":
		rest = num_1 * num_2
	case "sum":
		rest = num_1 + num_2
	case "res":
		rest = num_1 - num_2
	case "div":
		rest = num_1 / num_2
	}
	return rest
}
