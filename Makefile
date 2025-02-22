.PHONY: dev swagger build test clean seed install local-ssl-proxy

dev:
	air & local-ssl-proxy --source 3000 --target 8080 --cert certs/localhost.pem --key certs/localhost-key.pem

swagger:
	swag init

build:
	go build -o bin/schichtplaner

test:
	go test -v ./...

clean:
	rm -rf bin/
	rm -rf tmp/
	rm -f *.db

seed:
	go run cmd/seed.go

install:
	go mod download
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

local-ssl-proxy:
	local-ssl-proxy --source 3000 --target 8080 --cert localhost.pem --key localhost-key.pem
