.PHONY: build, run

build:
	cd front && yarn build
	go build -o ./bin/blive
