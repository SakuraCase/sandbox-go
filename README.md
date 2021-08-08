## cmd

### コンテナの立ち上げ

```
docker-compose up -d
```

### 作成したファイルのフォーマット & 実行

```
docker-compose exec app go fmt main.go && docker-compose exec app go run main.go
```

### コンテナを閉じる

```
docker-compose down
```

## 参考

- [環境構築](https://zenn.dev/tomi/articles/2020-10-14-go-docker)

- [はじめての Go](https://gihyo.jp/dev/feature/01/go_4beginners)

- [公式 Tutorials](https://golang.org/doc/tutorial/)
