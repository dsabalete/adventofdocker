# Day 9 - Environment Variables & Configuration

Hey there! Welcome to Day 9 of Advent of Docker. Today we’re going to talk about something super important: how to configure your Docker containers using environment variables and build arguments.

## Why Environment Variables?

Environment variables are a great way to:

- Configure your application without changing code
- Handle different environments (development, staging, production)
- Manage sensitive information like API keys and passwords

Let’s modify our HTTP server from Day 5 to use environment variables:

```go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Get configuration from environment variables
    port := os.Getenv("PORT")
    secretMessage := os.Getenv("SECRET_MESSAGE")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!\nSecret message: %s", secretMessage)
    })

    fmt.Printf("Listening on port %s\n", port)
    http.ListenAndServe(":" + port, nil)
}

```

Here we simply take the PORT and SECRET_MESSAGE from the environment variables and use them in our application.

Our Dockerfile stays the same:

```Dockerfile
FROM golang
COPY . .
RUN go build -o main main.go
CMD ["./main"]

```

## Running with Environment Variables

There are two main ways to pass environment variables to a container:

PS: Dont forget to build the image before running it!

1. Using the -e flag:

```bash
docker run -e SECRET_MESSAGE="Hello, Docker!" -e PORT=8080 my-app

```

You might not want to leak your secrets to the command line, so you can also use a `.env` file:

2. Using the --env-file flag:

Create a file named .env:

```plaintext
SECRET_MESSAGE="Hello, Docker!"
PORT=8080

```

Run the container with the --env-file flag:

```bash
docker run --env-file .env my-app

```

Internally, Docker will insert the environment variables into the container at runtime, like normal process environment variables. During the build process, the environment variables are not available.

## Build Arguments vs Environment Variables

There’s another way to configure Docker: build arguments. The key difference is:

- Build args are only available during image build
- Environment variables are available at runtime

Build arguments are made to configure the build process. For example, you might want to compile your application slightly differently for production and development. Let’s see an example.

Modify the Dockerfile:

```Dockerfile
FROM golang
COPY . .
ARG IS_PRODUCTION=false
# if its production, add a compilation flag
RUN if [ "$IS_PRODUCTION" = "true" ]; then go build -o main main.go -ldflags "-s -w"; else go build -o main main.go; fi
CMD ["./main"]

```

In this Dockerfile we add

```Dockerfile
ARG IS_PRODUCTION=false

```

This defines that the build argument IS_PRODUCTION is set to false by default. Build the image with the build argument:

```bash
docker build -t my-app  --build-arg IS_PRODUCTION=true .

```

If we specify `IS_PRODUCTION=true`, the build will execute the if statement and add the -ldflags "-s -w" flag to the build command. If we don’t specify it, it will use the default value (false) and not add the flag.

Now is a great time to think about your usecases. What would you use build arguments for? What would you use environment variables for? Why did you choose one over the other? An example can only show you so much, try it out! If you’re not sure, feel free to message me :)

## Best Practices

As always, there are some best practices to follow:

1. Never hardcode secrets in your Dockerfile or source code
2. Don’t commit .env files to version control
   - Use `.env.example` to show what variables are needed
   - Add `.env` to `.gitignore`
3. Use build args for build-time configuration
4. Use env vars for runtime configuration
5. Provide defaults for non-sensitive configuration
6. Validate required variables on startup

Here’s a better version of our app following these practices:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}
func main() {
	// Optional env var with default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Required env var
	secretMessage := getRequiredEnv("SECRET_MESSAGE")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\nSecret message: %s", secretMessage)
	})
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

```

Here we changed the SECRET_MESSAGE to be required and added a default value for the PORT environment variable. That makes our code already much more robust and usable. Additionally, we should create a .env.example file to show what environment variables are needed:

```plaintext
SECRET_MESSAGE=
PORT=8080

```

Now anyone using your app knows what environment variables they need to set!
