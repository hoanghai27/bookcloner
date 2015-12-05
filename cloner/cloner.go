package cloner

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"path/filepath"
	"strconv"
	"strings"
	"os"
)

// Chapter
type Chapter struct {
	Title string
	Url   string
}

// Start clone all of book chapters from url to output file
func Start(url string, outFile string) {
	outFile, _ = filepath.Abs(outFile)
	fmt.Printf("Clone book from %s to %s? (yes/no): ", url, outFile)
	var choise string
	fmt.Scanf("%s", &choise)
	if choise == "yes" {
		chapters := getChapters(url)
		fmt.Printf("Total: %d chap(s) was found\n", len(chapters))
		saved := saveChapterContent(chapters, outFile)
		if saved {
			fmt.Printf("All chapter(s) was saved in %s\n", outFile)
		} else {
			fmt.Printf("Problem when saving chapter(s) contents! Aborted!\n")
		}

	} else {
		fmt.Printf("User aborted!\n")
	}
}

// getChapters get all URL of book chapters
// It returns URL as slice
func getChapters(url string) []Chapter {
	fmt.Printf("Getting chapter list ...\n")
	doc, err := goquery.NewDocument(url)
	var chaps []Chapter
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil
	} else {
		// Find the last page of chapters
		maxPage := 0;
		doc.Find(".paging ul > li > a").Each(func(i int, s *goquery.Selection) {
			page := s.Text()
			if pageNum, err := strconv.Atoi(page); err == nil {
				if pageNum > maxPage {
					maxPage = pageNum
				}
			}
		})

		// Get chapter list
		for i := 1; i <= maxPage; i = i + 1 {
			page, err := goquery.NewDocument(url + "?page=" + strconv.Itoa(i))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return nil
			} else {
				page.Find("#dschuong > div > .jblack").Each(func(i int, s *goquery.Selection) {
					link, _ := s.Attr("href")
					text := strings.TrimSpace(s.Text())
					chaps = append(chaps, Chapter{text, strings.TrimSpace(link)})
					// fmt.Printf("Found chap %s on %s\n", text, link)
				})
			}

		}

		return chaps
	}
}

func saveChapterContent(chaps []Chapter, outFile string) bool {
	chapCount := len(chaps)
	fmt.Printf("Get %d chapter%s contents ...", chapCount, func() string {
		if chapCount > 1 {
			return "s"
		}
		return ""
	})

	// Create file
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Can not create file: %s\n", outFile)
		return false
	}

	// Get contents then write to file
	for i, chap := range chaps {
		fmt.Printf("- Downloading (%d/%d): %s ...\n", i+1, chapCount, chap.Url)

		contents := getChapterContents(i+1, chap)
		file.WriteString(contents)
	}

	// Close file
	err = file.Close()
	if err != nil {
		fmt.Printf("Can not close file: %s\n", outFile)
		return false
	}

	return true
}

func getChapterContents(order int, chap Chapter) string {
	doc, err := goquery.NewDocument(chap.Url)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return ""
	} else {
		content, _ := doc.Find("#id_noidung_chuong").Html()
		return fmt.Sprintf("<center><h3 id=\"chap-%d\">%s</h3></center>\n", order, chap.Title) + content
	}
}