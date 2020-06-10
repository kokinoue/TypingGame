package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	score int
	text  string
)

// 問題を出す
func question() {
	// text := "table"
	rand.Seed(time.Now().UnixNano())
	words := [...]string{"table", "chair", "pen", "water"}
	//FIXME: questionが次のinputとして見られてしまう
	fmt.Println("\ntype this: ", words[rand.Intn(len(words))])
	fmt.Print("> ")
}

// 答えをチャネルに持たせる
func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func main() {
	// コンテキストによるタイムアウト
	bc := context.Background()
	t := 30 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	question()

	ch := input(os.Stdin)

	// for-select
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finish!!!")
			fmt.Println("Your score is", score)
			return
		case v := <-ch:
			if v == text {
				score++
				fmt.Println("◎")
			} else {
				fmt.Println("×")
			}
			question()
		}
	}

}
