package main
import (
	"context"
	"fmt"
	"os"
	"bufio"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	go func() { // run the work in the background
		if err := readfile(ctx, "./test.ini"); err != nil {
			log.Println(err)
		}
	}()
	// perform some operation and that causes error
	time.Sleep(time.Millisecond * 10)
	if true { // err != nil
		cancel()
	}

}

func readfile(ctx context.Context,filename string) error{
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(filename)
	if err != nil {
		return  err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan() // this moves to the next token
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}