package card

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := `---
id: test_card
deck: Go
tags:
  - runtime
  - gc
---

# Front

What is GC?

# Back

Garbage collector.
`
	c, err := Parse(input, "test.md")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.ID != "test_card" {
		t.Errorf("id = %q, want %q", c.ID, "test_card")
	}
	if c.Deck != "Go" {
		t.Errorf("deck = %q, want %q", c.Deck, "Go")
	}
	if len(c.Tags) != 2 || c.Tags[0] != "runtime" || c.Tags[1] != "gc" {
		t.Errorf("tags = %v, want [runtime gc]", c.Tags)
	}
	if c.Front != "What is GC?" {
		t.Errorf("front = %q, want %q", c.Front, "What is GC?")
	}
	if c.Back != "Garbage collector." {
		t.Errorf("back = %q, want %q", c.Back, "Garbage collector.")
	}
}

func TestParse_MissingID(t *testing.T) {
	input := `---
deck: Go
tags:
  - test
---

# Front

Q?

# Back

A.
`
	_, err := Parse(input, "test.md")
	if err == nil {
		t.Fatal("expected error for missing id")
	}
}

func TestParse_MissingFront(t *testing.T) {
	input := `---
id: test
deck: Go
---

# Back

A.
`
	_, err := Parse(input, "test.md")
	if err == nil {
		t.Fatal("expected error for missing # Front")
	}
}
