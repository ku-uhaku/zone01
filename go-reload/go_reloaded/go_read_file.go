package go_reloaded

import (
	"fmt"
	"os"
)

func Read_File(file string) []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error Opning the file")
		os.Exit(1)
	}
	return data
}
