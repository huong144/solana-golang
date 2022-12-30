
# PROJECT SETUP

Clone project

```bash
git clone git@github.com:sotatek-dev/solana-crawl-services.git
```

# BUILD PROJECT

```bash
cp .env.example .env
go get .
go build -o build
```
or
```bash
make build
```
# MIGRATION DATABASE
```bash
  go run main.go migration
```


# DEPLOY PROJECT

## Crawl to latest

```bash
  go run main.go solana-fetch-new
```
or
```bash
./build/solana-crawl-service solana-fetch-new
 ```

## Crawl old slot

```bash
  go run main.go solana-fetch-old
```
or
```bash
./build/solana-crawl-service solana-fetch-old
 ```