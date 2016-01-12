package main

import (
	"fmt"
	"github.com/hoanghai27/bookcloner/thichdoctruyen"
	"github.com/hoanghai27/bookcloner/truyenyy"
	"os"
	"regexp"
)

// main function will be call when program start
func main() {
	if len(os.Args) < 2 {
		showHelp()
	} else {
		command := os.Args[1]

		if command == "help" || command == "--help" || command == "-h" {
			showHelp()
		} else {
			run()
		}
	}
}

// getParams get all ordered params from os.Args
// It returns 3 values: is all params is valid?, the url and the output path
func run() {
	url := os.Args[1]
	outFile := "book.html"

	if len(os.Args) >= 3 {
		outFile = os.Args[2]
	}

	switch getCloner(url) {
	case "truyenyy":
		truyenyy.Start(url, outFile)
		break
	case "thichdoctruyen":
		thichdoctruyen.Start(url, outFile)
		break
	default:
		fmt.Printf("URL \"%s\" is not valid.\n", url)
		fmt.Printf("Valid URL must be like this: http://truyenyy.com/truyen/ten-truyen/.\n")
		fmt.Printf("For help, please enter: bookcloner help.\n")
		break
	}
}

// linkValid check url provide form user is valid or not
// It returns cloner name if url is valid or empty string if it is invalid
func getCloner(url string) string {
	if valid, _ := regexp.MatchString("http:\\/\\/truyenyy.com\\/truyen\\/(.*?)\\/", url); valid {
		return "truyenyy"
	} else if valid, _ := regexp.MatchString(`http:\/\/thichdoctruyen\.com\/doc-truyen\/.*-[\d]+`, url); valid {
		return "thichdoctruyen"
	}
	return ""
}

// showHelp show help text
func showHelp() {
	fmt.Printf("Usage: bookcloner link [outputFile]\n")
	fmt.Printf("Params: \n")
	fmt.Printf("- link: Book url from website, example: http://truyenyy.com/truyen/ten-truyen/\n")
	fmt.Printf("- outputFile (optional): Output path\n")
	fmt.Printf("Example:  bookcloner http://truyenyy.com/truyen/ac-ma-phap-tac/ am-ma-phap-tac.html\n")
}
