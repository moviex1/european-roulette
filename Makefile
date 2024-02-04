run:
	go run main.go
build:
	go build -o bin/eu_roulette main.go

compile:
	echo "Building for Linux..."
	GOOS=linux GOARCH=arm64 go build -o bin/eu_roulette-linux-arm64 main.go
	echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/eu_roulette-windows-amd64 main.go
