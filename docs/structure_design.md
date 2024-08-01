# Design

## Directories & Modules Structure

### Overview

- ディレクトリ構成は [golang-standards/project-layout](https://github.com/golang-standards/project-layout) を参考にしています.
- モジュール構成はレイヤードアーキテクチャを採用しています.

### Details

```plaintext
.
├── build
├── cmd                
│   └── api
├── docs
├── internal
│   ├── base
│   └── invoice
│       ├── di
│       ├── domain
│       ├── infra
│       ├── presentation
│       └── usecase
├── migrations
└── scripts
```

- build: コンテナやビルドに関する設定ファイルを配置
- cmd: 実行可能なプロジェクト
  - api: APIサーバー
- docs: ドキュメントを配置
- internal: 内部で利用するプログラム
  - base: 業務のコンテキストを含まない処理
  - invoice: 請求書管理のコンテキストに関する処理
    - di: DIコンテナ
    - domain: ドメイン層 業務で扱う言葉によるデータやロジックを配置
    - infra: インフラ層 DBや外部APIとの通信, 技術的な詳細を扱う
    - presentation: プレゼンテーション層 クライアントとの通信に関する処理を行う
    - usecase: ユースケース層 ドメインオブジェクトの受け渡しや業務ロジックを起動する処理を配置
- migrations: DBのマイグレーションファイルを配置
- scripts: ローカル環境やCI向けに利用するスクリプトを配置

