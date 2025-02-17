# go-sandbox

## 起動

```
docker compose build
docker compose up -d
docker logs go-sandbox-container
```

## テストの実行

```
docker-compose run --rm go-sandbox go test -v ./...
```
