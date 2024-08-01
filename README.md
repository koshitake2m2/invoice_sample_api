# invoice_sample_api

## Abstract
請求書管理を題材にしたAPIサーバーのサンプルです.

## Documents

- [ディレクトリ・モジュール設計](./docs/structure_design.md)

## Setup

```bash
anyenv install goenv
exec $SHELL -l
goenv install 1.22.5
goenv local 1.22.5
go install github.com/air-verse/air@latest
```

## Run

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
