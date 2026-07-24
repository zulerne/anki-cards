---
id: slices_kakoy_strategiey_rosta_polzuetsya
deck: Go
tags:
  - slices
---

# Front

Какой стратегией роста пользуется слайс при `append`?

# Back

Точная стратегия роста — деталь реализации runtime и не гарантируется языком. Не полагайся на конкретные коэффициенты или пороги; если нужна ёмкость, заранее используй `make([]T, 0, capacity)`.
