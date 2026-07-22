---
id: architecture_solid_chto_oznachaet_o
deck: Go
tags:
  - architecture
---

# Front

SOLID: что означает O — Open/Closed Principle?

# Back

Открыт для расширения, закрыт для модификации. Новое поведение добавляется через новые реализации интерфейса, а не изменением существующего кода.

Пример: интерфейс `Notifier` с реализациями `Email`, `SMS`, `Push` — добавляем `Telegram` без правки существующего кода.
