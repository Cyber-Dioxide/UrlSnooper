package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
)

//var colorReset = "\033[0m"
//var colorYellow = "\033[33m"
//var colorWhite = "\033[37m"

func clear() {
	exec.Command("cls").Run()

}

func Banner() {
	logo := " ██████████   █████ ███████████   ███████████ \n░░███░░░░███ ░░███ ░░███░░░░░███ ░░███░░░░░███\n ░███   ░░███ ░███  ░███    ░███  ░███    ░███\n ░███    ░███ ░███  ░██████████   ░██████████ \n ░███    ░███ ░███  ░███░░░░░███  ░███░░░░░███\n ░███    ███  ░███  ░███    ░███  ░███    ░███\n ██████████   █████ █████   █████ ███████████ \n░░░░░░░░░░   ░░░░░ ░░░░░   ░░░░░ ░░░░░░░░░░░  \n                                              \n                                              \n       CODED BY: [ S A A D - K H A N ]\n\t\t[C Y B E R D I O X I D E]\n                                "
	//fmt.Println(logo, colorReset)
	color.Yellow(logo)

}
func readerFile(filename string) []string {
	file, err := os.Open(filename)
	HandleError(err)

	defer file.Close()
	read := bufio.NewScanner(file)
	read.Split(bufio.ScanLines)
	var passwords []string
	for read.Scan() {
		passwords = append(passwords, read.Text())
	}

	return passwords
}

func HandleError(err error) {
	if err != nil {
		color.Yellow(err.Error())
	}
}

func TestUrls(dir []string, url string) {
	var temp string
	valid := 0
	for _, j := range dir {
		temp = url + "/" + j
		req, err := http.Get(temp)
		HandleError(err)

		if req.StatusCode == 404 {
			continue
		} else {
			valid += 1
			color.Yellow("Testing: ")
			color.Cyan(temp)
			color.Yellow("\t\tStatus: [ ")
			color.Cyan(strconv.Itoa(req.StatusCode))
			color.Yellow(" ]\n")

			//fmt.Println("Testing:", temp, color.FgYellow, "\t\tStatus: [", color.FgCyan, req.StatusCode, color.FgYellow, "]")

		}
		defer req.Body.Close()
	}
	color.Yellow("Urls Tested: ")
	color.Cyan(strconv.Itoa(len(dir)))
	color.Yellow("\t\tValid Urls: ")
	color.Cyan(strconv.Itoa(valid) + "\n")
	//fmt.Println(color.FgYellow, "Urls Tested: ", color.FgWhite, len(dir), color.FgYellow, "\t\tValid Urls:", color.FgWhite, valid)
}

func main() {
	clear()
	Banner()
	fmt.Print("\n")
	var enter byte
	if len(os.Args) != 3 {
		color.White("Usage: UrlSnooper <wordlist> <url>\n")
		//fmt.Println("Checking color status")
		color.White("Press Enter to Exit...")
		fmt.Scanf("%s", enter)
		os.Exit(0)
	}
	url := os.Args[2]
	wordlist := os.Args[1]
	splinted := readerFile(wordlist)
	TestUrls(splinted, url)

	color.White("Press Enter to Exit...")
	fmt.Scanf("%s", enter)

}
