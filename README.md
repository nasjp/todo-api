[![Actions Status](https://github.com/nasjp/todo-api/workflows/CI/badge.svg)](https://github.com/nasjp/todo-api/actions)

# todo api

## run

```sh
go run main.go
```
## curl

```sh
$ curl -s -XPOST -d "{\"title\":\"buy pc\"}" localhost:8080/create | jq .
{
  "id": "1dfb5278-be18-496c-8a8c-7a77ae8ce15c",
  "title": "buy pc"
}

$ curl -s -XPOST -d "{\"title\":\"buy house\"}" localhost:8080/create | jq .
{
  "id": "9139c021-b985-461d-829f-aed68eaa74a5",
  "title": "buy house"
}


$ curl -s localhost:8080/list | jq .
[
  {
    "id": "1dfb5278-be18-496c-8a8c-7a77ae8ce15c",
    "title": "buy pc"
  },
  {
    "id": "9139c021-b985-461d-829f-aed68eaa74a5",
    "title": "buy house"
  }
]
```
