# spanner-test

## 環境変数

* bash

```bash
export SPANNER_EMULATOR_HOST=localhost:9010
export CREDENTIALS=PATH/TO/firebase-auth.json
```

* fish

```fish
export SPANNER_EMULATOR_HOST=localhost:9010
export CREDENTIALS=PATH/TO/firebase-auth.json
```

## 実行

```bash
docker-compose up spanner
docker-compose run spanner-init
docker-compose run create
```

## spannerへの接続

```bash
docker-compose run spanner-cli
```
