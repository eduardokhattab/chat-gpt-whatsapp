## Running project

You can run the project in docker or locally.

To run in docker:

```sh
docker-compose up -d
docker-compose exec chatservice bash

# change env to DB_HOST=mysql
# and request to localhost:8081
go run cmd/chatservice/main.go
```

To run locally, you need to have the `libtiktoken.a` on your golang path, to do that, you can build the lib with cargo and move to your machine:

In container, the lib will be at: `/go/pkg/mod/github.com/j178/tiktoken-go@v0.2.1/tiktoken-cffi/target/release/libtiktoken.a`. Just copy it and paste:

- If you are running with pkg: `/home/gympasser/projects/pkg/mod/github.com/j178/tiktoken-go@v0.2.1/tiktoken-cffi/target/release/libtiktoken.a`
- If you are running with vendor: `/home/gympasser/projects/chat-gpt-whatsapp/chatservice/vendor/github.com/j178/tiktoken-go/tiktoken-cffi/target/release/libtiktoken.a`

- change env to `DB_HOST=localhost`
- and request to `localhost:8080`
