package numbertowords

func Rupees(amount float64) string {

	num := int(amount)

	if num == 0 {
		return "Zero Rupees Only"
	}

	return convert(num) + " Rupees Only"
}

func convert(num int) string {

	ones := []string{
		"",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
	}

	tens := []string{
		"",
		"",
		"Twenty",
		"Thirty",
		"Forty",
		"Fifty",
		"Sixty",
		"Seventy",
		"Eighty",
		"Ninety",
	}

	switch {

	case num < 20:
		return ones[num]

	case num < 100:
		return tens[num/10] + " " + ones[num%10]

	case num < 1000:
		return ones[num/100] +
			" Hundred " +
			convert(num%100)

	case num < 100000:
		return convert(num/1000) +
			" Thousand " +
			convert(num%1000)

	case num < 10000000:
		return convert(num/100000) +
			" Lakh " +
			convert(num%100000)

	default:
		return convert(num/10000000) +
			" Crore " +
			convert(num%10000000)
	}
}
