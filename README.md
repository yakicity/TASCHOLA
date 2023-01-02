# TASCHOLA

TSSCHOLA is a simple todo list manager which can sync with Tokyo Tech T2SCHOLA.

## Set Up

1. Clone this repository
2. `docker-compose up -d`

## For Developers

- Docker Compose について

### Frontend

本プロジェクトにて採用した記述等に関する説明

- React + TypeScript

  Next.js を使用しなかった理由

### Backend

- CORS について

- TablePlus について

- MySQL との接続について

  wait-for-it.sh を使用しているのには理由がある。すくなくとも okoge の環境では、docker compose up で立ち上がる順番が go, mysql となってしまい、go が mysql に接続できないことがあった。そのため、wait-for-it.sh を使用して、mysql が立ち上がるまで待機するようにしている。

  同様の現象は、以下の記事においても確認されている。

  - [docker-compose up で MySQL が起動するまで待つ方法](https://qiita.com/study-toto/items/256c2d306b3c6c8f86cd)
  - [docker-compose で golang と MySQL をつなぐ](https://zenn.dev/ajapa/articles/443c396a2c5dd1)

  この挙動に関しては、もう少し調査したほうがよいと思われる。

### DB

### 参考資料

#### Frontend

#### Backend

- [Go Gin CORS + 認証](https://qiita.com/bty__/items/f8c4393bd7701a1d703c)

#### DB
