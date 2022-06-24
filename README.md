# todoApp

go :1.18ver

## 概要
- [heroku](https://devcenter.heroku.com/categories/reference)でアプリの公開
  - postgresqlでDBを作成

## heroku手順
① アプリのコードを連携
1. githubでソースコードをcommit
1. heroku login
1. heroku create
1. git push heroku main

②postgresqlを作成
1. heroku-addson:create postfresql:hobby-dev
1. heroku pg:psql // CLIを使用して、postgresにログイン
1. \i example.sql // 用意したsqlを実行
1. .q // postgresからログアウト

【構成】

```
  .
├── README.md
├── app
│   ├── controllers
│   │   ├── route.go
│   │   ├── route_auth.go
│   │   └── server.go
│   ├── models
│   │   ├── base.go
│   │   ├── todo.go
│   │   └── users.go
│   └── views
│       ├── css
│       │   ├── bootstrap.min.css
│       │   ├── bootstrap.min.css.map
│       │   └── form.css
│       ├── js
│       │   ├── bootstrap.bundle.min.js
│       │   ├── bootstrap.bundle.min.js.map
│       │   └── jquery-3.6.0.min.js
│       └── templates
│           ├── index.html
│           ├── layout.html
│           ├── login.html
│           ├── private_navbar.html
│           ├── public_navbar.html
│           ├── signup.html
│           ├── todo_edit.html
│           ├── todo_new.html
│           └── top.html
├── config
│   └── config.go
├── config.ini
├── example.sql
├── go.mod
├── go.sum
├── main.go

```
  
