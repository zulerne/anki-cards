---
id: sync_dlya_chego_nuzhen_syncwaitgroup
deck: Go
tags:
  - sync
---

# Front

Для чего нужен sync.WaitGroup?

# Back

Ожидание завершения группы горутин. Add(n) увеличивает счётчик, Done() уменьшает (= Add(-1)), Wait() блокирует до счётчика 0.
