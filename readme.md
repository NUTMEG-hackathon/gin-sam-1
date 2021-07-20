# git cloneしたら
`docker-compose build`

# 立ち上げ
`docker-compose up`
`https://localhost:8080`へアクセスするとhello worldがでる．

# route
## index
`https://localhost:8080/users/index`でuser一覧が得られる
現在はupするたびにシードが作られる状態になっているので，upするたびにサンプルデータが作られるようになってしまっている．

## delete
`http://localhost:8080/users/delete/:id`をするとidのuserを削除する

# dbのデータ削除
`docker-compose down -v`をするとdbのデータを削除する


