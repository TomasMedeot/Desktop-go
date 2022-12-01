package functions

import "strconv"

func Calc(var_1 string, var_2 string, calc string) (rest int) {

	num_1, err := strconv.Atoi(var_1)
	if err != nil {
		num_1 = 0
	}

	num_2, err := strconv.Atoi(var_2)
	if err != nil {
		num_2 = 0
	}

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
