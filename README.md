# Regisys Server

Backend Register system for MMA Junk-Ichi 2017

## サーバー起動

* 必要に応じて`jwtkey/{jwt.key,jwt.key.pub}`を差し替え
* portは`8080`でlisten
* docker-composeのデフォルトは`8080:8080`

	$ docker-compose -f docker-compose.prod.yml run --rm app go-wrapper download
	$ docker-compose -f docker-compose.prod.yml up -d
	
configはdocker-compose.prod.ymlを修正

## 開発環境

depが必要。リポジトリはGOPATH以下にcloneする
    $ go get -u github.com/golang/dep/cmd/dep
	$ dep ensure

### 開発用サーバー起動

	$ docker-compose run --rm app go-wrapper download
	$ docker-compose up -d

### 概略

* `/app, /client, /models, /tool`は`/design`のDSLをもとに生成されるため、編集しない
* `/design`以下を編集した場合は`make generate`で生成する
* 認証にはBasicと [jwt](https://jwt.io/) を使用

#### ドキュメント

1. ブラウザでサーバー (デフォルトで`localhost:8080`)にアクセスし、`http://petstore.swagger.io/v2/swagger.json`を`http://localhost:8080/swagger.json`に入れ替えて`Explore` でswagger-uiが起動する

1. JWT`/token`のAuthorizationに`<username>:<password>`をbase64エンコードして先頭に`Basic `をつけたもの(ex. `Basic Zm9vOnBhc3N3b3JkCg==`)を入れて「Try it out!」

1. `Response Headers`の`authorization`をコピー(ex.`Bearer Bearer eyJhbGciOiJ......`)して、ページの一番上の「Authorize」をクリックし、「Api key authorization」の`value`にペースとし、「Authorize」

1. 各サービスにアクセスできる
