# Day 13 - Multistage builds

Welcome to Day 13 of our Docker adventure! Today, we’ll explore multistage builds - a powerful feature that helps us create smaller, more secure Docker images. We’ll transform our previous Go application into a lean, production-ready container using multistage builds.

## What are Multistage Builds?

Multistage builds allow us to use multiple FROM statements in our Dockerfile. Each FROM instruction starts a new build stage, and we can selectively copy artifacts from one stage to another. This means we can include build tools and dependencies in earlier stages without bloating our final image.

## Before: Single Stage Build

Let’s look at our previous Go application Dockerfile:

```Dockerfile
FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
```

This Dockerfile works, but it includes the entire Go toolchain and build dependencies in our final image. Let’s check its size:

```bash
$ docker build -t hello-world-go .
$ docker images hello-world-go
REPOSITORY TAG IMAGE ID CREATED SIZE
hello-world-go latest 6209fff2e822 2 days ago 928MB

```

Nearly 1GB for a simple Go application! 😱

## After: Multistage Build#

Here’s how we can improve it using multistage builds:

```Dockerfile
# Build stage

FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN go build -o main .

# Final stage

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]

```

Let’s break down what’s happening:

1. The first stage (named builder) uses the full Go image to compile our application
2. The second stage starts fresh from a minimal Alpine Linux image
3. We copy only the compiled binary from the builder stage
4. Everything else from the build stage is discarded

Now let’s check the size:

```bash
$ docker build -t hello-world-go-multistage .
$ docker images hello-world-go-multistage
REPOSITORY TAG IMAGE ID CREATED SIZE
hello-world-go-multistage latest dc5a7b64247a 6 seconds ago 14.9MB

```

That’s a 98% reduction in size! 🎉

## More Examples

Multistage builds are not only useful for Go applications. They can be applied to any language or framework. Here are some more examples:

### Node.js Application

Here’s how a multistage build looks for a Node.js static website (like adventofdocker.com):

```Dockerfile
# Build stage - compiles the application

FROM node:lts AS base
WORKDIR /app
COPY package\*.json ./
RUN npm install
COPY . .
RUN npm run build

# Runtime stage - serves the static files

FROM nginx:mainline-alpine-slim AS runtime
COPY --from=base ./app/dist /usr/share/nginx/html
EXPOSE 80

```

This example shows how to build and serve a static website using two stages:

1. The build stage uses Node.js to install dependencies and build the application
2. The runtime stage uses a lightweight nginx image to serve only the built static files

This approach significantly reduces the final image size by excluding Node.js, npm, and all development dependencies from the runtime image.

## Benefits of Multistage Builds

1. Smaller Image Size: Final images contain only what’s necessary to run the application
2. Better Security: Fewer components mean a smaller attack surface
3. Faster Deployments: Smaller images are faster to push and pull
4. Clean Separation: Build-time dependencies are completely separated from runtime

## Best Practices

- Name your build stages for better readability
- Use specific base image versions instead of latest (e.g. FROM node:20 instead of FROM node:latest)
- Keep frequently changing files (like source code) in later COPY commands
- Consider using distroless or minimal base images for the final stage
