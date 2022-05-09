package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	resultChan := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go t(ctx, resultChan, i)
	}

	tim := time.NewTimer(time.Second * 3)
	defer tim.Stop()
	select {
	case <-tim.C:
		fmt.Println("all goroutine timeout, start cancel...")
		cancel()
	case <-resultChan:
		fmt.Println("allready find the result, start cancel...")
		cancel()
	}

	fmt.Println("main end...")
}

func t(ctx context.Context, resultChan chan struct{}, index int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel()....")
			return
		default:
			fmt.Printf("index %d is looking for...\n", index)
			if index == 2 {
				time.Sleep(time.Second * 4)
				fmt.Println("index 2 find the result, done")
				resultChan <- struct{}{}
				return
			} else {
				time.Sleep(time.Second * 6)
			}

			//if FindTarget(){
			//	resultChan <- struct{}{}
			//}
			//return
		}
	}
}

func FindTarget() bool {

	return true
}

func TestOther() {
	text := "abcd1234浮生无事"
	textLen := len(text)
	fmt.Printf("len:" + strconv.FormatInt(int64(textLen), 10))
	for i := 0; i <= textLen-1; i++ {
		fmt.Printf(fmt.Sprintf("word:%s", text[i:i+1]))
	}
}

func TestOther2() {
	text := "abcd1234浮生无事"
	textRune := []rune(text)
	textLen := len(textRune)
	fmt.Printf("len:" + strconv.FormatInt(int64(textLen), 10))
	for i := 0; i <= textLen-1; i++ {
		fmt.Printf(fmt.Sprintf("word:%s", string(textRune[i:i+1])))
	}
}
