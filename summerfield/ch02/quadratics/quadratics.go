package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <style>
      .error{
	color: #ff0000;
      }
    </style>
    <title>Quadratic Equations</title>
  </head>
  <body>
    <p>
      Solves quadratic equations
    </p>`
	form = `<form action="/" method="POST">
      <label for="numbers">Equation:</label><br />
      <input type="text" name="a" size="3">x<sup>2</sup> + 
      <input type="text" name="b" size="3">x + 
      <input type="text" name="c" size="3"> = 0<br />
      <input type="submit" value="Calculate">
    </form>`
	pageBottom = `</body>
  </html>`
	anError = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	var numbers = make(map[string]string)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		i := 0
		for k, v := range request.Form {
			if _, message, ok := processRequest(k, request); ok {
				numbers[k] = strings.Join(v, "")
			} else if message != "" {
				fmt.Fprintf(writer, anError, message)
				os.Exit(-1)
			}
			i += 1
		}
	}
	a, err := strconv.ParseFloat(numbers["a"], 64)
	b, err := strconv.ParseFloat(numbers["b"], 64)
	c, err := strconv.ParseFloat(numbers["c"], 64)
	solution := getX(a, b, c)
	fmt.Fprintf(writer, formatAnswer(solution))
	fmt.Fprint(writer, pageBottom)
}

func processRequest(field string, request *http.Request) (float64, string, bool) {
	var val float64
	if slice, found := request.Form[field]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return val, "'" + field + "' is invalid", false
			} else {
				val = x
			}
		}
	}
	if val == 0 {
		return val, "", false
	}
	return val, "", true

}

func formatAnswer(solution Quadratics) string {
	return fmt.Sprintf(`<p>%.3fx<sup>2</sup> + %.3fx + %.3f --> x = %.3f or x = %.3f</p>`, solution.a, solution.b, solution.c, solution.x1, solution.x2)
}

type Quadratics struct {
	a  float64
	b  float64
	c  float64
	x1 float64
	x2 float64
}

func getX(a float64, b float64, c float64) (solution Quadratics) {
	solution.a = a
	solution.b = b
	solution.c = c
	solution.x1, solution.x2 = quadraticFormula(a, b, c)
	return solution
}

func quadraticFormula(a float64, b float64, c float64) (float64, float64) {
	discriminant := math.Pow(b, 2) - (4.0 * a * c)
	x1 := ((b * -1.0) + math.Sqrt(discriminant)) / (2.0 * a)
	x2 := ((b * -1.0) - math.Sqrt(discriminant)) / (2.0 * a)
	return x1, x2
}
