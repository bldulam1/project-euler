package problem16

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxUINT = 0xFFFFFFFF

// BigNum is a number in string format
type BigNum struct {
	num string
}

func getMaxLen(addends []string) int {
	maxLen := 0

	for _, addend := range addends {
		addendLen := len(addend)
		if maxLen < addendLen {
			maxLen = addendLen
		}
	}

	return maxLen
}

// BigSum returns the sum of an array of string
func BigSum(addends []string) string {

	maxLen := getMaxLen(addends)
	var co uint64
	sum := ""

	for i := 0; i < maxLen; i++ {
		rem := uint64(0) + co
		for _, addend := range addends {
			aIndex := len(addend) - 1 - i
			if aIndex >= 0 {
				rem += uint64(addend[aIndex] - 48)
			}
		}
		co = rem / 10
		sum = strconv.Itoa(int(rem%10)) + sum
	}

	return sum
}

// BigProduct returns the product of two big numbers
func BigProduct(a string, b string, baseZeros int) string {
	num1 := BigNum{a}.Digits(baseZeros)
	num2 := BigNum{b}.Digits(baseZeros)

	base := uint(math.Pow10(baseZeros))
	zeroPadding := strings.Repeat("0", baseZeros)

	addends := make([]string, len(num2))

	for n2i := len(num2) - 1; n2i >= 0; n2i-- {
		carryOver := uint(0)
		tempSum := uint(0)
		addend := ""
		prod := uint(0)
		for n1i := len(num1) - 1; n1i >= 0; n1i-- {
			tempSum = num1[n1i]*num2[n2i] + carryOver

			carryOver = tempSum / base
			prod = tempSum % base

			if prod > 0 {
				prodS := fmt.Sprintf("%d", prod)
				addend = prodS + addend
				if len(prodS) < baseZeros {
					addend = strings.Repeat("0", baseZeros-len(prodS)) + addend
				}
			} else {
				addend = zeroPadding + addend
			}
		}
		addend += strings.Repeat(zeroPadding, len(num2)-n2i-1)
		if carryOver > 0 {
			addend = fmt.Sprintf("%d", carryOver) + addend
		}
		addends[n2i] = addend

	}

	return BigSum(addends)
}

// BigPower returns the product of a number num, n times
func BigPower(num string, n uint) string {
	acc := "1"
	for i := uint(0); i < n; i++ {
		acc = BigProduct(acc, num, 5)
	}
	return acc
}

// Digits returns the digits of the big number
func (bn BigNum) Digits(digLen int) []uint {
	var digits []uint

	isLenNoRemainder := len(bn.num)%digLen > 0

	if isLenNoRemainder {
		digits = make([]uint, len(bn.num)/digLen+1)
	} else {
		digits = make([]uint, len(bn.num)/digLen)
	}

	var digit uint64
	bigNumLen := len(bn.num)

	for i := bigNumLen; i > 0; i -= digLen {
		if i-digLen > 0 {
			digit, _ = strconv.ParseUint(bn.num[i-digLen:i], 10, 64)
		} else {
			digit, _ = strconv.ParseUint(bn.num[0:i], 10, 64)
		}

		if isLenNoRemainder {
			digits[i/digLen] = uint(digit)
		} else {
			digits[i/digLen-1] = uint(digit)
		}

	}

	return digits
}
