---
id: concurrency-patterns_obyasni_fanout__fanin
deck: Go
tags:
  - concurrency-patterns
---

# Front

Объясни Fan-Out / Fan-In. В чём разница?

# Back

Fan-Out: один канал читают несколько горутин — распределение работы (параллелизация).Fan-In: несколько каналов сливаются в один — объединение результатов. Часто используются вместе: Fan-Out для обработки, Fan-In для сбора.
