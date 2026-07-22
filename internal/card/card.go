package card

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Card struct {
	ID       string
	Deck     string
	Tags     []string
	Front    string
	Back     string
	FilePath string
}

func LoadAll(cardsDir string) ([]Card, error) {
	var cards []Card

	err := filepath.WalkDir(cardsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		c, err := ParseFile(path)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		cards = append(cards, c)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk cards: %w", err)
	}
	return cards, nil
}

func ParseFile(path string) (Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Card{}, fmt.Errorf("read file: %w", err)
	}
	return Parse(string(data), path)
}

func Parse(content, filePath string) (Card, error) {
	c := Card{FilePath: filePath}

	frontmatter, body, err := splitFrontmatter(content)
	if err != nil {
		return Card{}, err
	}

	c.ID, c.Deck, c.Tags, err = parseFrontmatter(frontmatter)
	if err != nil {
		return Card{}, fmt.Errorf("frontmatter: %w", err)
	}

	c.Front, c.Back, err = parseBody(body)
	if err != nil {
		return Card{}, fmt.Errorf("body: %w", err)
	}

	return c, nil
}

func splitFrontmatter(content string) (frontmatter, body string, err error) {
	const delimiter = "---"
	lines := strings.Split(content, "\n")

	if len(lines) == 0 || strings.TrimSpace(lines[0]) != delimiter {
		return "", "", errors.New("missing opening ---")
	}

	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == delimiter {
			end = i
			break
		}
	}
	if end == -1 {
		return "", "", errors.New("missing closing ---")
	}

	frontmatter = strings.Join(lines[1:end], "\n")
	body = strings.Join(lines[end+1:], "\n")
	return frontmatter, body, nil
}

func parseFrontmatter(fm string) (id, deck string, tags []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(fm))
	inTags := false

	for scanner.Scan() {
		line := scanner.Text()

		if inTags {
			trimmed := strings.TrimSpace(line)
			if tag, ok := strings.CutPrefix(trimmed, "- "); ok {
				tags = append(tags, strings.TrimSpace(tag))
				continue
			}
			inTags = false
		}

		key, value, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		switch key {
		case "id":
			id = value
		case "deck":
			deck = value
		case "tags":
			inTags = true
		}
	}

	if id == "" {
		return "", "", nil, errors.New("missing id")
	}
	if deck == "" {
		return "", "", nil, errors.New("missing deck")
	}
	return id, deck, tags, nil
}

func parseBody(body string) (front, back string, err error) {
	lines := strings.Split(body, "\n")
	frontHeader, backHeader := -1, -1
	inCodeBlock := false
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			continue
		}
		switch strings.TrimSpace(line) {
		case "# Front":
			frontHeader = i
		case "# Back":
			backHeader = i
		}
	}

	if frontHeader == -1 {
		return "", "", errors.New("missing # Front section")
	}
	if backHeader == -1 {
		return "", "", errors.New("missing # Back section")
	}
	if frontHeader > backHeader {
		return "", "", errors.New("# Front must come before # Back")
	}

	front = strings.TrimSpace(strings.Join(lines[frontHeader+1:backHeader], "\n"))
	back = strings.TrimSpace(strings.Join(lines[backHeader+1:], "\n"))

	return front, back, nil
}
