---
id: concurrency-patterns_obyasni_pattern_pipeline
deck: Go
tags:
  - concurrency-patterns
---

# Front

Объясни паттерн Pipeline.

# Back

Цепочка стадий обработки, каждая стадия — горутина с входным и выходным каналом. Данные текут последовательно: gen -> square -> print. Каждая стадия работает конкурентно. Удобно для потоковой обработки данных, ETL.
