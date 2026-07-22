---
id: stdlib_chem_jsonmarshal_otlichaetsya_ot
deck: Go
tags:
  - stdlib
---

# Front

Чем `json.Marshal` отличается от `json.NewEncoder`? Когда что использовать?

# Back

`json.Marshal` — возвращает `[]byte`, весь JSON в памяти. Для небольших данных, тестов, сохранения в переменную.

`json.NewEncoder(w)` — пишет сразу в `io.Writer` (HTTP response, файл). Для потоковой записи, экономит аллокацию промежуточного буфера.
