package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan_lines(filename string) ([]string, error) {
	file, _ := os.Open(filename)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
func main() {
	lines, _ := scan_lines("directory.txt")

	fmt.Println(lines)

}
