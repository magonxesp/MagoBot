.PHONY: clean build build-docker-image

clean:
	rm MagoBot

build:
	go build

build-docker-image:
	bash scripts/build-docker-image.sh
