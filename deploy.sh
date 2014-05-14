#!/bin/sh

ssh root@162.243.248.235 <<'ENDSSH'
  pid=$(pidof go-chat)
  kill -15 $pid
  cd /var/www/go-chat
  git pull origin master
  go build
  ./go-chat & disown
ENDSSH