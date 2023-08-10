frontend_API_URL := /api

.PHONY: clean setup build-all

setup:
	cd ./frontend/breaklist; npm install
	
clean:
	rm -rf build

build-all: build/reportGenerator/reportGenerator build/backend/backend build/frontend


build/backend/backend: ./backend/*
	cd ./backend; go build -o ../build/backend/ .

build/reportGenerator/reportGenerator: ./reportGenerator/*
	cd ./reportGenerator; go build -o ../build/reportGenerator/ .; cp -r weathercodes wkhtmltopdf ../build/reportGenerator/

build/frontend: ./frontend/breaklist/src/routes/* ./frontend/breaklist/* ./frontend/breaklist/src/* 
	mkdir -p tmp; cp ./frontend/breaklist/src/routes/+page.svelte ./tmp/+page.svelte.backup
	sed -i 's,http://localhost:3000,$(frontend_API_URL),g' ./frontend/breaklist/src/routes/+page.svelte
	cd ./frontend/breaklist; npm run build; cp -r ./build/ ../../build/frontend/
	cp ./tmp/+page.svelte.backup ./frontend/breaklist/src/routes/+page.svelte
	rm -rf ./tmp