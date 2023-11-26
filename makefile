.PHONY: clean setup all release

all: build/reportGenerator build/webserver build/static

release: all
	goreleaser release --snapshot --clean

setup:
	cd ./frontend/breaklist; npm install

clean:
	rm -rf build
	rm -rf dist


build/webserver: ./backend/*
	cd ./backend; go mod tidy; go build -ldflags '-w -s' -o ../build/webserver .

build/reportGenerator: ./reportGenerator/*
	cd ./reportGenerator; go mod tidy; go build -ldflags '-w -s' -o ../build/reportGenerator .; cp -r weathercodes template.html ../build/
	cp ./reportGenerator/.env.example build/.env.example

build/static: ./frontend/breaklist/src/* ./frontend/breaklist/src/routes/* ./frontend/breaklist/static/* ./frontend/breaklist/*
	cd ./frontend/breaklist; npm run build; cp -r ./build ../../build/static/