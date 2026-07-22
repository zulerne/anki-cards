---
id: interfaces_gde_sleduet_obyavlyat_interfeys
deck: Go
tags:
  - interfaces
---

# Front

Где следует объявлять интерфейс — в пакете с реализацией или с использованием?

# Back

В пакете-потребителе (consumer), не производителе.Идиоматика Go: «Accept interfaces, return structs.»Это снижает связность — реализация не знает о интерфейсе, потребитель определяет минимальный контракт.Исключения: стандартные интерфейсы (io.Reader, error, fmt.Stringer).
