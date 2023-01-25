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

Dockerファイルをいじった時
```
docker-compose down
docker images
docker rmi chat_web
docker-compose up -d 
docker-compose build
docker-compose logs
docker-compose ps
```
