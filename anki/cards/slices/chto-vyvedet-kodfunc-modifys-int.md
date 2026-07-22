---
id: slices_chto_vyvedet_kodfunc_modifys
deck: Go
tags:
  - slices
---

# Front

Что выведет код?

```go
func modify(s []int) { s = append(s, 4) }
s := []int{1, 2, 3}
modify(s)
fmt.Println(s)
```

# Back

`[1 2 3]`. Слайс передаётся по значению (копия дескриптора). `append` внутри функции может аллоцировать новый массив или изменить `len` только у локальной копии. Оригинальный дескриптор не изменился.
