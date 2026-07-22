---
id: channels_chto_proizoydet_pri_chtenii
deck: Go
tags:
  - channels
---

# Front

Что произойдёт при чтении из закрытого канала?

# Back

Вернёт оставшиеся значения из буфера, затем zero value + false (второе значение comma-ok).
