# anki-cards

Practice repo for rewriting Go code from Obsidian/Anki cards by hand.

## Workflow

```bash
task day                  # create today's directory (practice/mar22, etc.)
task new -- topic-name    # create a new topic with main.go template
task run -- topic-name    # run a topic
task build                # build all packages
task test                 # run all tests with race detector
```

Дневниковые решения находятся в `practice/daily/<date>/<topic>` и не хранятся
в Git. Для периодической очистки используй `task daily:clean`.

Эталонные runnable-примеры по темам находятся в
[`practice/reference`](practice/reference/README.md). Дневниковые решения
остаются в `practice/daily`, а reference-примеры поддерживаются как
идиоматичные варианты, связанные с карточками.
