package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var dogcounter uint64
	var fishcounter uint64
	var catcounter uint64

	wg.Add(3)
	dogch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)
	catch := make(chan struct{}, 1)

	go dog(&wg, dogcounter, dogch, fishch)
	go fish(&wg, fishcounter, fishch, catch)
	go cat(&wg, catcounter, catch, dogch)

	dogch <- struct{}{}
	wg.Wait()
}

func dog(wg *sync.WaitGroup, counter uint64, dogch, fishch chan struct{}) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-dogch
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishch <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, counter uint64, fishch, catch chan struct{}) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-fishch
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catch <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter uint64, catch, dogch chan struct{}) {
	for {
		if counter >= uint64(10) {
			wg.Done()
			return
		}
		<-catch
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogch <- struct{}{}
	}
}
