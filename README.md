# discord-minecraft-server-stats

a simple discord bot that displays information on a minecraft server

currently: name, current/max players, and player names

uses [mcapi](https://mcapi.us) but also has internal functions for [mcsrvstat](https://api.mcsrvstat.us/)

```text
Usage of ./bot:
  -domain string
        server nickname
  -loop int
        seconds between messages (default 60)
  -status int
        0: playing, 1: listening
  -token string
        discord bot token
```

## docker

### command line

```shell
docker run -e "TOKEN=XXX..XXX" -e "DOMAIN=sdomain.tld" -e "STATUS=0" -e "LOOP=5" ghcr.io/rssnyder/discord-minecraft-server-stats
```

### docker compose

```yaml
---
version: "3"
services:
  discord-minecraft-server-stats:
    image: ghcr.io/rssnyder/discord-minecraft-server-stats
    environment:
      TOKEN: XXX..XXX
      DOMAIN: sdomain.tld
      NAMES: true
      STATUS: 0
      LOOP: 5
```

## command line

### download binary

grab a download link from [here](https://github.com/rssnyder/discord-minecraft-server-stats/releases).
```shell
curl -L https://github.com/rssnyder/discord-minecraft-server-stats/releases/download/v<version>/discord-minecraft-server-stats_<version>_<os>_<arch>.tar.gz -o discord-minecraft-server-stats.tar.gz
tar zxf discord-minecraft-server-stats.tar.gz
```

### run

```shell
./discord-minecraft-server-stats -token "XXX..XXX" -nickname "some nickname" -activity "some activity" -status "0" -refresh "5"
```
