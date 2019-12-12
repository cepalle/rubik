all: build

dir:
	mkdir -p build/

build: dir
	go build -o build/rubik cmd/rubik/main.go

fclean:
	rm -rf build

re: fclean all

bfs_gen: dir
	go build -o build/bfs_gen cmd/rubik/bfs_gen.go

learn: dir
	go build -o build/learn cmd/rubik/learn.go
