---
id: concurrency_chto_vyvedet_kodfor_i
deck: Go
tags:
  - concurrency
---

# Front

Что выведет код?for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()
}
time.Sleep(time.Second)

# Back

До Go 1.22: 3 3 3 — все горутины замкнуты на одну переменную i, к моменту выполнения цикл завершился.С Go 1.22+: 0 1 2 (в произвольном порядке) — каждая итерация создаёт свою копию.Fix для старых версий: go func(i int) { fmt.Println(i) }(i).
