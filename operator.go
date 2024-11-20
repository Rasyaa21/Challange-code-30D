package main

import (
	"fmt"
	"math"
)

func solve(meal_cost float64, tip_percent int32, tax_percent int32) int {
	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100

	total := int(math.Round(meal_cost + tip + tax))
	fmt.Println(total)
	return total
}

func main() {
	solve(15, 20, 9)
}
