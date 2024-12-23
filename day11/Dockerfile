FROM golang
COPY . .
ARG IS_PRODUCTION=false
# if its production, add a compilation flag
RUN if [ "$IS_PRODUCTION" = "true" ]; then go build -o main main.go -ldflags "-s -w"; else go build -o main main.go; fi
CMD ["./main"]

