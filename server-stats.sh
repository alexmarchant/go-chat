#!/bin/sh

ssh root@162.243.248.235 <<'ENDSSH'
  pid=$(pidof go-chat)
  ps -p $pid -o %cpu,%mem,cmd
ENDSSH
