
# LABEL maintainer="Maycon Moreira <mayconvm@gmail.com>"

# Build project
FROM golang:alpine as builder

RUN mkdir /pismo
ADD . /pismo/
WORKDIR /pismo

RUN apk add --no-cache git

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/kataras/tablewriter
RUN go get -u github.com/landoop/tableprinter

RUN CGO_ENABLED=0 go test

RUN go build -o mainPismo .

# Run project
FROM alpine
RUN adduser -S -D -H -h /app pismo
USER pismo
EXPOSE 8090

COPY --from=builder /pismo/mainPismo /app/
WORKDIR /app
CMD ["./mainPismo"]