package RPN

import (
	"errors"

	"strconv"

	"example.com/addition"

	"example.com/subtraction"

	"example.com/multiplication"

	"example.com/division"
)

func Encode(calc string) ([]string, []float64, error) {

	var opAcc []string
	var numAcc []float64

	//Allows for temporaryily storing operator
	var tempOpAcc []string
	var tempRanking []float64

	longNumbAcc := ""

	var i float64
	i = 0.0

	for j, num := range calc {

		//Check if the string contains a number
		numIsANumber := 47 < float64(num) && float64(num) < 58

		//Check if we are using legal operand
		// look at ascii table for help
		opIsLegal := 39 < num && num < 44 || num == 94 || num == 45 || num == 47

		if numIsANumber {
			//String number
			number := string(num)
			longNumbAcc += number

			// if we are on the last part of the formula, we need to add the number, a formula either ends on a number
			// or a closed parenthesis
			if len(calc)-1 <= j {
				number, _ := strconv.Atoi(longNumbAcc)
				numAcc = append(numAcc, float64(number))
				longNumbAcc = ""
				opAcc = append(opAcc, "")
				break
			}
			continue
		}

		op := string(num)
		// We need to prioritize the part of the function inside ()
		// We don't want the precedence to be equal to any from earlier rounds
		if op == "(" {
			i += 3.0
			continue
		}
		if op == ")" {
			i -= 3.0
			// if we are on the last part of the formula, we need to add the number
			if len(calc)-1 <= j {
				number, _ := strconv.Atoi(longNumbAcc)
				numAcc = append(numAcc, float64(number))
				longNumbAcc = ""
				opAcc = append(opAcc, "")
				break
			}

			continue
		}

		//if i goes below zero, we know that the () has been messed up
		if i < 0 {
			var emptyOp []string
			var emptyNum []float64
			return emptyOp, emptyNum, errors.New("missing ()")
		}

		// If long number accumulate isn't we know we are dealing with a number
		// If there is anything in the number accumulater, we know that we either
		// Have an operand, or we are done and need to add the last number
		if longNumbAcc != "" {
			number, _ := strconv.Atoi(longNumbAcc)
			numAcc = append(numAcc, float64(number))
			longNumbAcc = ""
			opAcc = append(opAcc, "")

			if opIsLegal {

				var legalOps = map[string]float64{
					"^": 0.0,
					"*": 1.0,
					"/": 1.0,
					"+": 2.0,
					"-": 2.0,
				}

				op := string(num)

				if len(tempOpAcc) > 0 {
					// We neeed to see ranking of top of stack, to compare to new
					//priorityOp := tempOpAcc[len(tempOpAcc)-1]
					rankingPriority := tempRanking[len(tempRanking)-1]

					var rankingNew float64 = legalOps[op] - i

					// Based on the precedence of the top of the temp stack, we need to do different things
					// If the new operation has lower we know it is more important
					// If the new operation has the same we need to pop it onto the permanent stack, and those
					// on the stack
					// If the new operation has higher

					if rankingPriority > rankingNew {
						// The new prospect is better, so we need to add it and it current ranking for the round
						// Can change thanks to ()
						tempOpAcc = append(tempOpAcc, op)
						tempRanking = append(tempRanking, rankingNew)
						continue

					} else if rankingPriority < rankingNew {
						// We need the one with smallest precedence on the top
						popPriOp := tempOpAcc[len(tempOpAcc)-1]
						popPriR := tempRanking[len(tempRanking)-1]
						tempOpAcc = tempOpAcc[:len(tempOpAcc)-1]

						// We need to reverse the ranking
						tempOpAcc = append(tempOpAcc, op)
						tempOpAcc = append(tempOpAcc, popPriOp)

						tempRanking = append(tempRanking, rankingNew)
						tempRanking = append(tempRanking, popPriR)

						continue
					} else {

						// If they have the same precedence we need to see if there are more on the stack
						popPri := tempOpAcc[:len(tempOpAcc)-1]
						opAcc = append(opAcc, popPri[0])

						continue

					}

				} else {
					//If the list is empty, we just push the operator onto the temporary stack
					tempOpAcc = append(tempOpAcc, op)

					tempRanking = append(tempRanking, legalOps[op]-i)

				}
			}
			continue

		}

		// If it is under 13 we know they are end of sentence character
		if 13 < num {
			break
		}

	}

	//If i is too big we know we are missing )
	if i > 0 {
		var emptyOp []string
		var emptyNum []float64
		return emptyOp, emptyNum, errors.New("missing ())")
	}

	//We need to check if the temporary stack is empty, otherwise we need to empty it

	if len(tempOpAcc) > 0 {

		//We need to reverse the list
		remainingOps := len(tempOpAcc)
		for j := 0; j < remainingOps; j++ {

			popPri := tempOpAcc[len(tempOpAcc)-1]
			tempOpAcc = tempOpAcc[:len(tempOpAcc)-1]

			opAcc = append(opAcc, popPri)
		}
		opAcc = append(opAcc, tempOpAcc...)
	}
	return opAcc, numAcc, nil

}

func Solve(opAcc []string, numAcc []float64) float64 {
	result := 0.0
	var tempNumAcc []float64

	for _, char := range opAcc {

		// The way opACC is build, is that numbers are ""
		if char == "" {
			//We extraxt the first element
			var number float64
			number, numAcc = numAcc[0], numAcc[1:]
			tempNumAcc = append(tempNumAcc, number)

		} else {
			var number1 float64
			var number2 float64
			//We now know that it is an operator, so we pop the last two numbers
			number2, tempNumAcc = tempNumAcc[len(tempNumAcc)-1], tempNumAcc[:len(tempNumAcc)-1]
			number1, tempNumAcc = tempNumAcc[len(tempNumAcc)-1], tempNumAcc[:len(tempNumAcc)-1]

			if char == "^" {
				tempResult := number1
				intNumber2 := int(number2)

				for i := 0; i < intNumber2-1; i++ {
					tempResult = multiplication.Mul(tempResult, number1)
				}
				tempNumAcc = append(tempNumAcc, tempResult)
			}
			if char == "/" {
				subResult := division.Div(number1, number2)
				tempNumAcc = append(tempNumAcc, subResult)

			}
			if char == "*" {
				subResult := multiplication.Mul(number1, number2)
				tempNumAcc = append(tempNumAcc, subResult)

			}
			if char == "+" {
				subResult := addition.Add(number1, number2)
				tempNumAcc = append(tempNumAcc, subResult)
			}
			if char == "-" {

				subResult := subtraction.Sub(number1, number2)
				tempNumAcc = append(tempNumAcc, subResult)

			}

		}
	}
	result = tempNumAcc[len(tempNumAcc)-1]
	return result
}
