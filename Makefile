.PHONY: build

build:
	if [ ! -d "build" ]; then mkdir build; fi
	go build -o build/magobot
	chmod +x build/magobot
