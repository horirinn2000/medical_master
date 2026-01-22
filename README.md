# Medical Master Project

医療マスター（レセプト電算処理システム用）データを管理・提供するためのAPIサーバーおよびインポートバッチプロジェクトです。

## 概要

厚生労働省が公開している各種マスターファイルを解析し、PostgreSQLデータベースへの登録およびAPI経由でのデータ提供を行います。

### サポートしているマスター
- 医科診療行為マスター（および補足テーブル）
- 歯科診療行為マスター（および補足テーブル）
- 特定器材マスター
- 医薬品マスター（およびHOTコード）
- 傷病名マスター
- コメントマスター（および関連情報）
- 調剤行為マスター
- 訪問看護マスター
- 歯式マスター、病棟マスター

## 技術スタック
- **Language**: Go 1.21+
- **Framework**: Gin (API), GORM (ORM)
- **Database**: PostgreSQL
- **API Spec**: OpenAPI 3.0

## 構成

- `cmd/api/`: APIサーバーのエントリポイント
- `cmd/batch/`: インポート用バッチのエントリポイント
- `doc/`: OpenAPI 仕様書 (`openapi.yaml`) および仕様書PDF
- `internal/api/`: 自動生成された Go コード
- `internal/handler/`: API ハンドラー（ビジネスロジック）
- `internal/model/`: データベースモデル定義
- `internal/batch/`: マスター解析・インポートロジック
- `csv/`: マスターデータ（CSV/TXT）

## 使い方

### 1. 環境設定
`.env.example` をコピーして `.env` ファイルを作成し、データベース接続情報等を設定します。
```bash
cp .env.example .env
vi .env
```

### 2. データのインポート（バッチ実行）
以下のコマンドを実行してマスターデータをインポートします。
```bash
go run cmd/batch/main.go
```

### 3. APIサーバーの起動
```bash
go run cmd/api/main.go
```
起動後、以下の URL で Swagger UI を確認できます。
- URL: `http://localhost:8080/swagger/index.html`

## 開発者向け情報

このプロジェクトの開発に関する詳細なガイドライン（ディレクトリ構造の詳細、命名規則、API更新フロー等）については、**[AGENTS.md](./AGENTS.md)** を参照してください。

### APIコードの自動生成
`doc/openapi.yaml` を修正した後、以下のコマンドを実行してコードを更新します。
```bash
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest --config oapi-codegen.yaml doc/openapi.yaml
```
