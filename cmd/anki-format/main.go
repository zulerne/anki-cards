package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const cardsDir = "anki/cards"

var qualifiedIdent = regexp.MustCompile(`([a-z]\w*\.[A-Z]\w*(?:\.[A-Z]\w*)?)`)
var funcCall = regexp.MustCompile(`([a-z]\w*\.[A-Z]\w*\(\))`)
var goKeywords = regexp.MustCompile(`\b(ctx\.Done\(\)|ctx\.Err\(\)|context\.Background\(\)|context\.TODO\(\)|context\.WithCancel|context\.WithTimeout|context\.WithValue|context\.Context|sync\.Mutex|sync\.RWMutex|sync\.WaitGroup|sync\.Once|sync\.Map|sync\.Pool|errors\.Is|errors\.As|errors\.New|fmt\.Errorf|fmt\.Println|fmt\.Sprintf|fmt\.Stringer|io\.Reader|io\.Writer|io\.Closer|http\.Handler|http\.HandlerFunc|http\.Server|time\.Duration|time\.Ticker|time\.Timer)\b`)
var internalTypes = regexp.MustCompile(`\b(eface|iface|itab|_type|mcache|mspan|mheap|sudog|hchan|waitq)\b`)
var envVars = regexp.MustCompile(`\b(GOMAXPROCS|GOGC|GOMEMLIMIT|GOROOT|GOPATH)\b`)
var builtinFuncs = regexp.MustCompile(`\b(append|copy|delete|len|cap|close|panic|recover|make|new)\(`)
var percentVerbs = regexp.MustCompile(`(%[wvsdTqxp#])`)
var ptrType = regexp.MustCompile(`(\*[A-Z]\w+)\b`)
var interfaceEmpty = regexp.MustCompile(`\binterface\{\}`)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var files []string
	err := filepath.WalkDir(cardsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	modified := 0
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		original := string(data)
		result := processCard(original)
		if result != original {
			if err := os.WriteFile(path, []byte(result), 0o644); err != nil {
				return err
			}
			modified++
		}
	}
	fmt.Printf("modified %d/%d cards\n", modified, len(files))
	return nil
}

func processCard(content string) string {
	lines := strings.Split(content, "\n")
	inFrontmatter := false
	inCodeBlock := false
	var result []string

	for i, line := range lines {
		if i == 0 && strings.TrimSpace(line) == "---" {
			inFrontmatter = true
			result = append(result, line)
			continue
		}
		if inFrontmatter {
			if strings.TrimSpace(line) == "---" {
				inFrontmatter = false
			}
			result = append(result, line)
			continue
		}
		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			result = append(result, line)
			continue
		}
		if inCodeBlock || strings.HasPrefix(line, "# ") {
			result = append(result, line)
			continue
		}
		result = append(result, formatLine(line))
	}
	return strings.Join(result, "\n")
}

func formatLine(line string) string {
	line = applyRegex(line, goKeywords)
	line = applyRegex(line, qualifiedIdent)
	line = applyRegex(line, funcCall)
	line = applyRegex(line, internalTypes)
	line = applyRegex(line, envVars)
	line = applyRegex(line, percentVerbs)
	line = applyRegex(line, ptrType)
	line = applyRegex(line, interfaceEmpty)
	line = applyBuiltins(line)
	return line
}

func applyRegex(line string, re *regexp.Regexp) string {
	return re.ReplaceAllStringFunc(line, func(match string) string {
		return wrapSafe(line, match)
	})
}

func applyBuiltins(line string) string {
	return builtinFuncs.ReplaceAllStringFunc(line, func(match string) string {
		name := match[:len(match)-1]
		wrapped := wrapSafe(line, name)
		return wrapped + "("
	})
}

func wrapSafe(line, match string) string {
	idx := strings.Index(line, match)
	if idx == -1 {
		return match
	}
	if idx > 0 && line[idx-1] == '`' {
		return match
	}
	end := idx + len(match)
	if end < len(line) && line[end] == '`' {
		return match
	}
	backticksBeforeCount := strings.Count(line[:idx], "`")
	if backticksBeforeCount%2 == 1 {
		return match
	}
	return "`" + match + "`"
}
