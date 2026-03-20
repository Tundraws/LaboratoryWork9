package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCalculateSumOfSquares(t *testing.T) {
	const expectedResult = 14
	inputNumbers := []int{1, 2, 3}
	
	actualResult := CalculateSumOfSquares(inputNumbers)
	
	if actualResult != expectedResult {
		t.Errorf("Calculation failed: expected %d, got %d", expectedResult, actualResult)
	}
}

func TestProcessJsonStream(t *testing.T) {
	const validJsonInput = `{"numbers": [1, 2, 3]}`
	const expectedOutputFragment = `{"sum":14}`
	const invalidJsonInput = `{invalid}`

	t.Run("ValidInput", func(t *testing.T) {
		inputSource := strings.NewReader(validJsonInput)
		outputDestination := &bytes.Buffer{}

		if err := ProcessJsonStream(inputSource, outputDestination); err != nil {
			t.Fatalf("ProcessJsonStream failed: %v", err)
		}

		if !strings.Contains(outputDestination.String(), expectedOutputFragment) {
			t.Errorf("Expected %s, got %s", expectedOutputFragment, outputDestination.String())
		}
	})

	t.Run("InvalidInput", func(t *testing.T) {
		inputSource := strings.NewReader(invalidJsonInput)
		outputDestination := &bytes.Buffer{}

		if err := ProcessJsonStream(inputSource, outputDestination); err == nil {
			t.Errorf("Expected error for invalid JSON, but got nil")
		}
	})
}