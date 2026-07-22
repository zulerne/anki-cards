---
id: stdlib_chto_delaet_io_copy
deck: Go
tags:
  - io
  - streams
---

# Front

Для чего нужны `io.Copy` и `io.CopyBuffer`?

# Back

Они переносят байты из `io.Reader` в `io.Writer` до конца потока и возвращают
число скопированных байт и ошибку.

`io.Copy` подходит для обычной потоковой передачи без загрузки всего файла в
память. `io.CopyBuffer` позволяет передать переиспользуемый буфер.
