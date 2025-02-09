package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"maps"
	"math/rand"
	"slices"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/lipgloss"
)

//go:embed quotes/**/*.json
var quotesFiles embed.FS

var cli struct {
	NSFW          bool `help:"Enable NSFW quotes" short:"n"`
	Entertainment bool `help:"Display entertainment (movies and tv shows) quotes" short:"e"`
	Inspirational bool `help:"Display inspirational (famous figures) quotes" short:"i"`
}

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	NSFW      bool   `json:"nsfw"`
}

func main() {

	kongCli := kong.Parse(&cli)

	if cli.Entertainment && cli.Inspirational {
		kongCli.Fatalf("--entertainment (-e) and --inspirational (-i) are mutually exclusive. Please choose one.")
	}

	dirPath := "quotes/entertainment"
	if cli.Inspirational {
		dirPath = "quotes/inspirational"
	}

	// Get random JSON file
	files, _ := quotesFiles.ReadDir(dirPath)
	randFile := files[rand.Intn(len(files))]
	fileContent, _ := quotesFiles.ReadFile(dirPath + "/" + randFile.Name())

	var selectedQuotes map[string][]Quote
	if err := json.Unmarshal(fileContent, &selectedQuotes); err != nil {
		kongCli.Fatalf("error unmarshaling JSON: %v", err)
	}

	// Get random title
	titles := slices.Collect(maps.Keys(selectedQuotes))
	randomTitle := titles[rand.Intn(len(titles))]

	var filteredQuotes []Quote
	if cli.NSFW {
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
	if cli.Inspirational {
		fmt.Println(characterStyle.Render(fmt.Sprintf("─ %s", randomTitle)))
	} else {
		fmt.Println(characterStyle.Render(fmt.Sprintf("─ %s (%s) 🎬", randomQuote.Character, randomTitle)))
	}
}
