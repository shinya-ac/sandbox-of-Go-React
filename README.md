Goもモジュール（パッケージ）はモジュールモードで開発をする
（モジュールモードとGOPATHモードの違いは(こちら)[https://qiita.com/fetaro/items/31b02b940ce9ec579baf#%E6%8E%A2%E3%81%97%E3%81%AB%E8%A1%8C%E3%81%8F%E5%A0%B4%E6%89%80%E3%81%AF%E3%81%A9%E3%81%93%E3%81%8B-2]を参考に。）

ディレクトリに以下のファイルとフォルダがあるのを確認
```
chat
├── docker-compose.yml
├── app
│   ├── Dockerfile
│   └── src
│       ├── article
│       │   └── article.go
│       ├── go.mod
│       ├── go.sum
│       └── main.go
└── mysql
    ├── .env
    ├── Dockerfile
    └── init
        └── create_table.sh
```

`docker network create golang_1q1a_app_network`を行ってネットワークを手動で作成する
create_table.shファイルのパーミッションを645にする（末尾を5にすることでこのテーブルを作成してくれるシェルをgoが実行できるようにする）

参考：https://zenn.dev/ajapa/articles/443c396a2c5dd1

```
go mod tidy

docker-compose up -d
docker-compose build
```


アプリコンテナの入り方
```
docker psでアプリコンテナのIDを確認
docker exec -it 6af203f92b0b /bin/bashでコンテナに入る
```

Goのアプリファイルを編集したのちはアプリコンテナをリスタートする
`docker-compose restart 1q1a_app`

Dockerファイルをいじった時
```
docker-compose down
docker volume rm mysql_1q1a_app_volume
docker images　イメージの確認
docker rmi 1q1a_web
docker rmi 1q1a_db 
docker rmi 1q1a_client
docker-compose build --no-cache
docker-compose up
以下は参考(起動状態の確認)
docker-compose logs
docker-compose ps
```

ログインする際のリクエストは以下のようにする
ログイン機能には、以下のようなリクエストを送信することでログインできます。
* HTTPメソッド: POST
* URL: /login
* ボディ:perl
Copy code{ "email": "example@example.com", "password": "password" } 
email: ログインに使用するメールアドレス
* password: ログインに使用するパスワード
* コンテントタイプ: application/json
このリクエストを送信すると、サーバーは認証処理を行い、正常にログインできればアクセス用のトークンを返却する。

Reactにパッケージをインストールする方法
```
まず安定バージョンを調べる
miyoshishinya@miyoshishinyanoMacBook-Air-2 client % npm view axios version  
1.2.6

package.jsonにそのバージョンで記載する
...
"dependencies": {
    "axios": "1.2.6",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    ...

からのnpm install
miyoshishinya@miyoshishinyanoMacBook-Air-2 client % npm install           

added 9 packages, and audited 120 packages in 895ms

9 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
```