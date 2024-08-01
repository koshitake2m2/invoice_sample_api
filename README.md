# invoice_sample_api

## Abstract

請求書管理を題材にしたAPIサーバーのサンプルです.

## Documents

以下にドキュメントをまとめています.

- [ディレクトリ・モジュール設計](./docs/structure_design.md)
- [API仕様書](./docs/api.md)
- [タスク管理](./docs/tasks.md)

## Setup

```bash
anyenv install goenv
exec $SHELL -l
goenv install 1.22.5
goenv local 1.22.5
go install github.com/air-verse/air@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Tips

### Run Server

以下でServerを起動できます. `http://localhost:8080` でアクセスできます.

```bash
# Serverを起動します.
cd cmd/api
go run .
```

```bash
# ホットリロードで起動します.
cd cmd/api
air
```

### Format

以下で全てのモジュールのコード整形を行います.

```bash
# コード整形できてないファイルを列挙します.
./scripts/format_check.sh

# コード整形します.
./scripts/format_fix.sh
```

### Test

以下で全てのモジュールのテストを行います.

```bash
./scripts/test.sh
```

### Local DB

以下でローカルのDBの起動・停止・アクセス・マイグレーションを行います.

```bash
# Docker
docker compose -f build/package/postgres/compose.yml up -d
docker compose -f build/package/postgres/compose.yml down

# Envs
DB_NAME=mydb
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password

# Access
psql -p $DB_PORT -U $DB_USER -d $DB_NAME -h $DB_HOST

# Migrate
migrate --path migrations --database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up
```
