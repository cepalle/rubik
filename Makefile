all:
	mkdir -p build/
	go build -o build/rubik cmd/rubik/main.go

fclean:
	rm -rf build
