![ソースコードサイズ](https://img.shields.io/github/languages/code-size/k-jun/techtalk)


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
vegeta attack -rate=2000 -targets requests.txt -duration=10s | vegeta report
```


