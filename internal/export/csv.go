package export

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/zulerne/anki-cards/internal/card"
)

func WriteCSV(w io.Writer, cards []card.Card) error {
	cw := csv.NewWriter(w)
	cw.Comma = '\t'

	if err := cw.Write([]string{"id", "front", "back", "deck", "tags"}); err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	for _, c := range cards {
		record := []string{
			c.ID,
			c.Front,
			c.Back,
			c.Deck,
			strings.Join(c.Tags, " "),
		}
		if err := cw.Write(record); err != nil {
			return fmt.Errorf("write record %s: %w", c.ID, err)
		}
	}

	cw.Flush()
	return cw.Error()
}
