package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ankiConnectURL = "http://localhost:8765"
	deckName       = "Go Interview"
	outputDir      = "anki/cards"
)

type ankiRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params,omitempty"`
}

type findNotesResponse struct {
	Result []int64 `json:"result"`
	Error  *string `json:"error"`
}

type noteInfo struct {
	NoteID int64             `json:"noteId"`
	Tags   []string          `json:"tags"`
	Fields map[string]field  `json:"fields"`
}

type field struct {
	Value string `json:"value"`
	Order int    `json:"order"`
}

type notesInfoResponse struct {
	Result []noteInfo `json:"result"`
	Error  *string    `json:"error"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	noteIDs, err := findNotes()
	if err != nil {
		return fmt.Errorf("find notes: %w", err)
	}
	fmt.Printf("found %d notes\n", len(noteIDs))

	notes, err := getNotesInfo(noteIDs)
	if err != nil {
		return fmt.Errorf("get notes info: %w", err)
	}

	tagGroups := groupByPrimaryTag(notes)
	written := 0

	for tag, group := range tagGroups {
		dir := filepath.Join(outputDir, tag)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", dir, err)
		}

		for i, note := range group {
			filename := generateFilename(note, i)
			path := filepath.Join(dir, filename+".md")

			content := renderCard(note, tag)
			if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
				return fmt.Errorf("write %s: %w", path, err)
			}
			written++
		}
	}

	fmt.Printf("written %d cards to %s\n", written, outputDir)
	return nil
}

func findNotes() ([]int64, error) {
	req := ankiRequest{
		Action:  "findNotes",
		Version: 6,
		Params:  map[string]string{"query": fmt.Sprintf("deck:%q", deckName)},
	}

	var resp findNotesResponse
	if err := ankiCall(req, &resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("anki: %s", *resp.Error)
	}
	return resp.Result, nil
}

func getNotesInfo(ids []int64) ([]noteInfo, error) {
	req := ankiRequest{
		Action:  "notesInfo",
		Version: 6,
		Params:  map[string]any{"notes": ids},
	}

	var resp notesInfoResponse
	if err := ankiCall(req, &resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("anki: %s", *resp.Error)
	}
	return resp.Result, nil
}

func ankiCall(req ankiRequest, result any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	resp, err := http.Post(ankiConnectURL, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("post: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func groupByPrimaryTag(notes []noteInfo) map[string][]noteInfo {
	groups := make(map[string][]noteInfo)
	skipTags := map[string]bool{"go": true, "interview": true, "batch1": true, "batch2": true, "batch3": true}

	for _, note := range notes {
		primary := "general"
		for _, tag := range note.Tags {
			if !skipTags[tag] {
				primary = tag
				break
			}
		}
		groups[primary] = append(groups[primary], note)
	}
	return groups
}

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9-]`)

func generateFilename(note noteInfo, index int) string {
	front := note.Fields["Front"].Value
	front = stripHTML(front)

	words := strings.Fields(front)
	if len(words) > 5 {
		words = words[:5]
	}

	slug := strings.ToLower(strings.Join(words, "-"))
	slug = transliterate(slug)
	slug = nonAlphanumeric.ReplaceAllString(slug, "")
	slug = strings.Trim(slug, "-")

	if slug == "" {
		slug = fmt.Sprintf("card-%d", index+1)
	}
	if len(slug) > 50 {
		slug = slug[:50]
	}
	return slug
}

func generateID(note noteInfo, tag string) string {
	front := note.Fields["Front"].Value
	front = stripHTML(front)

	words := strings.Fields(front)
	if len(words) > 4 {
		words = words[:4]
	}

	slug := strings.ToLower(strings.Join(words, "_"))
	slug = transliterate(slug)
	slug = regexp.MustCompile(`[^a-z0-9_]`).ReplaceAllString(slug, "")
	slug = strings.Trim(slug, "_")

	if slug == "" {
		slug = fmt.Sprintf("%d", note.NoteID)
	}

	return tag + "_" + slug
}

func renderCard(note noteInfo, primaryTag string) string {
	front := note.Fields["Front"].Value
	back := note.Fields["Back"].Value

	front = stripHTML(front)
	back = stripHTML(back)

	tags := filterTags(note.Tags)
	id := generateID(note, primaryTag)

	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "id: %s\n", id)
	b.WriteString("deck: Go\n")
	b.WriteString("tags:\n")
	for _, t := range tags {
		fmt.Fprintf(&b, "  - %s\n", t)
	}
	b.WriteString("---\n\n")
	b.WriteString("# Front\n\n")
	b.WriteString(front)
	b.WriteString("\n\n# Back\n\n")
	b.WriteString(back)
	b.WriteString("\n")

	return b.String()
}

func filterTags(tags []string) []string {
	skip := map[string]bool{"go": true, "interview": true, "batch1": true, "batch2": true, "batch3": true}
	var result []string
	for _, t := range tags {
		if !skip[t] {
			result = append(result, t)
		}
	}
	if len(result) == 0 {
		result = []string{"general"}
	}
	return result
}

func stripHTML(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	s = re.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	return strings.TrimSpace(s)
}

var translitMap = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "e",
	'ж': "zh", 'з': "z", 'и': "i", 'й': "y", 'к': "k", 'л': "l", 'м': "m",
	'н': "n", 'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u",
	'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh", 'щ': "shch",
	'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu", 'я': "ya",
}

func transliterate(s string) string {
	var b strings.Builder
	for _, r := range s {
		if mapped, ok := translitMap[r]; ok {
			b.WriteString(mapped)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
