package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
FileOpen Opens a provided file from the parameter provided.
*/
func FileOpen(filename string) (*os.File, error) {
	fmt.Println("Opening...", filename)
	return os.Open(filename)
}

/*
FileClose Closes a file once read.
*/
func FileClose(file *os.File) {
	fmt.Println("Closing...")
	file.Close()
}

/*
GetFloats Reads a string from a line of a file and converts it to a float64.
*/
func GetFloats(filename string) ([]float64, error) { // return a slice not an array
	var numbers []float64 // nil by default, append treats nil like empty

	file, err := FileOpen(filename) // open the provided filename
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) // convert line by line string to float
		if err != nil {
			return nil, err // return nil, not the slice.
		}
		numbers = append(numbers, number)
	}

	FileClose(file)

	if scanner.Err() != nil { // error during a scan?
		return nil, scanner.Err() // return nil instead of the slice.
	}

	return numbers, nil // finally return the array of numbers and nil error
}