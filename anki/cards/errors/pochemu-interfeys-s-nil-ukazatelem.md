---
id: errors_pochemu_interfeys_s_nilukazatelem
deck: Go
tags:
  - errors
---

# Front

Почему интерфейс с nil-указателем != nil?var p *MyError = nilvar err error = p

# Back

Интерфейс == nil только когда ОБА поля (type и data) равны nil. Здесь type = *MyError (не nil), data = nil. Поэтому err != nil, хотя значение внутри — nil.
