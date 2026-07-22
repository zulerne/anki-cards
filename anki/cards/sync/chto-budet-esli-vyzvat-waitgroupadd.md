---
id: sync_chto_budet_esli_vyzvat
deck: Go
tags:
  - sync
---

# Front

Что будет, если вызвать `WaitGroup.Add` после `Wait`?

# Back

Data race, непредсказуемое поведение. Всегда вызывай `Add` до запуска горутины.
