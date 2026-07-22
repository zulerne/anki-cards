---
id: algorithms_chto_takoe_kheshtablitsa_kak
deck: Go
tags:
  - algorithms
---

# Front

Что такое хеш-таблица? Как разрешаются коллизии?

# Back

Структура: массив + хеш-функция (ключ -> индекс). Среднее время операций O(1).Коллизия — два ключа попадают в один индекс. Методы разрешения:Chaining — список в каждой ячейке (Go map использует buckets).Open addressing — ищем следующую свободную ячейку (linear probing, quadratic probing).
