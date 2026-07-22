---
id: stdlib_kak_rabotaet_encodingjson_chto
deck: Go
tags:
  - stdlib
---

# Front

Как работает `encoding/json`? Что такое struct tags?

# Back

`json.Marshal` сериализует структуру в JSON, `json.Unmarshal` — обратно.

Struct tags управляют именами и поведением:

- `json:"name,omitempty"` — имя поля в JSON + пропуск пустых значений.
- `json:"-"` — игнорировать поле.
