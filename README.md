# api-163-go
Yitimo's 163 api in golang. Study use only.

Notes:
1. this is just a RESTful api, so you still need to have a client.
2. This go api is under ``go-martini``, so after cloning thie repo, you should also ``go get`` the [martini](https://github.com/go-martini/martini)
3. modules in this api are all named by [``魔法少女まどか☆マギカ``](https://bangumi.bilibili.com/anime/2539?from=search&seid=10537413920560567412) : )
## How to run
```
go run main.go
```
or run after build:
```
go build
```

By default it will listen at ``http://localhost:9999``

## About authorization
Currently it's using a martini submodule.
* IP address white list

    only support the IPs configured by ``kyouko``
* Username & password authorization

    currently they're static ``yitimo`` and ``iamyitimo``

So if you want to call this api, just ``clone`` it, update with your auth config and run in your own server.
## Currently APIs
### Search
| param  | desc                         | is needed |
|--------|------------------------------|-----------|
| :words | search words                 | yes       |
| :page  | pagination, start from ``1`` | yes       |
| :limit | item number per page         | yes       |

| path           | desc          |
|----------------|---------------|
| /search        | search song   |
| /search/album  | search album  |
| /search/artist | search artist |

### Downlod
Edit soon ...

### Song
#### Info

#### Lyric

### Artist
#### TopSong

#### Album

