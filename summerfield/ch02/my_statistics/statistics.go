package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
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
    <title>Statistics</title>
  </head>
  <body>
    <p>
      Computs basic statistics for a given list of numbers
    </p>`
	form = `<form action="/" method="POST">
      <label for="numbers">Numbers (comma or space-separated):</label><br />
      <input type="text" name="numbers" size="30"><br />
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
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprintf(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
  <tr>
    <th colspan="2">Results</th>
  </tr>
  <tr>
    <td>Numbers</td>
    <td>%v</td>
  </tr>
  <tr>
    <td>Count</td>
    <td>%d</td>
  </tr>
  <tr>
    <td>Mean</td>
    <td>%f</td>
  </tr>
  <tr>
    <td>Median</td>
    <td>%f</td>
  </tr>
  <tr>
    <td>Mode</td>
    <td>%f</td>
  </tr>
  <tr>
    <td>Std. Dev.</td>
    <td>%f</td>
  </tr>
  </table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, mode(stats.numbers), stats.stdDev)
}

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	mode    float64
	stdDev  float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = mean(numbers)
	stats.median = median(numbers)
	sort.Float64s(mode(numbers))
	stats.stdDev = stdDev(numbers)
	return stats
}

func stdDev(numbers []float64) float64 {
	mean := mean(numbers)
	n := len(numbers)
	xMuSquared := make([]float64, n)
	for _, x := range numbers {
		xMu := x - mean
		xMuSquared = append(xMuSquared, math.Pow(xMu, 2))
	}
	ss := sum(xMuSquared)
	result := math.Sqrt(ss / float64(n-1))
	return result
}

func mean(numbers []float64) (total float64) {
	result := sum(numbers) / float64(len(numbers))
	return result
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle+1]) / 2
	}
	return result
}

func mode(numbers []float64) []float64 {
	var countHash = make(map[float64]float64)
	for _, x := range numbers {
		countHash[x] += 1
	}
	max := numbers[0]
	for k, v := range countHash {
		if v > countHash[max] {
			max = k
		}
	}
	results := make([]float64, len(numbers))
	results = append(results, max)
	for k, v := range countHash {
		if countHash[max] == v && k != max {
			results = append(results, k)
		}
	}

	for i := 0; i < len(results); i++ {
		newArr := make([]float64, len(results)-1)
		if results[i] == 0 {
			newArr = append(results[i:])
			results = newArr
		} else if results[i+1] != 0 {
			newArr = append(results[i:len(results)])
			results = newArr
		}
	}
	return results
}
