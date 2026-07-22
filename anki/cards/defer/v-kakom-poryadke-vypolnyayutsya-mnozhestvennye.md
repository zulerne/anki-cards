---
id: defer_v_kakom_poryadke_vypolnyayutsya
deck: Go
tags:
  - defer
---

# Front

В каком порядке выполняются множественные `defer`?

# Back

LIFO (стек) — последний `defer` выполняется первым.
