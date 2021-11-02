package functype

/*
同样输出Hello, World，只是给Greeting类型添加了say()方法。上面说了，函数类型是表示所有包含相同参数和返回类型的函数集合。我们在一开始先把func(name string) string这样的函数声明成Greeting类型，接着我们通过Greeting(english)将english函数转换成Greeting类型。通过这个转换以后，我们就可以借由变量g调用Greeting类型的say()方法。两段代码的差异就是go的类型系统添加方法和类C++语言添加类型方法的差异，具体讲解可以去查看《Go语言编程》第3章为类型添加方法这一节。

既然是函数集合，那么只有一个函数显然是不足以说明问题的。
*/
/*
import "fmt"

// Greeting function types
type Greeting func(name string) string

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func french(name string) string {
	return "Bonjour, " + name
}

func main() {
	g := Greeting(english)
	g.say("World")
	g = Greeting(french)
	g.say("World")
}
*/
