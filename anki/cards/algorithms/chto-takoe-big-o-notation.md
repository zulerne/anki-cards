---
id: algorithms_chto_takoe_big_o
deck: Go
tags:
  - algorithms
---

# Front

Что такое Big O notation? В чём разница между O(n) и O(n log n)?

# Back

Big O — верхняя граница роста времени/памяти при увеличении входных данных.

O(n) — линейный рост (один проход). O(n log n) — типичная сложность эффективных сортировок (merge sort, heap sort). При n=1M: O(n)=1M операций, O(n log n) ~20M операций. Разница растёт с размером данных.
