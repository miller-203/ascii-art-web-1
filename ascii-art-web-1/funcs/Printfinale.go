package funcs

import (
	"os"
	"strings"
)

func Printfinale(userText, choice string) (string, error) {
	results := ""
	var fileContent []byte
	var err error
	fileContent, err = os.ReadFile("src/" + choice + ".txt")
	// fmt.Println(choice)
	if err != nil {
		
		results = "Error: the choice you selected is not found, choose another one"
		return results, nil
		
	}
	fileString := string(fileContent)
	fileString = strings.ReplaceAll(fileString, "\r\n", "\n")
	lines := strings.Split(fileString, "\n")
	// input user splitted by \r\n
	inputText := strings.Split(userText, "\r\n")

	for _, V := range inputText {

		for i := 1; i < 9; i++ {
			for j := 0; j < len(V); j++ {

				asciiVal := (int(V[j]) - 32) * 9
				results += lines[asciiVal+i]
				if j == len(V)-1 {
					results += "\n"
				}

			}
		}
		results += "\n"
	}

	return results, err
}
