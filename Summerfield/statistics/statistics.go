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
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers   []float64
	mean      float64
	median    float64
	deviation float64
	modalMean []float64
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

//homePage handles the requests to home page of web app.
func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

//processRequest reads and returns data from html-form.
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
		return numbers, "", false
	}
	return numbers, "", true
}

//formatStats returns string whitch contain the formatted result.
func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%f</td></tr>
<tr><td>Std. Dev.</td><td>%v</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.deviation, stats.modalMean)
}

//getStats takes a slice with numbers and fills in the return value with the calculation results
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.deviation = stdDev(numbers)
	stats.modalMean = modalMean(numbers)
	return stats
}

//sum returns sum of all number in input-slice.
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

//median returns median of a slice numbers.
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

//maximum returns maximum of a slice numbers.
func maximum(numbers []int) int {
	max := int(math.Inf(-1))
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

//modalMean returns the mode of a slice numbers.
func modalMean(numbers []float64) []float64 {
	counts := make(map[float64]int)
	result := make([]float64, 0)
	coincidence := make([]int, 0)
	for _, num := range numbers {
		counts[num]++
	}
	for _, v := range counts {
		coincidence = append(coincidence, v)
	}
	if len(coincidence) != len(numbers) {
		for k, v := range counts {
			if v == maximum(coincidence) {
				result = append(result, k)
			}
		}
		return result
	}
	return nil
}

//stdDev returns the standard deviation of a slice numbers.
func stdDev(numbers []float64) float64 {
	var deviation float64
	median := sum(numbers) / float64(len(numbers))
	sum := 0.0
	for _, num := range numbers {
		sum += math.Pow(num-median, 2)
	}
	deviation = math.Sqrt(sum / float64((len(numbers) - 1)))
	return deviation
}
