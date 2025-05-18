# Golangで実装するAPIサーバーのテンプレリポジトリ

## DBマイグレーション手順
1. `schema.sql`を更新
2. `schema.sql`からスキーママイグレーション用のSQLファイルを生成（create_user_schemaは任意に変更可能）
    ```bash
    atlas migrate diff add_user_schema --env "local" --config "file://db/atlas.hcl" --to "file://db/schema.sql"

    # 以下でも実行可能
    atlas migrate diff --to "file://db/schema.sql" --dev-url "docker://postgres/15/dev" --dir "file://db/migrations"
    ```
3. マイグレーションを実行
    ``` bash
    atlas migrate apply --env "local" --config "file://db/atlas.hcl"
    ```

## 参考
- https://github.com/budougumi0617/go_todo_app/tree/main
- https://qiita.com/Imamotty/items/3fbe8ce6da4f1a653fae
