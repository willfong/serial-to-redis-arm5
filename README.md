# serial-to-redis

A simple Go app to read from an Arduino serial port and send the data to Redis.

```
docker run \
    -d --name serial_to_redis \
    --net aqua \
    --device=/dev/ttyACM0 \
    --restart unless-stopped \
    wfong/serial-to-redis
```

