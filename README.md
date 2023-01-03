# TASCHOLA

TSSCHOLA is a simple todo list manager which can sync with Tokyo Tech T2SCHOLA.

## Set Up

1. Clone this repository
2. `docker compose up -d`

注意: DB(MySQL) -> API(Go Gin) -> Frontend(React) の順に起動するようになっている。そのため、起動までに時間がかかることが予想される。特に、MySQL は初回起動時に時間がかかる。(それぞれの PC のスペックにもよるが初回起動時は 5〜8 分程度要する場合もありうる。)

注意: もし起動に失敗した場合は、`docker compose down`の後、`docker compose up`としてみよう。エラーメッセージが表示されているはずである。

## For Developers

### Frontend

本プロジェクトにて採用した記述等に関する説明

- React + TypeScript

  Next.js を使用しなかった理由

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

### 参考資料

#### Frontend

#### Backend

- [Go Gin CORS + 認証](https://qiita.com/bty__/items/f8c4393bd7701a1d703c)
- [docker-compose におけるヘルスチェック](https://qiita.com/hichika/items/9b96634d471246359e66)
- [Go Gin における CORS の設定](https://ti-tomo-knowledge.hatenablog.com/entry/2020/06/15/213401)

#### DB
