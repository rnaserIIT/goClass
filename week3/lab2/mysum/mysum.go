package mysum

func Sum(num1, num2 int) string {
	total := num1 + num2
	if total%2 == 0 {
		return ("even")
	} else {
		return ("odd")
	}

}
