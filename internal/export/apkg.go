package export

import (
	"archive/zip"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	"github.com/zulerne/anki-cards/internal/card"
)

const (
	modelID = 1700000000000
	deckID  = 1700000000001
)

func WriteAPKG(outputPath string, cards []card.Card, mediaDir string) error {
	tmpDB, err := os.CreateTemp("", "anki-*.db")
	if err != nil {
		return fmt.Errorf("create temp db: %w", err)
	}
	tmpDBPath := tmpDB.Name()
	_ = tmpDB.Close()
	defer func() { _ = os.Remove(tmpDBPath) }()

	if err := buildDB(tmpDBPath, cards); err != nil {
		return fmt.Errorf("build db: %w", err)
	}

	mediaFiles := collectMedia(cards, mediaDir)

	if err := packZip(outputPath, tmpDBPath, mediaFiles); err != nil {
		return fmt.Errorf("pack zip: %w", err)
	}

	return nil
}

func buildDB(dbPath string, cards []card.Card) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer func() { _ = db.Close() }()

	if err := createSchema(db); err != nil {
		return fmt.Errorf("create schema: %w", err)
	}

	if err := insertMetadata(db); err != nil {
		return fmt.Errorf("insert metadata: %w", err)
	}

	for _, c := range cards {
		if err := insertCard(db, c); err != nil {
			return fmt.Errorf("insert card %s: %w", c.ID, err)
		}
	}

	return nil
}

func createSchema(db *sql.DB) error {
	schema := `
		CREATE TABLE col (
			id INTEGER PRIMARY KEY,
			crt INTEGER NOT NULL,
			mod INTEGER NOT NULL,
			scm INTEGER NOT NULL,
			ver INTEGER NOT NULL,
			dty INTEGER NOT NULL,
			usn INTEGER NOT NULL,
			ls INTEGER NOT NULL,
			conf TEXT NOT NULL,
			models TEXT NOT NULL,
			decks TEXT NOT NULL,
			dconf TEXT NOT NULL,
			tags TEXT NOT NULL
		);
		CREATE TABLE notes (
			id INTEGER PRIMARY KEY,
			guid TEXT NOT NULL,
			mid INTEGER NOT NULL,
			mod INTEGER NOT NULL,
			usn INTEGER NOT NULL,
			tags TEXT NOT NULL,
			flds TEXT NOT NULL,
			sfld TEXT NOT NULL,
			csum INTEGER NOT NULL,
			flags INTEGER NOT NULL,
			data TEXT NOT NULL
		);
		CREATE TABLE cards (
			id INTEGER PRIMARY KEY,
			nid INTEGER NOT NULL,
			did INTEGER NOT NULL,
			ord INTEGER NOT NULL,
			mod INTEGER NOT NULL,
			usn INTEGER NOT NULL,
			type INTEGER NOT NULL,
			queue INTEGER NOT NULL,
			due INTEGER NOT NULL,
			ivl INTEGER NOT NULL,
			factor INTEGER NOT NULL,
			reps INTEGER NOT NULL,
			lapses INTEGER NOT NULL,
			left INTEGER NOT NULL,
			odue INTEGER NOT NULL,
			odid INTEGER NOT NULL,
			flags INTEGER NOT NULL,
			data TEXT NOT NULL
		);
		CREATE TABLE revlog (
			id INTEGER PRIMARY KEY,
			cid INTEGER NOT NULL,
			usn INTEGER NOT NULL,
			ease INTEGER NOT NULL,
			ivl INTEGER NOT NULL,
			lastIvl INTEGER NOT NULL,
			factor INTEGER NOT NULL,
			time INTEGER NOT NULL,
			type INTEGER NOT NULL
		);
		CREATE TABLE graves (
			usn INTEGER NOT NULL,
			oid INTEGER NOT NULL,
			type INTEGER NOT NULL
		);
	`
	_, err := db.Exec(schema)
	return err
}

func insertMetadata(db *sql.DB) error {
	now := time.Now().Unix()
	nowMs := time.Now().UnixMilli()

	models := map[string]any{
		fmt.Sprintf("%d", modelID): map[string]any{
			"id":    modelID,
			"name":  "Go Card",
			"type":  0,
			"mod":   now,
			"usn":   -1,
			"sortf": 0,
			"did":   deckID,
			"tmpls": []map[string]any{
				{
					"name":  "Card 1",
					"ord":   0,
					"qfmt":  "{{Front}}",
					"afmt":  "{{FrontSide}}<hr id=answer>{{Back}}",
					"bqfmt": "",
					"bafmt": "",
					"did":   nil,
					"bfont": "",
					"bsize": 0,
				},
			},
			"flds": []map[string]any{
				{"name": "Front", "ord": 0, "sticky": false, "rtl": false, "font": "Arial", "size": 20, "media": []any{}},
				{"name": "Back", "ord": 1, "sticky": false, "rtl": false, "font": "Arial", "size": 20, "media": []any{}},
			},
			"css":  ".card { font-family: -apple-system, sans-serif; font-size: 16px; line-height: 1.6; max-width: 640px; margin: 0 auto; padding: 20px; }",
			"latexPre": "",
			"latexPost": "",
			"latexsvg":  false,
			"req":       [][]any{{0, "all", []int{0}}},
			"tags":      []any{},
			"vers":      []any{},
		},
	}

	decks := map[string]any{
		"1": map[string]any{
			"id":             1,
			"mod":            now,
			"name":           "Default",
			"usn":            0,
			"lrnToday":       []int{0, 0},
			"revToday":       []int{0, 0},
			"newToday":       []int{0, 0},
			"timeToday":      []int{0, 0},
			"collapsed":      false,
			"browserCollapsed": false,
			"desc":           "",
			"dyn":            0,
			"conf":           1,
			"extendNew":      0,
			"extendRev":      0,
		},
		fmt.Sprintf("%d", deckID): map[string]any{
			"id":             deckID,
			"mod":            now,
			"name":           "Go",
			"usn":            -1,
			"lrnToday":       []int{0, 0},
			"revToday":       []int{0, 0},
			"newToday":       []int{0, 0},
			"timeToday":      []int{0, 0},
			"collapsed":      false,
			"browserCollapsed": false,
			"desc":           "",
			"dyn":            0,
			"conf":           1,
			"extendNew":      0,
			"extendRev":      0,
		},
	}

	dconf := map[string]any{
		"1": map[string]any{
			"id":       1,
			"mod":      0,
			"name":     "Default",
			"usn":      0,
			"maxTaken": 60,
			"autoplay": true,
			"timer":    0,
			"replayq":  true,
			"new": map[string]any{
				"bury":       false,
				"delays":     []float64{1, 10},
				"initialFactor": 2500,
				"ints":       []int{1, 4, 0},
				"order":      1,
				"perDay":     20,
			},
			"rev": map[string]any{
				"bury":     false,
				"ease4":    1.3,
				"ivlFct":   1,
				"maxIvl":   36500,
				"perDay":   200,
				"hardFactor": 1.2,
			},
			"lapse": map[string]any{
				"delays":   []float64{10},
				"leechAction": 1,
				"leechFails":  8,
				"minInt":   1,
				"mult":     0,
			},
		},
	}

	conf := map[string]any{
		"activeDecks":   []int{1},
		"curDeck":       1,
		"newSpread":     0,
		"collapseTime":  1200,
		"timeLim":       0,
		"estTimes":      true,
		"dueCounts":     true,
		"curModel":      modelID,
		"nextPos":       1,
		"sortType":      "noteFld",
		"sortBackwards": false,
		"addToCur":      true,
	}

	modelsJSON, _ := json.Marshal(models)
	decksJSON, _ := json.Marshal(decks)
	dconfJSON, _ := json.Marshal(dconf)
	confJSON, _ := json.Marshal(conf)

	_, err := db.Exec(`INSERT INTO col VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		1, now, nowMs, nowMs, 11, 0, -1, 0,
		string(confJSON), string(modelsJSON), string(decksJSON), string(dconfJSON), "{}")
	return err
}

func insertCard(db *sql.DB, c card.Card) error {
	noteID := guidToID(c.ID)
	cardID := noteID + 1
	now := time.Now().Unix()
	nowMs := time.Now().UnixMilli()

	guid := c.ID
	tags := " " + strings.Join(c.Tags, " ") + " "
	flds := c.Front + "\x1f" + c.Back
	sfld := c.Front
	csum := fieldChecksum(c.Front)

	_, err := db.Exec(`INSERT INTO notes VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		noteID, guid, modelID, now, -1, tags, flds, sfld, csum, 0, "")
	if err != nil {
		return fmt.Errorf("insert note: %w", err)
	}

	_, err = db.Exec(`INSERT INTO cards VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		cardID, noteID, deckID, 0, nowMs, -1, 0, 0, noteID, 0, 0, 0, 0, 0, 0, 0, 0, "")
	if err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}

func guidToID(id string) int64 {
	h := sha256.Sum256([]byte(id))
	v := int64(binary.BigEndian.Uint64(h[:8]))
	if v < 0 {
		v = -v
	}
	if v < 1_000_000_000_000 {
		v += 1_000_000_000_000
	}
	return v
}

func fieldChecksum(field string) int64 {
	h := sha256.Sum256([]byte(field))
	return int64(binary.BigEndian.Uint32(h[:4]))
}

func collectMedia(cards []card.Card, mediaDir string) map[string]string {
	media := make(map[string]string)
	for _, c := range cards {
		refs := extractMediaRefsFromCard(c)
		for _, ref := range refs {
			path := filepath.Join(mediaDir, ref)
			if _, err := os.Stat(path); err == nil {
				media[ref] = path
			}
		}
	}
	return media
}

func extractMediaRefsFromCard(c card.Card) []string {
	combined := c.Front + "\n" + c.Back
	var refs []string
	for {
		idx := strings.Index(combined, "![")
		if idx == -1 {
			break
		}
		combined = combined[idx:]
		start := strings.Index(combined, "(")
		end := strings.Index(combined, ")")
		if start == -1 || end == -1 || end < start {
			break
		}
		ref := combined[start+1 : end]
		if !strings.HasPrefix(ref, "http") {
			refs = append(refs, ref)
		}
		combined = combined[end+1:]
	}
	return refs
}

func packZip(outputPath, dbPath string, mediaFiles map[string]string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}

	zw := zip.NewWriter(f)

	dbFile, err := os.Open(dbPath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	w, err := zw.Create("collection.anki2")
	if err != nil {
		return fmt.Errorf("create zip entry: %w", err)
	}
	if _, err := io.Copy(w, dbFile); err != nil {
		return fmt.Errorf("copy db: %w", err)
	}
	_ = dbFile.Close()

	mediaMap := make(map[string]string)
	i := 0
	for name, path := range mediaFiles {
		key := fmt.Sprintf("%d", i)
		mediaMap[key] = name

		mf, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open media %s: %w", name, err)
		}
		mw, err := zw.Create(key)
		if err != nil {
			_ = mf.Close()
			return fmt.Errorf("create media entry: %w", err)
		}
		if _, err := io.Copy(mw, mf); err != nil {
			_ = mf.Close()
			return fmt.Errorf("copy media: %w", err)
		}
		_ = mf.Close()
		i++
	}

	mediaJSON, _ := json.Marshal(mediaMap)
	mw, err := zw.Create("media")
	if err != nil {
		return fmt.Errorf("create media json: %w", err)
	}
	if _, err := mw.Write(mediaJSON); err != nil {
		return fmt.Errorf("write media json: %w", err)
	}

	if err := zw.Close(); err != nil {
		return fmt.Errorf("close zip: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	return nil
}
