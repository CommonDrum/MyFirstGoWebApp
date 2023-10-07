FROM golang:alpine


WORKDIR app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o build/myapp

EXPOSE 8080
ENTRYPOINT [ "./build/myapp" ]