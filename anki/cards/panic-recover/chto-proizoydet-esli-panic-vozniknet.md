---
id: panic-recover_chto_proizoydet_esli_panic
deck: Go
tags:
  - panic-recover
---

# Front

Что произойдёт, если panic возникнет в горутине, а recover не установлен?

# Back

Упадёт вся программа. panic в горутине не перехватывается recover из другой горутины (в том числе main). Каждая горутина, которая может паниковать, должна иметь свой defer + recover.
