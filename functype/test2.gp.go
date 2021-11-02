package functype

/*
import "fmt"
//say()函数要求传入一个Greeting类型，因为english函数的参数和返回值跟Greeting一样，参考接口的概念这里可以做类型转换。我们换个方式来实现上面的功能:

// Greeting function types
type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	g := Greeting(english)
	g.say("World")
}
*/
