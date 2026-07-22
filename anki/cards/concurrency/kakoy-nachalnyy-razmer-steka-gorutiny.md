---
id: concurrency_kakoy_nachalnyy_razmer_steka
deck: Go
tags:
  - concurrency
---

# Front

Какой начальный размер стека горутины? Может ли он расти?

# Back

Начальный: 2KB. Растёт в 2x (contiguous stack — копируется в новое место). Максимум: 1GB (64-bit) или 250MB (32-bit).
