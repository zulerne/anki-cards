---
id: slices_kakoy_strategiey_rosta_polzuetsya
deck: Go
tags:
  - slices
---

# Front

Какой стратегией роста пользуется слайс при append?

# Back

До 256 элементов: рост ~2x. После 256 элементов: рост ~1.25x.
