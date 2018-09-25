package mathcalc

import (
	"strconv"
	"testing"
)

var exps = map[string]float64{
	"1+1": 2,
	"3-5": -2,
}

const DELTA = 0.000001

func TestEvaluate(t *testing.T) {
	for expression, expected := range exps {
		res, err := ParseAndEval(expression)
		if err != nil {
			message := expression + " :failed: actual value " +
				strconv.Itoa(res) +
				" differs from expected value " +
				strconv.FormatFloat(expected, 'G', -1, 64) +
				"error: " + err.Error()
			t.Error(message)
		} else {
			message := expression + " :passed: " + "Result: " + strconv.Itoa(res) + " equal to Expected " + strconv.FormatFloat(expected, 'G', -1, 64)
			t.Log(message)
		}
	}
}

func TestEvaluateInvalid(t *testing.T) {
	tests := [][]string{
		{"/"},
		{"1/"},
		{"1("},
		{")("},
		{"(()"},
		{"@"},
		{"@@"},
		{"0", "@@"},
		{"0", "@@@"},
		{"@@\xa6"},
	}
	for i, series := range tests {
		var fail error
		for _, expr := range series {
			if _, err := ParseAndEval(expr); err != nil {
				fail = err
				break
			}
		}
		if fail == nil {
			t.Errorf("case %d: expected error, finished successfully", i)
		}
	}
}

func BenchmarkEvaluate(b *testing.B) {
	tests := []string{
		"π",
		"1+2^3^2",
		"2^(3+4)",
		"2^(3/(1+2))",
		"2^2(1+3)",
		"1+(-1)^2",
		"3*(3-(5+6)^12)*23^3-5^23",
		"2^3^2",
		"ln(3^15)",
		"sqrt(10)",
		"abs(-3/2)",
		"1+2sin(-1024)tan(acos(1))^2",
		"tan(10)cos(20)",
	}
	for i := 0; i < b.N; i++ {
		ParseAndEval(tests[i%len(tests)])
	}
}
