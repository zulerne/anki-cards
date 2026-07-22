---
id: stdlib_chto_proizoydet_pri_marshalinge
deck: Go
tags:
  - stdlib
---

# Front

Что произойдёт при маршалинге структуры с неэкспортированным полем?

# Back

Неэкспортированные поля (с маленькой буквы) игнорируются encoding/json. Они не попадут в JSON при Marshal и не заполнятся при Unmarshal. Пакет json использует reflect, а reflect не может читать неэкспортированные поля.
