setup:
	cd ./frontend/breaklist; npm install
	
clean:
	rm -rf build

build: *
	mkdir -p build
	cd ./backend; go build -o ../build/backend/ .
	cd ./reportGenerator; go build -o ../build/reportGenerator/ .; cp -r weathercodes wkhtmltopdf ../build/reportGenerator/
	cd ./frontend/breaklist; npm run build; cp -r ./build/ ../../build/frontend/