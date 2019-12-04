all: build

dir:
	mkdir -p build/

build: dir
	go build -o build/rubik cmd/rubik/main.go

fclean:
	rm -rf build
