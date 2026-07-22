package lint

import (
	"testing"

	"github.com/zulerne/anki-cards/internal/card"
)

func TestCheckDuplicateIDs(t *testing.T) {
	cards := []card.Card{
		{ID: "a", Deck: "Go", Front: "Q1?", Back: "A1", FilePath: "a.md"},
		{ID: "a", Deck: "Go", Front: "Q2?", Back: "A2", FilePath: "b.md"},
	}
	issues := Check(cards, Config{AllowedDecks: []string{"Go"}})

	found := false
	for _, issue := range issues {
		if issue.Severity == Error && issue.FilePath == "b.md" {
			found = true
		}
	}
	if !found {
		t.Error("expected duplicate id error for b.md")
	}
}

func TestCheckUnknownDeck(t *testing.T) {
	cards := []card.Card{
		{ID: "a", Deck: "Python", Front: "Q?", Back: "A", FilePath: "a.md"},
	}
	issues := Check(cards, Config{AllowedDecks: []string{"Go"}})

	found := false
	for _, issue := range issues {
		if issue.Severity == Error && issue.FilePath == "a.md" {
			found = true
		}
	}
	if !found {
		t.Error("expected unknown deck error")
	}
}

func TestCheckEmptyFields(t *testing.T) {
	cards := []card.Card{
		{ID: "a", Deck: "Go", Front: "", Back: "A", FilePath: "a.md"},
	}
	issues := Check(cards, Config{AllowedDecks: []string{"Go"}})

	found := false
	for _, issue := range issues {
		if issue.Severity == Error && issue.FilePath == "a.md" {
			found = true
		}
	}
	if !found {
		t.Error("expected empty front error")
	}
}
