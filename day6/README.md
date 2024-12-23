# Day 6 - Docker Image Layers

Yesterday we played around with our Go HTTP server and learned how to interact with containers. Today, let’s understand how Docker images are actually built and why containers are ephemeral.

Remember our Dockerfile from the last two days?

```Dockerfile
FROM golang
COPY . .
RUN go build -o main main.go
CMD ["./main"]

```

Each instruction in this Dockerfile creates a new layer (or multiple). A layer is basically a set of file changes. Let’s break down what happens when Docker builds this image:

1. FROM golang - Pulls all base layers with the Go runtime
2. COPY . . - Adds a new layer containing our source files
3. RUN go build - Creates a layer with our compiled binary
4. CMD - Doesn’t create a layer (it’s just metadata)

You can actually see these layers with:

```bash
$ docker history hello-world-go
IMAGE CREATED CREATED BY SIZE COMMENT
a3c62afdbe9d 12 seconds ago CMD ["./main"] 0B buildkit.dockerfile.v0
<missing> 12 seconds ago EXPOSE map[8080/tcp:{}] 0B buildkit.dockerfile.v0
<missing> 12 seconds ago RUN /bin/sh -c go build -o main main.go # bu… 81MB buildkit.dockerfile.v0
<missing> 16 seconds ago COPY . . # buildkit 322B buildkit.dockerfile.v0
<missing> 3 days ago WORKDIR /go 0B buildkit.dockerfile.v0
<missing> 3 days ago RUN /bin/sh -c mkdir -p "$GOPATH/src" "$GOPA… 0B buildkit.dockerfile.v0
<missing> 3 days ago COPY /target/ / # buildkit 235MB buildkit.dockerfile.v0
<missing> 3 days ago ENV PATH=/go/bin:/usr/local/go/bin:/usr/loca… 0B buildkit.dockerfile.v0
<missing> 3 days ago ENV GOPATH=/go 0B buildkit.dockerfile.v0
<missing> 3 days ago ENV GOTOOLCHAIN=local 0B buildkit.dockerfile.v0
<missing> 3 days ago ENV GOLANG_VERSION=1.23.4 0B buildkit.dockerfile.v0
<missing> 3 days ago RUN /bin/sh -c set -eux; apt-get update; a… 239MB buildkit.dockerfile.v0
<missing> 11 months ago RUN /bin/sh -c set -eux; apt-get update; a… 183MB buildkit.dockerfile.v0
<missing> 19 months ago RUN /bin/sh -c set -eux; apt-get update; a… 47.8MB buildkit.dockerfile.v0
<missing> 19 months ago # debian.sh --arch 'arm64' out/ 'bookworm' '… 139MB debuerreotype 0.15

```

## Why Layers Matter

Layers are what make Docker so efficient:

1. Caching - If nothing changes in a layer, Docker reuses it from cache
2. Storage - Layers are shared between images that use the same base
3. Transfer - When pushing/pulling images, only changed layers are transferred

## Container Layer

When you run a container, Docker adds one more special writable layer on top of your image layers. This is where all changes made during container runtime are stored.

Let’s see this in action. Start our container and exec into it:

```bash
bash
$ docker run -d -p 8080:8080 hello-world-go
$ docker exec -it <container_id> bash
root@container:/# touch newfile.txt
root@container:/# exit

```

Now stop the container and start a new one:

```bash
bash
$ docker stop <container_id>
$ docker run -d -p 8080:8080 hello-world-go
$ docker exec -it <container_id> bash
root@container:/# ls newfile.txt
ls: cannot access 'newfile.txt': No such file or directory

```
