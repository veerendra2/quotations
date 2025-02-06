package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"maps"
	"math/rand"
	"os"
	"slices"

	"github.com/charmbracelet/lipgloss"
)

//go:embed quotes/*.json
var quotesFiles embed.FS

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	NSFW      bool   `json:"nsfw"`
}

func main() {
	nsfw := flag.Bool("nsfw", false, "Enable NSFW quotes")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "A tool to display famous TV Shows and Movies quotes.")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Get random JSON file
	files, _ := quotesFiles.ReadDir("quotes")
	randFile := files[rand.Intn(len(files))]
	fileContent, _ := quotesFiles.ReadFile("quotes/" + randFile.Name())

	var selectedQuotes map[string][]Quote
	if err := json.Unmarshal(fileContent, &selectedQuotes); err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}

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
	fmt.Println(characterStyle.Render(fmt.Sprintf("─ %s (%s) 🎬", randomQuote.Character, randomTitle)))
}
