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
	for index := range tokens {
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
					p.x0 = f64
				case '1':
					p.x1 = f64
				case '2':
					p.x2 = f64
				case '3':
					p.x3 = f64
				}
			}
		}
	}
	return nil
}

func main() {
	input := "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
	tokens := strings.Split(input, " ")
	polyn := Poly{}
	err := polyn.parseX(tokens)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("x0 %f, x1 %f, x2 %f, x3 %f", polyn.x0, polyn.x1, polyn.x2, polyn.x3)

}