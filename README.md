# spanner-test

## 環境変数

* bash

```bash
export GITHUB_TOKEN=YOUR_TOKEN
```

* fish

```fish
set -x GITHUB_TOKEN YOUR_TOKEN
```

## 実行

```bash
docker-compose up -d spanner
docker-compose run spanner-init
docker-compose run create
docker-compose run gen-model
```

## spannerへの接続

```bash
docker-compose run spanner-cli
```
