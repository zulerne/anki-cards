---
id: concurrency_chem_gorutina_otlichaetsya_ot
deck: Go
tags:
  - concurrency
---

# Front

Чем горутина отличается от потока ОС?

# Back

Горутина: начальный стек ~2KB (растёт до 1GB), переключение в userspace (Go runtime). Поток ОС: стек ~8MB (фиксированный), переключение через ядро.
