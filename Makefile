#@echo "Hello From Make"

run:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./out/golang_orders_rest .
