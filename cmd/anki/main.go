package main

import (
	"fmt"
	"os"

	"github.com/zulerne/anki-cards/internal/card"
	"github.com/zulerne/anki-cards/internal/export"
	"github.com/zulerne/anki-cards/internal/lint"
)

const (
	cardsDir     = "anki/cards"
	mediaDir     = "anki/media"
	templateDir  = "anki/templates"
	generatedDir = "anki/generated"
	apkgFile     = "anki/generated/go-deck.apkg"
)

var allowedDecks = []string{"Go"}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cards, err := card.LoadAll(cardsDir)
	if err != nil {
		return fmt.Errorf("load cards: %w", err)
	}

	if len(cards) == 0 {
		fmt.Println("no cards found")
		return nil
	}

	fmt.Printf("loaded %d cards\n", len(cards))

	cfg := lint.Config{
		AllowedDecks: allowedDecks,
		MediaDir:     mediaDir,
	}

	issues := lint.Check(cards, cfg)
	knowledgeIssues := lint.CheckKnowledge(cards)
	issues = append(issues, knowledgeIssues...)

	hasErrors := false
	for _, issue := range issues {
		fmt.Println(issue)
		if issue.Severity == lint.Error {
			hasErrors = true
		}
	}

	if hasErrors {
		return fmt.Errorf("lint failed with errors")
	}

	if err := os.MkdirAll(generatedDir, 0o755); err != nil {
		return fmt.Errorf("create generated dir: %w", err)
	}

	if err := export.WriteAPKG(apkgFile, cards, mediaDir, templateDir); err != nil {
		return fmt.Errorf("export apkg: %w", err)
	}

	fmt.Printf("exported %d cards to %s\n", len(cards), apkgFile)
	return nil
}
