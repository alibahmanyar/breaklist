.PHONY: clean setup all

all: build/reportGenerator build/webserver build/static

setup:
	cd ./frontend/breaklist; npm install

clean:
	rm -rf build


build/webserver: ./backend/*
	cd ./backend; go build -ldflags '-w -s' -o ../build/webserver .

build/reportGenerator: ./reportGenerator/*
	cd ./reportGenerator; go build -ldflags '-w -s' -o ../build/reportGenerator .; cp -r weathercodes template.html ../build/
	cp ./reportGenerator/.env.example build/.env
	- cp ./reportGenerator/.env build/.env

build/static: ./frontend/breaklist/src/* ./frontend/breaklist/src/routes/* ./frontend/breaklist/static/* ./frontend/breaklist/*
	cd ./frontend/breaklist; npm run build; cp -r ./build/ ../../build/static/