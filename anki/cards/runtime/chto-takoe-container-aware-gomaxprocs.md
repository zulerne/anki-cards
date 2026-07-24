---
id: runtime_chto_takoe_container_aware_gomaxprocs
deck: Go
tags:
  - runtime
  - concurrency
---

# Front

Что изменилось в default `GOMAXPROCS` в Go 1.25+?

# Back

Если `GOMAXPROCS` явно не задан, Go 1.25+ учитывает доступный CPU limit контейнера и может обновлять значение при его изменении.

Явная переменная окружения или вызов `runtime.GOMAXPROCS` имеют приоритет. `GOMAXPROCS` определяет доступный runtime уровень параллелизма, но не задаёт общее число созданных OS threads.
