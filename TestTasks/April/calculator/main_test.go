package main

import (
	"regexp"
	"strings"
	"testing"
)

type testData struct {
	expr string
	want float64
}

func TestCalculateExpr(t *testing.T) {
	tests := []testData{
		testData{expr: "1+1+1+1-1-1-1-1", want: 0},
		testData{expr: "100*2-111+1-1/1+1*1", want: 90},
		testData{expr: "1*2*3*4*5*6/6", want: 120},
		testData{expr: "3*2-3/3+5+5-5*5+3*5", want: 5},
		testData{expr: "2*-2", want: -4},
		testData{expr: "2+-2", want: 0},
		testData{expr: "2/-2", want: -1},
		testData{expr: "2--2", want: 4},
		testData{expr: "(2-2)*2", want: 0},
		testData{expr: "1+2*(3+4/2-(1+2))*2+1", want: 10},
	}

	re := regexp.MustCompile(`[+-]?\d+(\.\d+)?|[(|)]|[+|-|*|\/?]|[+|-]`)

	for _, test := range tests {
		expression := test.expr
		expr := strings.Replace(expression, " ", "+", -1)

		//Trick for accounting unary operators.
		signs := regexp.MustCompile(`[\d|\)][\-|+]\d`)
		for {
			changeSigns := signs.FindAllString(expr, -1)
			if len(changeSigns) > 0 {
				for _, s := range changeSigns {
					expr = strings.Replace(expr, s, s[:len(s)-1]+"+"+s[len(s)-1:], -1)
				}
			} else {
				break
			}
		}
		got := CalculateExpr(re.FindAllString(expr, -1))
		if got != test.want {
			t.Errorf("CalculateExpr(%#v) = \"%f\", want \"%f\"", test.expr, got, test.want)
		}
	}
}
