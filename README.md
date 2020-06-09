## environment valiables

- REDIS_ENDPOINT
- DB_HOST
- DB_USER
- DB_NAME
- DB_PASSWORD

```sh
export DB_USER="root" \
export DB_PASSWORD="password1!" \
export DB_HOST="localhost" \
export DB_NAME="mysqldb" \
export REDIS_ENDPOINT="localhost:6379"
```

## apis

in this case this is for load testing, if you send channel_id=0, the server fills the id with random number which is 1 to 1000

* GET "channels/{channel_id}/messages"
* POST "channels/{channel_id}/messages"
```json
{
  "user_id": "1",
  "type": "sample type",
  "body": "sample body"
}
```

* PUT "channels/{channel_id}/messages"
```json
{
  "id": "1",
  "type": "updated type",
  "body": "updated body"
}
```

* DELETE "channels/{channel_id}/messages"
```json
{
  "id": "1"
}
```


## TODOs

* change user in Dockerfile from root to application specific user for sequlity
* the app doesn't wait untill the db ready


## Content


* vegeta(https://github.com/tsenart/vegeta)
```
brew install vegeta
echo "GET http://18.180.59.175/channels/1/messages" | vegeta attack -rate=2000 -duration=10s | vegeta report
```

* Siege(https://github.com/JoeDog/siege)

```sh
brew install siege
# bench mark to decide concurrency and requests
siege -b -t 10s  http://18.180.59.175/channels/0/messaga
# concurrency and request per user
siege -c 20 -r 100  http://18.180.59.175/channels/0/messagas
```

* Apach Bench

```sh
ab -c 20 -n 10000  http://18.180.59.175/channels/1/messages
```


* wrk

```sh
brew install wrk
wrk -c 100 -d 5s http://18.180.59.175/channels/1/messages
```

* goku

```sh
brew install k-nasa/tap/goku
goku kamehameha -c 50 -n 1000 http://13.230.60.181/channels/1/message
```

https://qiita.com/jun2014/items/121f2dcb2de4c4e77421
>どんな負荷テスト想定しているかで変わります。
>「大規模負荷テスト」「サーバ負荷テスト」「処理単位の応答時間」
>高価なものから無償のものまでです。負荷テストツールの選定も大事な負荷テスト計画となりますね。



## tool list
- JMeter(http://jmeter.apache.org/) 
- k6(https://app.k6.io/projects/3494089)
https://qiita.com/navitime_tech/items/277fde79adbba3d15217
- gatling
- Tsung
- Locust
https://qiita.com/sho7650/items/58ec4ab4adc9f6b1129d
- loader(https://loader.io/)
https://qiita.com/furu8ma/items/fb7a45388bfe323b46c1
- wrk
- vegeta(https://github.com/tsenart/vegeta)
- artillery(https://artillery.io/)

- blazemeter(https://www.blazemeter.com/)



- gostress
https://github.com/karupanerura/gostress



