# Bookmark API Project

One of the project of golang [list](https://github.com/diaszakir/Backend-Projects/)

## Description

Main idea of this project is creating and managing bookmarks with name, URL, and tag.

You send request with JSON bookmark data and API stores it with a generated ID.

```json
{
    "name": "GitHub",
    "url": "https://github.com",
    "tag": "dev"
}
```

Bookmarks can be filtered by name or tag using query parameters.

## How to launch
Copy repository to your local machine using command:

```
git clone https://github.com/diaszakir/bookmark
```

In case if you do not have go.mod and go.sum

```
go get -u github.com/gin-gonic/gin
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

After

```
go run cmd/app/main.go
```

Loading documentation in case if you changed structure

```
swag init -g cmd/app/main.go
```

To access swagger go to `localhost:8080/swagger/index.html`

## API Endpoints
- `GET /health` - checks API
- `GET /bookmark` - getting all bookmarks (filter by `?name=` or `?tag=`)
- `GET /bookmark/:id` - getting specific bookmark
- `POST /bookmark` - creating a bookmark
- `PUT /bookmark/:id` - updating a bookmark
- `DELETE /bookmark/:id` - deleting specific bookmark
- `DELETE /bookmark` - deleting all bookmarks
