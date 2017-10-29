# Go API documents demo

## Memo

### Installation

```
npm install -g aglio
npm install -g drakov
npm install -g dredd
```

```
go get -u github.com/labstack/echo
go get -u github.com/snikch/goodman/cmd/goodman
```

```
dredd init
```

### JSON Schema generator

- https://jsonschema.net/#/editor

### Demo

#### API Test

```
dredd
```

#### Go API run

```
realize start
```




```

```

```
json-schema-generator json.json -o schema.json

aglio -i api.md --theme-template triple --server
```
