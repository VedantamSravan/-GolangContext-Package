package main

import (
	"context"
	"fmt"
	"time"
	"os"
	"bufio"
)

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(time.Second * 6)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	select {
	//reading file exceeds more than five then it prints  context error message 'context deadline exceeded'
	case <-time.After(time.Second * 5):
		readfile("test.ini")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func readfile(filename string){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan() // this moves to the next token
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}