package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokvic", court)

	court <- -1 // 往通道传递数据

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court // 往通道获取数据
		if !ok {
			fmt.Println("Player %s Won\n", name)
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			close(court)

			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball // 往通道传递数据
	}
}
