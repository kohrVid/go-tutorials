package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	//"strconv"
	//"strings"
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
      <label for="numbers">Numbers (comma or space-separated):</label><br />
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
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if a, b, c, message, ok := processRequest(request); ok {
			solution := getX(a, b, c)
			fmt.Fprintf(writer, formatAnswer(solution))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (float64, float64, float64, string, bool) {
	var a float64
	var b float64
	var c float64
	/*if slice, found := request.Form["a"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return a, b, c, "'" + field + "' is invalid", false
			} else {
				a = x
			}
		}
	}*/
	if a == 0 || b == 0 || c == 0 {
		return a, b, c, "", false
	}
	return a, b, c, "", true

}

func formatAnswer(solution Quadratics) string {
	return fmt.Sprintf(`<p>%fx<sup>2</sup> + %fx + %f --> x = %f or x = %f</p>`, solution.a, solution.b, solution.c, solution.x1, solution.x2)
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
	fmt.Printf("%f %f %f, %f %f", a, b, c, solution.x1, solution.x2)
	return solution
}

func quadraticFormula(a float64, b float64, c float64) (float64, float64) {
	discriminant := math.Pow(b, 2) - (4.0 * a * c)
	x1 := ((b * -1.0) + math.Sqrt(discriminant)) / (2.0 * a)
	x2 := ((b * -1.0) - math.Sqrt(discriminant)) / (2.0 * a)
	return x1, x2
}
