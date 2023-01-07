# TASCHOLA

TSSCHOLA is a simple todo list manager which can sync with Tokyo Tech T2SCHOLA.

## Set Up

Docker 環境が用意されているのは、MySQL(DB)と Go Gin(Backend)のみである。

フロントエンドについては、ローカルで開発するために、yarn を使用している。

アプリケーション全体を試す場合は、Terminal のタブを 2 つ用意し、片方で Backend の起動、もう片方で Frontend の起動を行う。
localhost は共有されているので、フロントエンドとバックエンドは CORS を通じて通信することができる。

### For Backend

1. Move to the directory `cd TASCHOLA`
2. `docker compose up -d`

ログを見たい場合は、`docker compose up`を使用する。
(`docker compose down`をするために Terminal のタブがもう一つ必要)

注意: DB(MySQL) -> API(Go Gin) の順に起動するようになっている。そのため、起動までに時間がかかることが予想される。特に、MySQL は初回起動時に時間がかかる。(それぞれの PC のスペックにもよるが初回起動時は 5〜8 分程度要する場合もありうる。)

注意: もし起動に失敗した場合は、`docker compose down`の後、`docker compose up`としてみよう。エラーメッセージが表示されているはずである。

### For Frontend

1. install yarn

   - For Mac

     ```bash
     brew install yarn
     ```

   - For Windows

     ```bash
     npm install --global yarn
     ```

     brew でも問題ありません

2. Move to the directory `cd frontend`
3. `yarn install`
4. `yarn dev`

## For Developers

### Frontend

- import path alias

  相対パスでの import は、ファイル構造を把握するのが難しくなるため、絶対パスでの import を行う。

  import パスのルートは `frontend/`である。

  参考: https://qiita.com/tatane616/items/e3ee99a181662ad6824b

- tailwind css

  CSS フレームワークとして、tailwind css を使用している。

  参考: https://tailwindcss.com/docs/guides/nextjs
  参考: [Next.js に Tailwind CSS を導入する](https://fwywd.com/tech/next-tailwind)

### Backend

- CORS について

  CORS 自体についての説明は以下

  https://developer.mozilla.org/ja/docs/Web/HTTP/CORS

  `backend/router/router.go`にて実装してある。
  詳細な設定方法は、ドキュメント参照のこと

  https://pkg.go.dev/github.com/gin-contrib/cors?utm_source=godoc#Config

  注意点として、CORS の設定は routing の前に行う必要がある。

  そのため、

  ```go
  // health check
  engine.GET("/health", controllers.HealthCheck)
  ```

  の前に、CORS の設定を書く必要がある。

- TablePlus について

  DB の状態を確認する目的や、ダミーデータを作成する目的で TablePlus を使用している。

  接続方法は、`Create a new connection` の後以下のように入力すればよい。
  Password は、docker-compose.yml に記載されているものを使用する。
  (`MYSQL_ROOT_PASSWORD`のところを参照)

  ![TablePlus](./public/tablepuls.png)

- MySQL との接続について

  `wait-for-it.sh`を用いて、MySQL が起動するまで待機するようにしていたが、`healthcheck:`を利用することにより、`docker-compose.yml`だけで依存関係と実行順番制御が行える。

  - heath check について (死活監視)

    Go Gin においては、`/health`エンドポイントを作成し、問題なく Gin が起動している場合は`200`を返すようにしている。

    また、MySQL については mysqladmin ping で問題なく接続できるかいなかを確認している。

  - health check の詳細

    `docker-compose.yml`の `healthcheck:`, `test:` に実際に確かめるために使用しているコマンドが記載されている。

    また、health check の周期なども`docker-compose.yml`で設定している。

### DB

- Set Up

1. `docker compose up -d`
2. install sql-migrate (下記の 1. Installation を参照)
3. `sql-migrate status`で、現在の DB の状態を確認する
4. `sql-migrate up`で、migration を実行する

`sql-migrate status`の具体例

![sql-migrate status](./public/migration.png)

- Migration

DB の中身を変更することがあると思うので、Migration を楽にするツールを導入した。

sql-migrate: https://github.com/rubenv/sql-migrate

1. Installation

   ```bash
   go get -u github.com/rubenv/sql-migrate/...
   ```

2. Check Status

   ```bash
   sql-migrate status
   ```

3. Create Migration File

   ```bash
   sql-migrate new <migration_name>
   ```

4. Apply Migration

   ```bash
   sql-migrate up
   ```

5. Rollback Migration

   ```bash
   sql-migrate down
   ```

### 参考資料

#### Frontend

#### Backend

- [Go Gin CORS + 認証](https://qiita.com/bty__/items/f8c4393bd7701a1d703c)
- [docker-compose におけるヘルスチェック](https://qiita.com/hichika/items/9b96634d471246359e66)
- [Go Gin における CORS の設定](https://ti-tomo-knowledge.hatenablog.com/entry/2020/06/15/213401)

#### DB

- [Go 製マイグレーションツール sql-migrate](https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0)
- [Go の migration ツールのデファクトってなくないですか？](https://onemuri.space/note/is3ev1-d1/)
- [sql-migrate の使い方](https://k2ss.info/archives/3342/)
