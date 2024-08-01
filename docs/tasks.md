# Tasks

## In Scope

以下のタスクをスコープ内として優先度順に実施する. タイムボックスにより実施しないタスクがある

### Must

- [x] スコープの定義
- [x] Goの環境構築
- [x] multi modulesディレクトリ構成
- [x] serverをダミーで起動
- [x] CI環境
- [x] ドキュメント作成 設計
- [x] DBテーブル作成
- [x] ユースケース 請求書データ作成 Interface定義(request type, response type)
- [x] ユースケース 請求書データ一覧取得 Interface定義(request type, response type)
- [x] ドキュメント作成(request type, response type) Interface & example
- [ ] ユースケースで利用するドメインオブジェクト作成
- [ ] 認証ダミー
- [ ] トランザクションダミー
- [ ] DBダミー
- [ ] ID生成ダミー
- [ ] DI
- [ ] serverの起動
- [ ] validation ユースケース 請求書データ作成 取引先の存在チェック
- [ ] エラーハンドリング 一部
- [ ] ドメインロジックの単体テスト作成 一部
- [ ] READMEドキュメント作成

### Better

- [ ] ID生成の実装
- [ ] ORMの構造体定義
- [ ] DBトランザクションの実装
- [ ] DBアクセスの実装
- [ ] validation 全部
  - [ ] validation ユースケース 請求書データ作成 各プロパティの事前条件validation
  - [ ] validation ユースケース 請求書データ一覧取得 指定期間の範囲validation
- [ ] エラーハンドリング 全部
- [ ] ドメインロジックの単体テスト作成 全部
- [ ] ローカルDBのdocker作成
- [ ] ローカル環境向けのconfigファイル作成
- [ ] ドキュメント作成 ローカル環境構築

## Out of Scope

- [ ] 認証の実装
- [ ] パブリッククラウド等のインフラ作成
- [ ] 自動デプロイ
- [ ] 監視
- [ ] 環境ごとの設定
