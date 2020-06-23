package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	score    int
	question string
	t        time.Duration
)

// 問題を出す
func q() {
	rand.Seed(time.Now().UnixNano())
	words := [...]string{"table", "chair", "pen", "water"}
	question = words[rand.Intn(len(words))]
	fmt.Println("\ntype this: ", question)
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
	flag.DurationVar(&t, "time", 30*time.Second, "answer time")
	flag.Parse()
	// コンテキストによるタイムアウト
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	q()

	ch := input(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finish!!!")
			fmt.Println("Your score is", score)
			return
		case answer := <-ch:
			if answer == question {
				score++
				fmt.Println("◎")
			} else {
				fmt.Println("×")
			}
			q()
		}
	}

}
