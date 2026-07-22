# Anki Cards System

Единственный источник истины для Go Anki-колоды. Все карточки хранятся в git как markdown-файлы.

## Структура

```
anki/
├── cards/          # карточки по темам
│   ├── runtime/
│   ├── concurrency/
│   ├── interfaces/
│   └── ...
├── templates/      # HTML/CSS шаблоны Anki
├── media/          # изображения для карточек
└── generated/      # результат сборки (в .gitignore)
```

## Формат карточки

Каждая карточка — отдельный `.md` файл с YAML frontmatter:

```markdown
---
id: unique_card_id
deck: Go
tags:
  - topic
  - subtopic
---

# Front

Вопрос карточки

# Back

Ответ карточки
```

### Правила

- `id` — уникальный идентификатор, snake_case
- `deck` — имя колоды (сейчас только `Go`)
- `tags` — список тегов для навигации в Anki
- `# Front` — вопрос (одна сторона карточки)
- `# Back` — ответ (обратная сторона)

## Как добавить карточку

1. Создать файл в подходящей директории:

```bash
touch anki/cards/runtime/new-topic.md
```

2. Заполнить по шаблону выше
3. Запустить проверку:

```bash
task anki:check
```

## Как изменить шаблон

Шаблоны Anki находятся в `anki/templates/`:

- `front.html` — лицевая сторона
- `back.html` — обратная сторона
- `styling.css` — стили карточек

После изменения шаблоны нужно вручную обновить в Anki Desktop (Manage Note Types → Cards).

## Как собрать колоду

```bash
task anki:build
```

Результат: `anki/generated/cards.tsv` — TSV-файл для импорта.

## Как импортировать в Anki

1. `task anki:build`
2. Anki Desktop → File → Import
3. Выбрать `anki/generated/cards.tsv`
4. Настройки импорта:
   - Type: Basic
   - Deck: Go
   - Field separator: Tab
   - Allow HTML in fields: Yes
   - Field 1 → id (используется для обновления)
   - Field 2 → Front
   - Field 3 → Back

## Проверки

`task anki:check` выполняет два уровня проверок:

### Формат (ошибки)

- Дубликаты id
- Дубликаты вопросов
- Отсутствует Front / Back
- Пустые поля
- Битые ссылки на изображения
- Неизвестный deck

### Знания (предупреждения)

- Ответ длиннее 250 слов
- Несколько вопросов на одной карточке
- Упоминание deprecated пакетов (io/ioutil и т.д.)
- Устаревшие концепции (cooperative scheduler, interface{})
