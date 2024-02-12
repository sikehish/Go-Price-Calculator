package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() { //It'll iterate as long scanner.Scan() returns true.It returns false when the scan stops, either by reaching the end of the input or an error.
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() //After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil.

	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file content failed")
	}

	file.Close()
	return lines, nil
}

// any and interface{} are the same
func WriteJSON(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	file.Close()
	return nil
}

// NOTE: Both marshal and NewEncoder are used for encoding Go data structures into JSON format. The main difference between them is the way they handle the output.

// marshal is a function from the encoding/json package that takes a Go data structure and returns a byte slice containing the JSON representation of that structure. It is useful when you want to encode the data directly into a byte slice.

// On the other hand, NewEncoder is a type from the same package that provides an interface for writing JSON data to an output stream, such as a file or network connection. It allows you to encode the data directly into the output stream without having to create an intermediate byte slice.

// In the practice project, marshal is used because the goal is to encode the data into a byte slice and then write it to a file using ioutil.WriteFile. In this case, using NewEncoder would require additional steps to write the data to a file.

// To choose between marshal and NewEncoder, consider whether you need to encode the data into a byte slice or directly into an output stream. If you need to write the data to a file or network connection, NewEncoder is a good choice. If you just need the JSON representation as a byte slice, marshal is more suitable.
