package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zulerne/anki-cards/internal/card"
)

type Severity int

const (
	Error Severity = iota
	Warning
)

type Issue struct {
	Severity Severity
	FilePath string
	Message  string
}

func (i Issue) String() string {
	prefix := "⚠"
	if i.Severity == Error {
		prefix = "✗"
	}
	return fmt.Sprintf("%s %s: %s", prefix, i.FilePath, i.Message)
}

type Config struct {
	AllowedDecks []string
	MediaDir     string
}

func Check(cards []card.Card, cfg Config) []Issue {
	var issues []Issue

	issues = append(issues, checkDuplicateIDs(cards)...)
	issues = append(issues, checkDuplicateQuestions(cards)...)
	issues = append(issues, checkEmptyFields(cards)...)
	issues = append(issues, checkDecks(cards, cfg.AllowedDecks)...)
	issues = append(issues, checkMedia(cards, cfg.MediaDir)...)

	return issues
}

func checkDuplicateIDs(cards []card.Card) []Issue {
	var issues []Issue
	seen := make(map[string]string)

	for _, c := range cards {
		if prev, exists := seen[c.ID]; exists {
			issues = append(issues, Issue{
				Severity: Error,
				FilePath: c.FilePath,
				Message:  fmt.Sprintf("duplicate id %q (also in %s)", c.ID, prev),
			})
		}
		seen[c.ID] = c.FilePath
	}
	return issues
}

func checkDuplicateQuestions(cards []card.Card) []Issue {
	var issues []Issue
	seen := make(map[string]string)

	for _, c := range cards {
		normalized := strings.ToLower(strings.TrimSpace(c.Front))
		if prev, exists := seen[normalized]; exists {
			issues = append(issues, Issue{
				Severity: Error,
				FilePath: c.FilePath,
				Message:  fmt.Sprintf("duplicate question (also in %s)", prev),
			})
		}
		seen[normalized] = c.FilePath
	}
	return issues
}

func checkEmptyFields(cards []card.Card) []Issue {
	var issues []Issue

	for _, c := range cards {
		if c.Front == "" {
			issues = append(issues, Issue{
				Severity: Error,
				FilePath: c.FilePath,
				Message:  "empty Front section",
			})
		}
		if c.Back == "" {
			issues = append(issues, Issue{
				Severity: Error,
				FilePath: c.FilePath,
				Message:  "empty Back section",
			})
		}
	}
	return issues
}

func checkDecks(cards []card.Card, allowed []string) []Issue {
	if len(allowed) == 0 {
		return nil
	}

	var issues []Issue
	set := make(map[string]struct{}, len(allowed))
	for _, d := range allowed {
		set[d] = struct{}{}
	}

	for _, c := range cards {
		if _, ok := set[c.Deck]; !ok {
			issues = append(issues, Issue{
				Severity: Error,
				FilePath: c.FilePath,
				Message:  fmt.Sprintf("unknown deck %q (allowed: %s)", c.Deck, strings.Join(allowed, ", ")),
			})
		}
	}
	return issues
}

func checkMedia(cards []card.Card, mediaDir string) []Issue {
	var issues []Issue

	for _, c := range cards {
		refs := extractMediaRefs(c.Front + "\n" + c.Back)
		for _, ref := range refs {
			path := filepath.Join(mediaDir, ref)
			if _, err := os.Stat(path); err != nil {
				issues = append(issues, Issue{
					Severity: Error,
					FilePath: c.FilePath,
					Message:  fmt.Sprintf("media not found: %s", ref),
				})
			}
		}
	}
	return issues
}

func extractMediaRefs(text string) []string {
	var refs []string
	for {
		idx := strings.Index(text, "![")
		if idx == -1 {
			break
		}
		text = text[idx:]
		start := strings.Index(text, "(")
		end := strings.Index(text, ")")
		if start == -1 || end == -1 || end < start {
			break
		}
		ref := text[start+1 : end]
		if !strings.HasPrefix(ref, "http") {
			refs = append(refs, ref)
		}
		text = text[end+1:]
	}
	return refs
}
