package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("Error Found: ", err)
	}

}
func scanLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
func Tester() {
	filename := os.Args[1]
	url := os.Args[2]
	words, err := scanLines(filename)
	handleError(err)
	for i := 0; i < len(words); i++ {
		req, err1 := http.Get(url + "/" + words[i])
		handleError(err1)
		defer req.Body.Close()
		fmt.Println("Test: ", url+"/"+words[i], "\t\tStatus", req.StatusCode)
		if i == len(words) {
			fmt.Println("All Urls tested")
		}

	}

}
func main() {
	Tester()
}
