---
id: algorithms_nazovi_algoritmy_sortirovki_i
deck: Go
tags:
  - algorithms
---

# Front

Назови алгоритмы сортировки и их сложность.

# Back

Bubble Sort: O(n^2) — учебный.

Insertion Sort: O(n^2) — хорош для почти отсортированных.

Merge Sort: O(n log n) — стабильный, доп. память O(n).

Quick Sort: O(n log n) среднее, O(n^2) худшее — на практике самый быстрый.

Heap Sort: O(n log n) — in-place, нестабильный.

Go `sort.Sort` использует pattern-defeating quicksort (pdqsort).
