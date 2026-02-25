# cloudflare-mcp-go

Cloudflare API の MCP (Model Context Protocol) サーバー。Go で実装。

## 必要要件

- Go 1.21 以上
- Cloudflare API トークン

## セットアップ

```bash
go build -o cloudflare-mcp-go .
```

## 使い方

環境変数 `CLOUDFLARE_API_TOKEN` に Cloudflare の API トークンを設定して実行します。

```bash
export CLOUDFLARE_API_TOKEN="your-api-token"
./cloudflare-mcp-go
```

サーバーは stdio トランスポートで動作し、MCP クライアントから接続できます。

## 利用可能なツール

### list_zones

Cloudflare アカウントのゾーン一覧を取得します。

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|---|---|---|---|
| name | string | No | ドメイン名でフィルタ |
| page | int | No | ページ番号 (デフォルト: 1) |
| per_page | int | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |

## MCP クライアント設定例

```json
{
  "mcpServers": {
    "cloudflare": {
      "command": "/path/to/cloudflare-mcp-go",
      "env": {
        "CLOUDFLARE_API_TOKEN": "your-api-token"
      }
    }
  }
}
```
