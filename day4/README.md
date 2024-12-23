# Day 4 - Building your first container

It is finally time to build your first container! Today we will go through the process of dockerizing a simple application.

We will use a simple Go application that prints “Hello, World!” to the console. I purposefully chose Go because it is a very simple language and you probably don’t have it installed yet. This will show you one of the benefits of dockerizing your applications, you can use any language you want without actually needing to install it on your machine!

To get started create a new directory for your project and navigate into it. Then create a new file called main.go and add the following code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

As you can probably see, this is only going to print “Hello, World!” to the console and the exit. Baby steps! Now if you have go installed you can just execute the code with:

```bash
$ go run main.go
Hello, World!

```

I don’t expect you to actually run this on your machine, we will just use docker to build and run it!

Now that we have our application, we can start to dockerize it. To do this we will create a file called Dockerfile that will define how to build and run our application.

```dockerfile
FROM golang

COPY . .

RUN go build -o main main.go

CMD ["./main"]

```

A dockerfile is a simple text file that contains a set of instructions for docker on how to build and run your application. It is executed in order from top to bottom. Each instruction generally adds a new layer to the image, or in other words, it adds an additional set of files to your final image.

The first instruction FROM golang tells docker to use the official golang image as a base image. This image already contains the golang compiler and other dependencies we need to build our application. This is basically the same as if you would manually install golang.

```dockerfile
FROM golang

```

The next instruction COPY . . copies all files from the current directory into the image.

```dockerfile
COPY . .

```

The next instruction RUN go build -o main main.go builds our application and names the output file main.

```dockerfile
RUN go build -o main main.go

```

The last instruction CMD ["./main"] tells docker to run the main executable when the container starts.

```dockerfile
CMD ["./main"]

```

The Dockerfile itself does not do anything. To create an image from it, we will need to run the docker build command.

```bash
$ docker build -t hello-world-go .

```

This will create a new image with the name hello-world-go. The dot at the end tells Docker which directory to use as the root for the build.

The first time you run this command, Docker will start to download the golang base image from Docker Hub. This might take a while if you don’t have it yet! Subsequent runs will be faster because Docker is quite good at caching steps that do not change.

If the build was successful, you should be able to see it when you list all images with docker images.

```bash
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED
hello-world-go      latest              <image-id>          <timestamp>

```

If you don’t see the image, please let me know and I will help you out :)

Great job! You just built your first docker image! 🎉

Now that we have our image, we can run it with the docker run command.

```bash
$ docker run hello-world-go
Hello, World!

```

This will start a new container from our image and execute the main executable, because that is what we defined with the CMD instruction in our Dockerfile.

As you can see, the container exits immediately because that is what our application does.

If you run docker ps --all, you should see that the container exited immediately.

```bash
$ docker ps -a
CONTAINER ID     IMAGE            COMMAND     CREATED          STATUS
<container-id>   hello-world-go   "/main"     <timestamp>      Exited (0)

```

Make sure to run docker ps --all and not docker ps because docker ps only shows running containers by default. The --all flag is used to show all containers, including the ones that have already exited.

Now its your turn to play around with this small setup! Change some lines in the Dockerfile and see what happens. Here are some ideas:

- Change the CMD instruction to RUN go run main.go and see what happens. Does it still work? Why?
- Change the RUN instruction to RUN go build -o main main.go, why does it fail?
- Add a new COPY instruction to copy some other files into the image, does it work?
- Checkout the WORKDIR instruction and try to use it in your Dockerfile.
