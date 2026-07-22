---
id: architecture_solid_chto_oznachaet_s
deck: Go
tags:
  - architecture
---

# Front

SOLID: что означает S — Single Responsibility Principle?

# Back

Класс/модуль должен иметь одну причину для изменения. Пример: отдельно UserRepository (хранение), отдельно UserService (бизнес-логика), отдельно UserHandler (HTTP). Каждый модуль меняется по одной причине.
