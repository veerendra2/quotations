package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"maps"
	"math/rand"
	"slices"

	"github.com/charmbracelet/lipgloss"
)

//go:embed quotes/**/*.json
var quotesFiles embed.FS

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	NSFW      bool   `json:"nsfw"`
}

func main() {
	nsfw := flag.Bool("n", false, "Enable NSFW quotes")
	entertainment := flag.Bool("e", false, "Display entertainment quotes (default)")
	inspirational := flag.Bool("i", false, "Display inspirational quotes")
	flag.Parse()

	if *entertainment && *inspirational {
		fmt.Println("Please choose either -e or -i, not both.")
		return
	}

	dirPath := "quotes/entertainment"
	if *inspirational {
		dirPath = "quotes/inspirational"
	}

	// Get random JSON file
	files, _ := quotesFiles.ReadDir(dirPath)
	randFile := files[rand.Intn(len(files))]
	fileContent, _ := quotesFiles.ReadFile(dirPath + "/" + randFile.Name())

	var selectedQuotes map[string][]Quote
	_ = json.Unmarshal(fileContent, &selectedQuotes)

	// Get random title
	titles := slices.Collect(maps.Keys(selectedQuotes))
	randomTitle := titles[rand.Intn(len(titles))]

	var filteredQuotes []Quote
	if *nsfw {
		filteredQuotes = selectedQuotes[randomTitle]
	} else {
		filteredQuotes = slices.DeleteFunc(selectedQuotes[randomTitle], func(q Quote) bool {
			return q.NSFW
		})
	}

	// Get random quote from selected title
	randomQuote := filteredQuotes[rand.Intn(len(filteredQuotes))]

	var quoteStyle = lipgloss.NewStyle().Italic(true).Bold(true).Width(70).Foreground(lipgloss.Color("229"))
	var characterStyle = lipgloss.NewStyle().Bold(true).PaddingLeft(30).Foreground(lipgloss.Color("6"))

	fmt.Println(quoteStyle.Render(fmt.Sprintf("❝%s❞", randomQuote.Quote)))
	if *inspirational {
		fmt.Println(characterStyle.Render(fmt.Sprintf("─ %s", randomTitle)))
	} else {
		fmt.Println(characterStyle.Render(fmt.Sprintf("─ %s (%s)", randomQuote.Character, randomTitle)))
	}
}
