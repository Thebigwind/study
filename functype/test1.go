package functype

//A function type denotes the set of all functions with the same parameter and result types.
import "fmt"

// Greeting function types
type Greeting func(name string) string

func say(g Greeting, n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	say(english, "World")
}
