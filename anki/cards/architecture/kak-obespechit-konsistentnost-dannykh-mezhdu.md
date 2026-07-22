---
id: architecture_kak_obespechit_konsistentnost_dannykh
deck: Go
tags:
  - architecture
---

# Front

Как обеспечить консистентность данных между микросервисами? (Saga pattern)

# Back

Saga — последовательность локальных транзакций с компенсирующими действиями при откате. Два типа:

Choreography — сервисы слушают события друг друга.

Orchestration — центральный координатор управляет шагами.

Если шаг 3 упал — выполняются компенсации шагов 2 и 1.
