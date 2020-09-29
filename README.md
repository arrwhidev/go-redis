# go-redis

Redis-like in memory cache adhering to the [RESP protocol.](https://redis.io/topics/protocol)

This project is an effort to learn more Go. I've toyed with Go over the years but want a large project focused on data structures and networking to dig deep into the language.

Currently supported commands;

 - PING: ✅
 - ECHO: ✅
 - QUIT: ✅
 - SET: ✅ (EX & PX args)
 - GET: ✅ 
 - DEL: ✅ 
 - KEYS: ✅ (* only)

## Running locally

To start the server;

    go run cmd/main.go

To start a redis-cli (replace with local ip);

    docker run --rm --name redis-cli -it goodsmileduck/redis-cli redis-cli -h 192.168.10.1 -p 6379