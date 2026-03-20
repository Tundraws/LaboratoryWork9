package main

import (
	"encoding/json"
	"io"
	"os"
)

type CalculationInput struct {
	Numbers []int `json:"numbers"`
}

type CalculationOutput struct {
	Sum int `json:"sum"`
}

func CalculateSumOfSquares(numbers []int) int {
	totalSum := 0
	for _, currentNumber := range numbers {
		totalSum += currentNumber * currentNumber
	}
	return totalSum
}

func ProcessJsonStream(inputSource io.Reader, outputDestination io.Writer) error {
	var input CalculationInput
	if err := json.NewDecoder(inputSource).Decode(&input); err != nil {
		return err
	}

	result := CalculateSumOfSquares(input.Numbers)
	output := CalculationOutput{Sum: result}

	return json.NewEncoder(outputDestination).Encode(output)
}

func main() {
	const errorExitCode = 1
	if err := ProcessJsonStream(os.Stdin, os.Stdout); err != nil {
		os.Exit(errorExitCode)
	}
}