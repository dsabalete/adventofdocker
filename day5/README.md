# Day 5 - Interacting with Containers

Hey there! Congrats to making it to day 5. Today we will continue with our Go application from Day 4 and make it a bit more interesting. We will start by running a simple HTTP server and then see how we can interact with it!

As a quick reminder, yesterday we created a Dockerfile that looked like this:

```Dockerfile
FROM golang

COPY . .

RUN go build -o main main.go

CMD ["./main"]

```

And our main.go file looked like this:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

If you didn’t create it yesterday, please create a new directory and add the files to it!

Printing hello world is of course a bit boring, so let’s make our application a bit more interesting by adding a simple HTTP server.

To do this we will modify our main.go file to look like this:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!") // this will be printed as a response when you request /
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil) // this will start the HTTP server (blocking operation) and listen on port 8080
}

```

Here we basically just create one route (/) that will print “Hello World!” when someone visits the root path. The HTTP server will be listening on port 8080. If you feel adventurous, you can try to add some more routes, change the port or the message.

You can still use the same Dockerfile from yesterday and build the image just like we did yesterday.

```bash
$ docker build -t hello-world-go .

```

And then run it just like we did yesterday.

```bash
$ docker run hello-world-go

```
