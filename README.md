# TASCHOLA

TSSCHOLA is a simple todo list manager which can sync with Tokyo Tech T2SCHOLA.

## Set Up

1. Clone this repository
2. `docker-compose up -d`

## For Developers

### Frontend

### Backend

- CORS について

- TablePlus について

- MySQLとの接続について

  wait-for-it.sh を使用しているのには理由がある。すくなくともokogeの環境では、docker compose upで立ち上がる順番が go, mysqlとなってしまい、goがmysqlに接続できないことがあった。そのため、wait-for-it.shを使用して、mysqlが立ち上がるまで待機するようにしている。

  同様の現象は、以下の記事においても確認されている。
  - [docker-compose upでMySQLが起動するまで待つ方法](https://qiita.com/study-toto/items/256c2d306b3c6c8f86cd)
  - [docker-compose でgolangとMySQLをつなぐ](https://zenn.dev/ajapa/articles/443c396a2c5dd1)

  この挙動に関しては、もう少し調査したほうがよいと思われる。

### DB

