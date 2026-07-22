---
id: slices_kak_pravilno_skopirovat_slays
deck: Go
tags:
  - slices
---

# Front

Как правильно скопировать слайс, чтобы изменения не влияли на оригинал?

# Back

`dst := make([]int, len(src))` + `copy(dst, src)`.

Или в Go 1.21+: `dst := slices.Clone(src)`. Это shallow copy: элементы-указатели
не клонируются. `slices.Clone` также сохраняет nilness исходного слайса.
