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

func getDegreeStr(value float64, degree int, hdegree int) string {
	strDegree := []string{"X^0", "X^1", "X^2", "X^3"}
	if (value == 0 && degree > hdegree) {
		return ""
	} else if (value < 0){
		value = value * -1.0
		return "- " + strconv.FormatFloat(value, 'f', -1, 64) + " * " + strDegree[degree] + " "
	} else {
		return "+ " + strconv.FormatFloat(value, 'f', -1, 64) + " * " + strDegree[degree] + " "
	}
}

func getDegree(x0 float64, x1 float64, x2 float64, x3 float64) int {
	if x3 != 0 {
		return 3
	} else if x2 != 0 {
		return 2
	} else if x1 != 0 {
		return 1
	}
	return 0
}

func getPolyReducedForm(poly1 Poly, poly2 Poly) (string, int) {
	x0 := poly1.x0 - poly2.x0
	x1 := poly1.x1 - poly2.x1
	x2 := poly1.x2 - poly2.x2
	x3 := poly1.x3 - poly2.x3

	degree := getDegree(x0, x1, x2, x3)
	res := getDegreeStr(x0, 0, degree)+
		getDegreeStr(x1, 1, degree) +
		getDegreeStr(x2, 2, degree) +
		getDegreeStr(x3, 3, degree)
	
	if res == "" {
        return "0 = 0", 0
    }

	if(res[0] == '+') {
		res = res[2:]
	}
	return res + "= 0", degree
}

func givenInput(input string) (string, int) {
	sides := strings.Split(input, "=")
	polyInstance, err := getPoly(sides[0])
	polyInstance2, err2 := getPoly(sides[1])
	if(err != nil || err2 != nil) {
		fmt.Println("Error", err)
		return "", -1
	}
	return getPolyReducedForm(polyInstance, polyInstance2)
}

func main() {
	println("Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0")
	reduced, degree := givenInput("5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	fmt.Printf("Reduced form: %s\nPolynomial degree: %d\n", reduced, degree)

	println("Reduced form: 1 * X^0 + 4 * X^1 = 0")
	reduced, degree = givenInput("5 * X^0 + 4 * X^1 = 4 * X^0")
	fmt.Printf("Reduced form: %s\nPolynomial degree: %d\n", reduced, degree)

	println("Reduced form: 5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0")
	reduced, degree = givenInput("8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0")
	fmt.Printf("Reduced form: %s\nPolynomial degree: %d\n", reduced, degree)
}