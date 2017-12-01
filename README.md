# Go API documents demo

```
docker-compose up
```

## Demo

### Test

```
dredd
```

```
go test ./...
```

### Go API run

```
realize start
```

### DB migration

```
goose -dir ./migrations mysql 'root@(db:3306)/sample' status
goose -dir ./migrations create add_tag_table sql
```

### Dcoument render

```
aglio -i simple.apib --theme-template triple --server
```
