package main
import(
	"fmt"
	"strings"
	"bufio"
	"os"
)
func CountLastWordLen(str string)int{
	if len(str) == 0 || len(str)>=5000{
		return 0
	}
	arr := strings.Split(str," ")
	return len(arr[len(arr)-1])
}

func main(){
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.Trim(a,"\n")
	lastWordLen := CountLastWordLen(a)
	fmt.Println(lastWordLen)
}


