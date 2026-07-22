---
id: channels_select_s_neskolkimi_gotovymi
deck: Go
tags:
  - channels
---

# Front

select с несколькими готовыми каналами — какой case выберется?

# Back

Один из готовых выбирается псевдослучайно. Это сделано для fairness — предотвращает starvation, когда один канал всегда доминирует.
