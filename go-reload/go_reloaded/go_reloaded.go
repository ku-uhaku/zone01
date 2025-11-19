package go_reloaded

import (
	"fmt"
	"os"
)

func Go_reloaded() {
	if len(os.Args) == 3 {
		input := os.Args[1]
		output := os.Args[2]
		content := Read_File(input)
		token := Tokenize(string(content))
		s := ProcessTokens(token)
		CreatFile(output, []byte(s))
	} else {
		fmt.Println("Error in argument")
	}
}
