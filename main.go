package main

import (
	"fmt"
	"regexp"
	"os"
	"github.com/hoanghai27/bookcloner/cloner"
)

// main function will be call when program start
func main() {
	command := os.Args[1]

	if command == "help" || command == "--help" || command == "-h" {
		showHelp();
	} else {
		valid, url, outFile := getParams()
		if valid {
			cloner.Start(url, outFile);
		}
	}
}

// getParams get all ordered params from os.Args
// It returns 3 values: is all params is valid?, the url and the output path
func getParams() (bool, string, string) {
	url := os.Args[1]
	outFile := "book.html";

	if linkValid(url) == false {
		fmt.Printf("URL \"%s\" is not valid.\n", url);
		fmt.Printf("Valid URL must be like this: http://truyenyy.com/truyen/ten-truyen/.\n");
		fmt.Printf("For help, please enter: bookcloner help.\n");
		return false, "", "";
	}

	if len(os.Args) >= 3 {
		outFile = os.Args[2];
	}

	return true, url, outFile;
}

// linkValid check url provide form user is valid or not
// It returns true if url is valid or false if it is invalid
func linkValid(url string) bool {
	valid, _ := regexp.MatchString("http:\\/\\/truyenyy.com\\/truyen\\/(.*?)\\/", url)
	return valid
}

// showHelp show help text
func showHelp() {
	fmt.Printf("Usage: bookcloner link [outputFile]\n")
	fmt.Printf("Params: \n")
	fmt.Printf("- link: Book url from truyenyy.com, example: http://truyenyy.com/truyen/ten-truyen/\n")
	fmt.Printf("- outputFile (optional): Output path\n")
	fmt.Printf("Example:  bookcloner http://truyenyy.com/truyen/ac-ma-phap-tac/ am-ma-phap-tac.html\n")
}