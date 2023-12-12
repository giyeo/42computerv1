package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Poly struct {
	x0	float64
	x1	float64
	x2	float64
	x3	float64
}

func (p *Poly) parseX(tokens []string) error {
	signal := 1.0
	for index := range tokens {
		if(tokens[index] == "-") {
			signal = -1.0
		}
		if(tokens[index] == "+") {
			signal = 1.0
		}
		f64 := 1.0
		if(len(tokens[index]) > 2 && tokens[index][0] == 'X') {
			if(index > 1 && tokens[index - 1][0] == '*') {
				//check for multiplication factors
				var err error
				f64, err = strconv.ParseFloat(tokens[index - 2], 64)
				if err != nil {
					fmt.Println("Error:", err)
					return err
				}
			}
			if len(tokens[index]) > 2 {
				switch (tokens[index][2]) {
				case '0':
					p.x0 = f64 * signal
				case '1':
					p.x1 = f64 * signal
				case '2':
					p.x2 = f64 * signal
				case '3':
					p.x3 = f64 * signal
				}
			}
		}
	}
	return nil
}

func (p Poly) printPoly() {
	fmt.Printf("x0: %f, x1: %f, x2: %f, x3: %f\n", p.x0, p.x1, p.x2, p.x3)
}

func getPoly(s string) (Poly, error) {
	poly := Poly{}
	tokens := strings.Split(s, " ")
	err := poly.parseX(tokens)
	if err != nil {
		fmt.Println("Error", err)
		return Poly{}, err
	}
	return poly, nil
}

func isNegative(number int) bool {
    return number < 0
}

func getDegreeStr(value float64, degree string) string {
	if (value == 0) {
		return ""
	} else if (value > 0) {
		return "+ " + strconv.FormatFloat(value, 'f', -1, 64) + " * " + degree + " " 
	} else {
		value = value * -1.0
		return "- " + strconv.FormatFloat(value, 'f', -1, 64) + " * " + degree + " "
	}
}

func getPolyReducedForm(poly1 Poly, poly2 Poly) string {
	x0 := poly1.x0 - poly2.x0
	x1 := poly1.x1 - poly2.x1
	x2 := poly1.x2 - poly2.x2
	x3 := poly1.x3 - poly2.x3
	s0, s1, s2, s3 := "X^0", "X^1", "X^2", "X^3"
	res := getDegreeStr(x0, s0) + getDegreeStr(x1, s1) + getDegreeStr(x2, s2) + getDegreeStr(x3, s3)

	if(res[0] == '+') {
		res = res[1:]
	}
	res = res + "= 0"

	return res
}

func givenInput(input string) string {
	sides := strings.Split(input, "=")
	polyInstance, err := getPoly(sides[0])
	polyInstance2, err2 := getPoly(sides[1])
	if(err != nil || err2 != nil) {
		fmt.Println("Error", err)
		return ""
	}
	return getPolyReducedForm(polyInstance, polyInstance2)
	
}

func main() {
	givenInput("5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
}