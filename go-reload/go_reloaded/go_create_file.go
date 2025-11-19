package go_reloaded

import (
	"fmt"
	"os"
)

func CreatFile(file string, data []byte) {
	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Can not open the file")
		os.Exit(1)
	}
	defer output.Close()

	_, err = output.Write(data)
	if err != nil {
		fmt.Println("Cannot write to the file")
		os.Exit(1)
	}
}
