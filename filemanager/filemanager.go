package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(ipPath, opPath string) FileManager {
	return FileManager{
		InputFilePath:  ipPath,
		OutputFilePath: opPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	defer file.Close() // It executes at the end, and will execute even in an event where an error occurs.THey are executed even after a return statement. The defer statements are placed on a stack, and they are executed in reverse order (last to first) when the surrounding function returns.

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() { //It'll iterate as long scanner.Scan() returns true.It returns false when the scan stops, either by reaching the end of the input or an error.
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() //After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil.

	if err != nil {
		// file.Close()
		return nil, errors.New("Reading the file content failed")
	}

	// file.Close()
	return lines, nil
}

// any and interface{} are the same
func (fm FileManager) WriteResult(data any) error {

	//Creates a directory if it doesnt exist
	if err := os.MkdirAll("results", 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close() // The defer statements are executed even after a return statement. The defer statements are placed on a stack, and they are executed in reverse order (last to first) when the surrounding function returns.

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	// file.Close()
	return nil

	// //OR The below lines of code (using json.Marshal) can also be used instead of using the above approach
	// jsonData, err := json.Marshal(data) //for the note's fields to be accessible, the fields need to start with capital letters

	// if err != nil {
	// 	return err
	// }
	// return os.WriteFile(fm.OutputFilePath, jsonData, 0644)
}

// NOTE: Both marshal and NewEncoder are used for encoding Go data structures into JSON format. The main difference between them is the way they handle the output.

// marshal is a function from the encoding/json package that takes a Go data structure and returns a byte slice containing the JSON representation of that structure. It is useful when you want to encode the data directly into a byte slice.

// On the other hand, NewEncoder is a type from the same package that provides an interface for writing JSON data to an output stream, such as a file or network connection. It allows you to encode the data directly into the output stream without having to create an intermediate byte slice.

// In the practice project, marshal is used because the goal is to encode the data into a byte slice and then write it to a file using ioutil.WriteFile. In this case, using NewEncoder would require additional steps to write the data to a file.

// To choose between marshal and NewEncoder, consider whether you need to encode the data into a byte slice or directly into an output stream. If you need to write the data to a file or network connection, NewEncoder is a good choice. If you just need the JSON representation as a byte slice, marshal is more suitable.
