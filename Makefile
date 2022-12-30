build:
	go get .
	go build -o build/
migrate:
	./build/solana-crawl-service migrate