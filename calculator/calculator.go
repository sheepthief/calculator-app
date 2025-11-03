package calculator

import (
	"example.com/RPN"

	"strconv"

	"log"
)

func Solve(formula string) string {

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	opAcc, numAcc, err := RPN.Encode(formula)

	if err != nil {
		log.Fatal(err)
	}

	result := RPN.Solve(opAcc, numAcc)
	strResult := strconv.FormatFloat(result, 'g', 5, 64)

	return strResult
}
