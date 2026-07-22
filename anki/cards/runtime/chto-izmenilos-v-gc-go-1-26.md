---
id: runtime_chto_izmenilos_v_gc_go_1_26
deck: Go
tags:
  - runtime
  - gc
---

# Front

Что изменилось в GC в Go 1.26?

# Back

В стандартном runtime Go 1.26 Green Tea GC включён по умолчанию. Он меняет
организацию mark-and-sweep для лучшей locality и масштабирования сканирования
небольших объектов.

Это деталь реализации, а не гарантия спецификации. В Go 1.26 его можно временно
отключить при сборке через `GOEXPERIMENT=nogreenteagc`; ожидается, что opt-out
будет удалён в Go 1.27.
