package main

import (
	"fmt"
	"regexp"
	"os"
	"github.com/hoanghai27/bookcloner/cloner"
)

func main() {
	command := os.Args[1]

	if command == "help" || command == "--help" || command == "-h" {
		showHelp();
	} else {
		valid, url, outDir := getParams()
		if valid {
			cloner.Start(url, outDir);
		}
	}
}

func getParams() (bool, string, string) {
	url := os.Args[1]
	outDir := "output";

	if linkValid(url) == false {
		fmt.Printf("URL \"%s\" is not valid.\n", url);
		fmt.Printf("Valid URL must be like this: http://truyenyy.com/truyen/ten-truyen/.\n");
		fmt.Printf("For help, please enter: bookcloner help.\n");
		return false, "", "";
	}

	if len(os.Args) >= 3 {
		outDir = os.Args[2];
	}

	return true, url, outDir;
}

func linkValid(url string) bool {
	valid, _ := regexp.MatchString("http:\\/\\/truyenyy.com\\/truyen\\/(.*?)\\/", url)
	return valid
}

func showHelp() {
	fmt.Printf("Usage: bookcloner link [outputDir]\n")
	fmt.Printf("Params: \n")
	fmt.Printf("- link: Book url from truyenyy.com, example: http://truyenyy.com/truyen/ten-truyen/\n")
	fmt.Printf("- outputDir (optional): Output path\n")
	fmt.Printf("Example:  bookcloner http://truyenyy.com/truyen/ac-ma-phap-tac/ am-ma-phap-tac\n")
}