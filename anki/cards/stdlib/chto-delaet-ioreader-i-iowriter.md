---
id: stdlib_chto_delaet_ioreader_i
deck: Go
tags:
  - stdlib
---

# Front

Что делает io.Reader и io.Writer? Почему это главные интерфейсы в Go?

# Back

io.Reader: один метод Read(p []byte) (n int, err error). io.Writer: один метод Write(p []byte) (n int, err error). Абстрагируют источник/приёмник байтов (файл, сеть, буфер, сжатие). Минимальные интерфейсы — позволяют компоновать через io.Copy, bufio, gzip и т.д.
