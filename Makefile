run:
	go run ./cmd/server/main.go

watch:
	CompileDaemon -command="go run ./cmd/server/main.go" -directory="./cmd/server"