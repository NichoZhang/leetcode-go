##

## 目的

同步leetcode上的答题

* Array and String (5 Chapters, 29 Items, 9 Done)
* Queue and Stack (5 Chapters, 32 Items, 0 Done)
* Trie (4 Chapters, 14 Items, 0 Done)
* Binary Search (8 Chapters, 30 Items, 0 Done)
* Binary Search Tree (4 Chapters, 21 Items, 0 Done)


## `go test` 使用方式

本项目`go test`使用并非规范使用，仅仅是测试是否满足条件

``` golang
$ cd $GOPATH/src/leetcode
$ cd arrayStruct
$ go test -v PascalsTriangle*
output
=== RUN   TestPascalsTriangle1
--- PASS: TestPascalsTriangle1 (0.00s)
	Array_PascalsTriangle_test.go:6: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
=== RUN   TestPascalsTriangle2
--- PASS: TestPascalsTriangle2 (0.00s)
	PascalsTriangle_test.go:10: []
PASS
ok  	command-line-arguments	(cached)
```
