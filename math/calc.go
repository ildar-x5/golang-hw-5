package math

import (
	"github.com/GDG-Cloud-Innopolis/Go-begginners/l3/test"
)


// Calc is a function that return an operation function
func Calc(operation test.CalcOperation) func(a, b float64) float64 {
	var resultFunc func(a, b float64) float64
	switch operation {
	case "+" : 
		resultFunc = func(a, b float64) float64 { return a + b }
	case "-": 
		resultFunc = func(a, b float64) float64 { return a - b }
	case "*" : 
		resultFunc = func(a, b float64) float64 { return a * b }
	case "/" : 
		resultFunc = func(a, b float64) float64 { return a / b }
	}
	return resultFunc
}
