# SimpleAccountBook

## API requests
### register
#### request
``` shell
curl --location --request POST 'http://127.0.0.1:8080/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "username",
    "password": "password"
}'
```
#### response
```json
{
    "status_code": 0,
    "status_message": " success"
}
```

### login
#### request
``` shell
curl --location --request POST 'http://127.0.0.1:8080/user/login/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "string",
    "password": "string"
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string",
    "expire": "string"
}
```

### refresh token
#### request
``` shell
curl --location --request POST 'http://127.0.0.1:8080/user/refresh/' \
--header 'Authorization: <Authorization>' \
```
#### response
```json
{
    "status_code": "string",
    "status_msg": 0,
    "token": "string",
    "expire": "string"
}
```

### get user information
#### request
``` shell
curl --location --request GET 'http://127.0.0.1:8080/user/' \
--header 'Authorization: <Authorization>' \
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string",
    "user": {
        "user_id": 0,
        "username": "string"
    }
}
```

### create account
#### request
```shell
curl --location --request POST 'http://127.0.0.1:8080/account/' \
--header 'Authorization: <Authorization>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_name": "string",
    "balance": 0
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### get account list
#### request
```shell
curl --location --request GET 'http://127.0.0.1:8080/account/list/' \
--header 'Authorization: <Authorization>' 
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string",
    "data": [
        {
            "account_id": 0,
            "account_name": "string",
            "balance": 0
        }
    ]
}
```

### modify account information
#### request
```shell
curl --location --request PUT 'http://127.0.0.1:8080/account/<account_id>/' \
--header 'Authorization: <Authorization>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_name": "string",
    "balance": 0
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### delete account
#### request
```shell
curl --location --request PUT 'http://127.0.0.1:8080/account/<account_id>/' \
--header 'Authorization: <Authorization>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_name": "string",
    "balance": 0
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### delete account
#### request
```shell
curl --location --request DELETE 'http://127.0.0.1:8080/account/<account_id>/' \
--header 'Authorization: <Authorization>'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### get account information
#### request
```shell
curl --location --request GET 'http://127.0.0.1:8080/account/<account_id>/' \
--header 'Authorization: <Authorization>'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string",
    "data": {
        "account_id": 0,
        "account_name": "string",
        "balance": 0
    }
}
```

### create bill
#### request
```shell
curl --location --request POST 'http://127.0.0.1:8080/bill/<account_id>/' \
--header 'Authorization: <Authorization>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": 0,
    "amount": 0,
    "category": "string",
    "comment": "string"
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### delete bill
#### request
```shell
curl --location --request DELETE 'http://127.0.0.1:8080/bill/<bill_id>' \
--header 'Authorization: <Authorization>'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### modify bill information
#### request
```shell
curl --location --request PUT 'http://127.0.0.1:8080/bill/1' \
--header 'Authorization: <Authorization>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 0,
    "account_id": 0,
    "type": 0,
    "amount": 0,
    "category": "string",
    "comment": "string"
}'
```
#### response
```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

### get bill information
#### request
```shell
curl --location --request GET 'http://127.0.0.1:8080/bill/<bill_id>/' \
--header 'Authorization: <Authorization>'
```
#### response
```json
{
  "status_code": 0,
  "status_msg": "string",
  "data": {
    "bill_id": 0,
    "account_id": 0,
    "type": 0,
    "amount": 0,
    "category": "string",
    "comment": "string"
  }
}
```

### list bills
#### request
```shell
curl --location --request GET 'http://127.0.0.1:8080/bill/list/<account_id>/' \
--header 'Authorization: <Authorization>'
```
#### response
```json
{
  "status_code": 0,
  "status_msg": "string",
  "data": [
    {
      "bill_id": 0,
      "account_id": 0,
      "type": 0,
      "amount": 0,
      "category": "string",
      "comment": "string"
    }
  ]
}
```