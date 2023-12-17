build:
	@go build -o ./bin/serveme ./cmd/server/main.go 
	@go build -o ./bin/sendme ./cmd/client/main.go 
