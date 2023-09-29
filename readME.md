 docker run --name some-redis -d redis

 docker run -d -p 6379:6379 --name my-redis -e REDIS_PASSWORD= -e REDIS_HOST=0.0.0.0 -e REDIS_DB=0 redis
