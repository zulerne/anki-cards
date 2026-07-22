---
id: strings_chem_otlichaetsya_iteratsiya_for
deck: Go
tags:
  - strings
---

# Front

Чем отличается итерация for i, r := range str от for i := 0; i < len(str); i++?

# Back

range итерирует по рунам (Unicode code points): i — байтовый offset, r — руна (rune).Индексная итерация по str[i] — по отдельным байтам (byte).Для ASCII разницы нет, для мультибайтовых символов range корректно декодирует UTF-8.
