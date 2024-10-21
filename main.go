package main
import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)
func calculateAverage(numbers []float64) float64 {
	total := 0.0
	count := len(numbers)
	// Calculate the total sum of numbers
	for _, num := range numbers {
		total += num
	}
	// Calculate the average
	average := total / float64(count)
	return average
}
func calculateStandardDeviation(numbers []float64) float64 {
	mean := calculateAverage(numbers)
	variance := 0.0
	// Calculate the sum of squares of differences from the mean
	for _, num := range numbers {
		diff := num - mean
		variance += diff * diff
	}
	// Calculate the variance and take the square root to get the standard deviation
	variance /= float64(len(numbers))
	stdDev := math.Sqrt(variance)
	return stdDev
}
func calculateMedian(numbers []float64) float64 {
	// Sort the numbers
	sort.Float64s(numbers)
	length := len(numbers)
	if length == 0 {
		return 0 // Handle empty slice
	}
	// Calculate the median based on whether the number of elements is odd or even
	if length%2 == 0 {
		// If the number of elements is even, take the average of the middle two elements
		mid := length / 2
		return (numbers[mid-1] + numbers[mid]) / 2
	} else {
		// If the number of elements is odd, take the middle element
		return numbers[length/2]
	}
}
func calculateVariance(data []float64) float64 {
	// Step 1: Find the mean
	var sum float64
	for _, value := range data {
		sum += value
	}
	mean := sum / float64(len(data))
	// Step 2: Calculate the sum of squared differences
	var squaredDiffs float64
	for _, value := range data {
		squaredDiffs += math.Pow(value-mean, 2)
	}
	// Step 3: Calculate the variance
	variance := squaredDiffs / float64(len(data))
	return variance
}
func checkerr(err error) {
	if err != nil {
		fmt.Println("fi weror ya 7marr 🫰", err)
		return
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("fi weror ya 7marr 🫰")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	checkerr(err)
	defer file.Close()
	myArr, err := ioutil.ReadAll(file)
	checkerr(err)
	myStr := string(myArr)
	mystrs := strings.Split(myStr, "\n")
	numbers := []float64{}
	for _, r := range mystrs {
		c, err := strconv.Atoi(r)
		checkerr(err)
		numbers = append(numbers, float64(c))
	}
	fmt.Println("Average: ", int(math.Round(calculateAverage(numbers))))
	fmt.Println("Median: ", int(math.Round(calculateMedian(numbers))))
	fmt.Println("Variance: ", int(math.Round(calculateVariance(numbers))))
	fmt.Println("Standard Deviation: ", int(math.Round(calculateStandardDeviation(numbers))))
}
