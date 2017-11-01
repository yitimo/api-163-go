# api-163-go
yitimo's 163 api in golang
## How to run
```
go run main.go
```
or run after build:
```
go build
api-163-go.exe
```

By default it will listen at ``http://localhost:9999``

## Currently APIs
### Search
* ``/search/:words/:page/:limit``

example:
* ``/search/再见二丁目/1/10``

### Downlod
* ``/download/:id``

example:

* ``/download/123456``

Note that the id param should get from search api or song's info api.

### Info

Completing...(the API is arrival)
