---
id: design-patterns_chto_takoe_dependency_injection
deck: Go
tags:
  - design-patterns
---

# Front

Что такое Dependency Injection в Go? Как реализуется без фреймворков?

# Back

Передача зависимостей через параметры конструктора (интерфейсы). Вместо создания зависимости внутри — принимаем извне:func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}В тестах передаём мок, в проде — реальную реализацию. Фреймворки (wire, fx) не обязательны.
