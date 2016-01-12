package thichdoctruyen

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Chapter
type Chapter struct {
	Title   string
	Url     string
	Content string
}

// Start clone all of book chapters from url to output file
func Start(url string, outFile string) {
	outFile, _ = filepath.Abs(outFile)
	fmt.Printf("Clone book from %s to %s? (yes/no): ", url, outFile)
	var choice string
	fmt.Scanf("%s", &choice)
	if choice == "yes" {
		chapters := getChapters(url)
		fmt.Printf("Total: %d chap(s) was found\n", len(chapters))
		saved := saveChapters(chapters, outFile)
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
func getChapters(Url string) []Chapter {
	fmt.Printf("Getting chapter list ...\n")

	// Find id
	regex := regexp.MustCompile(`(.*)-([\d]+)$`)
	matches := regex.FindStringSubmatch(Url)

	res, err := http.PostForm("http://thichdoctruyen.com/actions/ajaxTruyen/ajaxLoadChap.php", url.Values{"id_story": {matches[2]}, "name_story": {"Anything"}})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	var chaps []Chapter
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil
	} else {
		// Find the last page of chapters
		doc.Find("option").Each(func(i int, s *goquery.Selection) {
			chapUrl := matches[1] + "/" + s.AttrOr("value", "")
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				chaps = append(chaps, Chapter{"", chapUrl, ""})
			}
		})

		return chaps
	}
}

// saveChapters save all chapters content to output file
// It returns true if save success
func saveChapters(chaps []Chapter, outFile string) bool {
	chapCount := len(chaps)
	fmt.Printf("Get %d chapter(s) contents ...", chapCount)

	// Create file
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Can not create file: %s\n", outFile)
		return false
	}

	// Get contents then write to file
	for i, chap := range chaps {
		fmt.Printf("- Downloading (%d/%d): %s ...\n", i+1, chapCount, chap.Url)

		title, contents := getChapterContents(i+1, chap)
		chaps[i].Title = title
		chaps[i].Content = contents
	}

	// build bookmarks
	bookmarks := getTheBookmark(chaps)

	// Write HTML header

	file.WriteString("<html>\n")
	file.WriteString("<head>\n")
	file.WriteString("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />\n")
	file.WriteString("</head>\n")
	file.WriteString("<body>\n")
	file.WriteString(bookmarks)

	// Get contents then write to file
	for _, chap := range chaps {
		file.WriteString(chap.Content)
	}

	file.WriteString("</body>\n")
	file.WriteString("</html>\n")

	// Close file
	err = file.Close()
	if err != nil {
		fmt.Printf("Can not close file: %s\n", outFile)
		return false
	}

	return true
}

// getChapterContents download content of a chapter
// It returns content of the chapter
func getChapterContents(order int, chap Chapter) (string, string) {
	doc, err := goquery.NewDocument(chap.Url)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", ""
	} else {
		chapTitle := strings.Trim(doc.Find("p.tenchuong").Text(), " \n\r\t")
		content, _ := doc.Find(".boxview").Html()
		return chapTitle, fmt.Sprintf("<center><h3 id=\"chap-%d\">%s</h3></center>\n", order, chapTitle) + content
	}
}

func getTheBookmark(chaps []Chapter) string {
	html := "<h2>BOOKMARKS:</h2>\n"
	html += "<ul>\n"
	for i, chap := range chaps {
		html += fmt.Sprintf("<li><a href=\"#chap-%d\">%s</a></li>", i, chap.Title)
	}
	html += "</ul>\n"
	return html
}
