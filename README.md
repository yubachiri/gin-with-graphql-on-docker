## 社内諸々レビューサービス m-share
公共事業です

## 初回
```
$ cp .env.example .env
```
DB情報は変更なしで動くはず

## 起動
```
$ docker-compose build
$ docker-compose up
```

## 使ってるもの
- Go 1.15
- [Gin](https://github.com/gin-gonic/gin) 1.6.3
- MySQL 5.7
