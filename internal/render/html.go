package render

import (
	"strings"
)

func MarkdownToHTML(md string) string {
	lines := strings.Split(md, "\n")
	var out strings.Builder
	var i int

	for i < len(lines) {
		line := lines[i]

		if strings.HasPrefix(line, "```") {
			i++
			lang, _ := strings.CutPrefix(line, "```")
			var code strings.Builder
			for i < len(lines) && !strings.HasPrefix(lines[i], "```") {
				if code.Len() > 0 {
					code.WriteByte('\n')
				}
				code.WriteString(escapeHTML(lines[i]))
				i++
			}
			if i < len(lines) {
				i++
			}
			if lang != "" {
				out.WriteString(`<pre><code class="language-`)
				out.WriteString(lang)
				out.WriteString(`">`)
			} else {
				out.WriteString("<pre><code>")
			}
			out.WriteString(code.String())
			out.WriteString("</code></pre>\n")
			continue
		}

		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			out.WriteString("<ul>\n")
			for i < len(lines) && (strings.HasPrefix(lines[i], "- ") || strings.HasPrefix(lines[i], "* ")) {
				item := strings.TrimPrefix(strings.TrimPrefix(lines[i], "- "), "* ")
				out.WriteString("<li>")
				out.WriteString(inlineFormat(item))
				out.WriteString("</li>\n")
				i++
			}
			out.WriteString("</ul>\n")
			continue
		}

		if matched, rest := cutOrderedPrefix(line); matched {
			out.WriteString("<ol>\n")
			out.WriteString("<li>")
			out.WriteString(inlineFormat(rest))
			out.WriteString("</li>\n")
			i++
			for i < len(lines) {
				if m, r := cutOrderedPrefix(lines[i]); m {
					out.WriteString("<li>")
					out.WriteString(inlineFormat(r))
					out.WriteString("</li>\n")
					i++
				} else {
					break
				}
			}
			out.WriteString("</ol>\n")
			continue
		}

		if strings.TrimSpace(line) == "" {
			out.WriteString("<br>\n")
			i++
			continue
		}

		out.WriteString("<p>")
		out.WriteString(inlineFormat(line))
		out.WriteString("</p>\n")
		i++
	}

	return strings.TrimSpace(out.String())
}

func inlineFormat(s string) string {
	s = escapeHTML(s)
	s = replaceInline(s, "**", "<b>", "</b>")
	s = replaceInline(s, "__", "<b>", "</b>")
	s = replaceInlineCode(s)
	return s
}

func replaceInline(s, delim, open, close string) string {
	var out strings.Builder
	for {
		start := strings.Index(s, delim)
		if start == -1 {
			out.WriteString(s)
			break
		}
		end := strings.Index(s[start+len(delim):], delim)
		if end == -1 {
			out.WriteString(s)
			break
		}
		end += start + len(delim)
		out.WriteString(s[:start])
		out.WriteString(open)
		out.WriteString(s[start+len(delim) : end])
		out.WriteString(close)
		s = s[end+len(delim):]
	}
	return out.String()
}

func replaceInlineCode(s string) string {
	var out strings.Builder
	for {
		start := strings.Index(s, "`")
		if start == -1 {
			out.WriteString(s)
			break
		}
		end := strings.Index(s[start+1:], "`")
		if end == -1 {
			out.WriteString(s)
			break
		}
		end += start + 1
		out.WriteString(s[:start])
		out.WriteString("<code>")
		out.WriteString(s[start+1 : end])
		out.WriteString("</code>")
		s = s[end+1:]
	}
	return out.String()
}

func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func cutOrderedPrefix(line string) (bool, string) {
	for i, c := range line {
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '.' && i > 0 && i < len(line)-1 && line[i+1] == ' ' {
			return true, line[i+2:]
		}
		break
	}
	return false, ""
}
