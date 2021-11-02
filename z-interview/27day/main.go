package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
	//根据 iota 的用法推断出 South 的值是 2；另外，如果类型定义了 String() 方法，当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印。

}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func test() {
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)
}
