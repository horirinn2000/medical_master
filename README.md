# Medical Master Project

医療マスターデータを管理・提供するためのAPIサーバープロジェクトです。

## API仕様とコード生成

このプロジェクトでは、OpenAPI (Swagger) 仕様書から Go のコードを自動生成しています。
API仕様の管理には `doc/openapi.yaml` を使用しており、これを元にサーバーコードを生成しています。

### 使用ツール

- [oapi-codegen/v2](https://github.com/oapi-codegen/oapi-codegen) (v2.5.1)

### 自動生成の手順

`doc/openapi.yaml` を修正した後、以下のコマンドを実行することで API サーバーのインターフェースおよびモデル定義が更新されます。

```bash
oapi-codegen -package api -generate gin,models,spec doc/openapi.yaml > internal/api/medical_api.gen.go
```

※ `internal/api/medical_api.gen.go` は手動で編集しないでください。

### Swagger UI の確認

API サーバーを起動すると、以下の URL で Swagger UI を確認できます。

- URL: `http://localhost:8080/swagger/index.html`

サーバーの起動方法:
```bash
go run cmd/api/main.go
```

## 構成

- `doc/`: OpenAPI 仕様書 (`openapi.yaml`)
- `internal/api/`: 自動生成された Go コードおよび API 実装
- `cmd/api/`: API サーバーのエントリポイント
- `model/`: データベースモデル定義
- `csv/`: マスターデータのCSVファイル
