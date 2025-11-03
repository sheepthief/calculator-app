package RPN

import (
	"testing"

	"log"

	"reflect"
)

func TestSubtraction(t *testing.T) {

	formula := "4-3"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"-",
	}
	wantNumAcc := []float64{
		4.0,
		3.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 1 {
		t.Errorf("wrong answer")
	}
}

func TestAddition(t *testing.T) {

	formula := "4+3"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"+",
	}
	wantNumAcc := []float64{
		4.0,
		3.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 7 {
		t.Errorf("wrong answer")
	}
}

func TestMultiplication(t *testing.T) {

	formula := "4*3"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"*",
	}
	wantNumAcc := []float64{
		4.0,
		3.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 12 {
		t.Errorf("wrong answer")
	}
}

func TestDivision(t *testing.T) {

	formula := "4/3"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"/",
	}
	wantNumAcc := []float64{
		4.0,
		3.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 4.0/3.0 {
		t.Errorf("wrong answer")
	}
}

func TestPower(t *testing.T) {

	formula := "4^3"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"^",
	}
	wantNumAcc := []float64{
		4.0,
		3.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 64 {
		t.Errorf("wrong answer")
	}
}

func TestComplex(t *testing.T) {

	formula := "3+4*(2-1)"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"",
		"",
		"-",
		"*",
		"+",
	}
	wantNumAcc := []float64{
		3.0,
		4.0,
		2.0,
		1.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 7 {
		t.Errorf("wrong answer")
	}
}

func TestLongNumber(t *testing.T) {

	formula := "30+4"

	opAcc, numAcc, err := Encode(formula)

	log.SetPrefix("Calculator: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	wantOpAcc := []string{

		"",
		"",
		"+",
	}
	wantNumAcc := []float64{
		30.0,
		4.0,
	}

	if !reflect.DeepEqual(wantNumAcc, numAcc) {
		t.Errorf("Fail in numAcc:")
		t.Error(numAcc)
		t.Error(len(formula))

	}
	if !reflect.DeepEqual(wantOpAcc, opAcc) {
		t.Errorf("Fail in opAcc")
		t.Error(opAcc)
		t.Error(len(opAcc))

	}

	result := Solve(opAcc, numAcc)

	if result != 34 {
		t.Errorf("wrong answer")
	}
}
