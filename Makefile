run: build
	./bin/expTracker

build:
	go build -o ./bin/expTracker main.go