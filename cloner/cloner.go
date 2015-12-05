package cloner

import (
	"fmt"
	//"regexp"
	//"net/http"
	"path/filepath"
)

func Start(url string, outDir string) {
	outDir, _ = filepath.Abs(outDir)
	fmt.Printf("Clone book from %s to %s? (yes/no): ", url, outDir)
	var choise string
	fmt.Scanf("%s", &choise)
	if choise == "yes" {
		fmt.Printf("Just kidding :D, we'll handle this function later!\n")
	} else {
		fmt.Printf("User aborted!\n")
	}
}
