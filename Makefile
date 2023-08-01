hello:
	echo "Hello"

buildwin:
	GOOS=windows GOARCH=amd64 go build -o bin/wbuddy.exe main.go

buildlinuxamd64:
	GOOS=linux GOARCH=amd64 go build -o bin/wbuddy main.go

run:
	go run main.go Tokyo
