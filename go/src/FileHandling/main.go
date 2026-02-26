package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile() error {
	_, err := os.Open("non_existent_file.txt")
	if err != nil {
		fmt.Println(err)
		// The %w verb wraps the original error 'err'
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil
}

func main() {
	err := readFile()
	if err != nil {
		fmt.Println(err)
		// You can now check if the error is of a specific type
		if os.IsNotExist(err) {
			fmt.Println("The file does not exist! 📁")
		} else {
			fmt.Println("An unexpected error occurred.")
		}
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("This line will print. The file doesn't exist. ✅")
		} else {
			fmt.Println("Something else went wrong.")
		}

		// You can also unwrap the error to get the original one
		wrappedErr := fmt.Errorf("could not process data: %v", err)
		fmt.Println(wrappedErr)

		// Or check if it is a specific error
		// if e := os.Unwrap(wrappedErr); e != nil {
		// 	fmt.Println("The original error was:", e)
		// }
	}
}
