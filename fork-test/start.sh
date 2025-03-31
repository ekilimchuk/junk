#!/bin/sh

logfile="./start.log"
pidfile="./start.pid"
path=$PWD
telegraf_user=$USER

/usr/sbin/daemon -fcr -P ${pidfile} -u ${telegraf_user} -o ${logfile} \
    go run ${path}/main.go