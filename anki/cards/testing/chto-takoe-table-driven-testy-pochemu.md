---
id: testing_chto_takoe_tabledriven_testy
deck: Go
tags:
  - testing
---

# Front

Что такое table-driven тесты? Почему это идиоматичный подход в Go?

# Back

Массив тестовых случаев (слайс структур), итерация по ним через for _, tc := range tests + t.Run(tc.name, ...).Идиоматично: легко добавлять случаи без дублирования кода, каждый подтест именован (видно в выводе), можно запускать один: go test -run TestFoo/case_name.
