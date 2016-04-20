package main

import "math"

// Retire returns the amount needed in order to retire
func Retire(currAge, retireAge, deathAge, currIncome int, inflation, rateOfReturn, rateOfSpend float64) float64 {
	var fVArr []float64

	for i := retireAge; i <= deathAge; i++ {
		fVArr = append(fVArr, futureValue(float64(currIncome)*rateOfSpend, inflation, i-currAge))
	}
	var total = 0.0
	for _, num := range fVArr {
		total += num
	}
	return total
}

func futureValue(principal, inflation float64, years int) (result float64) {
	return principal * math.Pow(1+inflation, float64(years))
}

func presentValue(principal, inflation float64, years int) (result float64) {
	return principal / math.Pow(1+inflation, float64(years))
}
