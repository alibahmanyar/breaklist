frontend_API_URL := /api

.PHONY: clean setup build-all

setup:
	cd ./frontend/breaklist; npm install
	
clean:
	rm -rf build

build-all: build/reportGenerator build/webserver build/frontend


build/webserver: ./backend/*
	cd ./backend; go build -o ../build/webserver .


build/reportGenerator: ./reportGenerator/*
	cd ./reportGenerator; go build -o ../build/reportGenerator .; cp -r weathercodes template.html ../build/
	cp ./reportGenerator/.env.example build/.env
	- cp ./reportGenerator/.env build/.env

build/frontend: ./frontend/breaklist/src/routes/* ./frontend/breaklist/* ./frontend/breaklist/src/* 
	mkdir -p tmp; cp ./frontend/breaklist/src/routes/+page.svelte ./tmp/+page.svelte.backup
	sed -i 's,http://localhost:3000/api,$(frontend_API_URL),g' ./frontend/breaklist/src/routes/+page.svelte
	cd ./frontend/breaklist; npm run build; cp -r ./build/ ../../build/static/
	cp ./tmp/+page.svelte.backup ./frontend/breaklist/src/routes/+page.svelte
	rm -rf ./tmp