package problem17

import "fmt"

// Numbers is a map of numbers to strings
var Numbers map[uint16]string = map[uint16]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// NumberToWord converts number into words
func NumberToWord(num uint16) string {
	var word string

	// thou := num / 1000
	// hund := (num - thou*1000) / 100
	// tens := (num - thou*1000 + hund*100) / 10
	ones := num % 10
	tens := (num % 100) / 10
	hund := (num % 1000) / 100
	thou := num / 1000

	if thou > 0 {
		if hund > 0 || tens > 0 || ones > 0 {
			word = fmt.Sprintf("%s thousand %s", Numbers[num/1000], NumberToWord(num%1000))
		} else {
			word = fmt.Sprintf("%s thousand", Numbers[num/1000])
		}
	} else if hund > 0 {
		if tens > 0 || ones > 0 {
			word = fmt.Sprintf("%s hundred and %s", Numbers[num/100], NumberToWord(num%100))
		} else {
			word = fmt.Sprintf("%s hundred", Numbers[num/100])
		}
	} else if num < 20 || num%10 == 0 {
		word = Numbers[num]
	} else if tens > 1 {
		word = Numbers[tens*10] + "-" + Numbers[ones]
	} else {
		panic(fmt.Sprintf("cannot handle %d", num))
	}

	return word
}
