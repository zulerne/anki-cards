---
id: channels_kogda_ispolzovat_mutex_a
deck: Go
tags:
  - channels
---

# Front

Когда использовать Mutex, а когда канал?

# Back

`Mutex` — для защиты разделяемого состояния (shared state). Канал — для передачи владения данными между горутинами (communication).

Правило Go: «Share memory by communicating, don't communicate by sharing memory.»
