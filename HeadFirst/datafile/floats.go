//Package datafile is used to read some info from files.
package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//openFile opens file with name geting from command line.
func openFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

// GetFloats reads float64 value from every file string.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := openFile(fileName)
	if err != nil {
		return nil, err
	}
	defer closeFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
