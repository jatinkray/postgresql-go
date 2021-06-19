############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS build
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /app
ADD . .
# Build the binary.
COPY ./static /app/build/static/
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/build/goapp


############################
# STEP 2 binary release
############################
FROM alpine AS release
COPY --from=build /app/build/ /app
CMD ["/app/goapp"]