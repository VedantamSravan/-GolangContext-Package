## What is context API?

The context package in Go provides a mechanism to prevent the application from doing unnecessary work.
context is a standard package of Golang that makes it easy to pass request-scoped values, 
cancelation signals,and deadlines across API boundaries to all the goroutines involved in 
handling a request.

Package context defines the Context type, which carries deadlines,
cancellation signals, and other request-scoped values across API boundaries and between processes.

## Why Use Context?
- It simplifies the implementation for deadlines and cancellation across your processes or API.
- It prepares your code for scaling, for example, using Context will make your code clean 
 and easy to manipulate in the future by chaining all your process in a child parent relationship,
  you can tie/join any process together.
- Goroutine safe, i.e you can run the same context on different goroutines without leaks.



## Creating a Context
    The two ways to create a Context:
        - context.Background()
        - context.TODO()
        Both functions return a non-nil, empty context. The only time TODO is used instead of Background is when the implementation is unclear or the context is not yet known.

        context.TODO()function which is used to create an empty (or starting) context. We can use this when we don't know which context we want to use particularly. It will act like a placeholder.


 ## We have four options for making our context stop the program execution if a long period has passed:

    - context.Value
    - context.WithCancel
    - context.WithTimeout
    - context.WithDeadline

    1)context.value:
    One of the most common uses for contexts is to share data, or use request scoped values. 
    When you have multiple functions and you want to share data between them, 
    you can do so using contexts. The easiest way to do that is to use the function context.WithValue. 
    This function creates a new context based on a parent context and adds a value to a given key. 

Example:
    package main

    import (
        "context"
        "fmt"
    )
    
    func main() {
        ctx := context.Background()
        ctx = addValue(ctx)
        readValue(ctx)
    }
    
    func addValue(ctx context.Context) context.Context {
        return context.WithValue(ctx, "key", "test-value")
    }
    
    func readValue(ctx context.Context) {
        val := ctx.Value("key")
        fmt.Println(val)
    }
    
## 2)context.WithTimeout: Read file with Context Timeout

context.WithTimeout function to create a new context that is canceled when the specified timeout elapses. 
The function takes two arguments: an existing context and a duration for the timeout.

Example 1:

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
	/*context.WithTimeout to create a context that is canceled after 5 seconds */
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	select {
	/* reading file exceeds more than five then it prints  context error message 'context deadline  exceeded' */
	case <-time.After(time.Second * 4):
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
	scanner.Scan() 
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}


## Example2: Database Query Timeout using Context
    context package to set a deadline for the query execution.
    First, create a context with a timeout using the context.WithTimeout function. 
    Then, pass the context as the first argument to the query execution function 
    (such as db.QueryContext() or db.ExecContext()).

    package main
    import (
      "context"
      "database/sql"
      "fmt"
      "time"
    )
    
    func main() {
      /* Open a connection to the database */
      db, _ := sql.Open("driverName", "dataSourceName")
    
      /*
       Create a context with a timeout of 1 second
      If the query takes longer than 1 second to execute, it will be cancelled and an error will be returned.
      **/

      ctx, cancel := context.WithTimeout(context.Background(), time.Second)
      defer cancel()
    
      /* Execute the query with the context */
      rows, err := db.QueryContext(ctx, "SELECT * FROM table")
      if err != nil {
        fmt.Println(err)
      }
      defer rows.Close()
    
      /* Handle the query results */
      
    }


## Example4:Using Context for HTTP
    
  package main

  import (
      "context"
      "fmt"
      "io/ioutil"
      "net/http"
      "time"
  
      "golang.org/x/net/context/ctxhttp"
  )
  
  func main() {
      ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
      defer cancel()
    /*
     Request will be canceled if the Done channel is closed by the context, 
     but it will not close the connection. 
     It's the responsibility of the application to close the connection
    **/

    resp, err := ctxhttp.Get(ctx, nil, "https://example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
      defer resp.Body.Close()
  
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          fmt.Println("Error:", err)
          return
      }
  
      fmt.Println(string(body))
  }


##  context.WithDeadline
  
  context.WithDeadline function creates a new context with an associated deadline.
  The deadline is a specific point in time after which the context will be considered "dead" and any associated work will be cancelled.
  The function takes in two arguments: the existing context, and the deadline time. 
  It returns a new context that will be cancelled at the specified deadline.

  package main

  import (
      "context"
      "fmt"
      "time"
      "os"
      "bufio"
  )
  /* test.ini 
 {"NodeName":"ncvm1","NodeIP":"127.0.0.1","Upordown":"1","Port":"[0:1 Gbps  1:Down  ]","Status":"Running","Duration":"00:07:35:00","BeginTime":"2023-04-23 18:40:00","EndTime":"2023-04-24 02:15:00","License":"Evaluation","TimeZone":"UTC","PreCaptureFilter":"On","VirtualStorage":"1288.74GB","RealStorage":"455.00GB","Capturedrops":"0","BeginTimeSeconds":"1682275200","CaptureServerTime":"139728548030336","Throughput":"0.04","CompressionRatio":"2.83","ClusterCount":"0","tcppps":"7147","udppps":"318","otherpps":"0","totalpps":"7465","LogDataCompressionRatio":"1.00","PercentIOWait":"82.00","LoadAverage":"0.00 8.41 8.42"}
 **/

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


##  context.WithCancel:
  ==================
  context.WithCancel returns a new context and a cancel function.
  We defer the cancel function so that it is called when the main function exits. 
  it will check if the context has been done, if yes, the function will return.

## Why Do We Need Cancellation?
    - In short, we need cancellation to prevent our system from doing unnecessary work.
        Consider the common situation of an HTTP server making a call to a database, 
        and returning the queried data to the client:
    - But, what would happen if the client cancelled the request in the middle? 
        This could happen if, for example, the client closed their browser mid-request.
    - Without cancellation, the application server and database would continue to do their work, 
      even though the result of that work would be wasted:


Example:
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
	/* perform some operation and that causes error */
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