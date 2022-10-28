FROM golang:1.17-alpine
RUN mkdir api 
COPY . /api
WORKDIR /api
RUN go mod vendor
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 8080
