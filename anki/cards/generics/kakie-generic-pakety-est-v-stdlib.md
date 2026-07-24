---
id: generics_kakie_generic_pakety_est_v_stdlib
deck: Go
tags:
  - generics
  - stdlib
---

# Front

Какие generic-пакеты стандартной библиотеки Go полезны для коллекций?

# Back

С Go 1.21 доступны generic-пакеты `slices`, `maps` и `cmp`.

Примеры: `slices.Clone`, `slices.Sort`, `maps.Clone`, `maps.Equal` и `cmp.Compare`. Они работают с пользовательскими типами, совместимыми с нужными constraints.
