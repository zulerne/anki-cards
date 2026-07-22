package lint

import (
	"fmt"
	"strings"

	"github.com/zulerne/anki-cards/internal/card"
)

var deprecatedPackages = []string{
	"io/ioutil",
	"golang.org/x/net/context",
}

var outdatedConcepts = []struct {
	pattern string
	note    string
}{
	{"cooperative scheduler", "Go scheduler is preemptive since 1.14"},
	{"GOMAXPROCS(1)", "default is NumCPU since Go 1.5"},
	{"interface{}", "use 'any' since Go 1.18"},
}

const maxAnswerWords = 250

func CheckKnowledge(cards []card.Card) []Issue {
	var issues []Issue

	for _, c := range cards {
		issues = append(issues, checkAnswerLength(c)...)
		issues = append(issues, checkMultipleQuestions(c)...)
		issues = append(issues, checkDeprecated(c)...)
		issues = append(issues, checkOutdated(c)...)
	}
	return issues
}

func checkAnswerLength(c card.Card) []Issue {
	words := len(strings.Fields(c.Back))
	if words > maxAnswerWords {
		return []Issue{{
			Severity: Warning,
			FilePath: c.FilePath,
			Message:  fmt.Sprintf("answer is %d words (max recommended: %d)", words, maxAnswerWords),
		}}
	}
	return nil
}

func checkMultipleQuestions(c card.Card) []Issue {
	count := strings.Count(c.Front, "?")
	if count > 1 {
		return []Issue{{
			Severity: Warning,
			FilePath: c.FilePath,
			Message:  fmt.Sprintf("front contains %d questions — consider splitting", count),
		}}
	}
	return nil
}

func checkDeprecated(c card.Card) []Issue {
	var issues []Issue
	combined := c.Front + "\n" + c.Back

	for _, pkg := range deprecatedPackages {
		if strings.Contains(combined, pkg) {
			issues = append(issues, Issue{
				Severity: Warning,
				FilePath: c.FilePath,
				Message:  fmt.Sprintf("mentions deprecated package %s", pkg),
			})
		}
	}
	return issues
}

func checkOutdated(c card.Card) []Issue {
	var issues []Issue
	combined := c.Front + "\n" + c.Back

	for _, concept := range outdatedConcepts {
		if strings.Contains(combined, concept.pattern) {
			issues = append(issues, Issue{
				Severity: Warning,
				FilePath: c.FilePath,
				Message:  fmt.Sprintf("mentions %q — %s", concept.pattern, concept.note),
			})
		}
	}
	return issues
}
