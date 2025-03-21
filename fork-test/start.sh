#!/bin/sh

logfile="./start.log"
pidfile="./start.pid"

/usr/sbin/daemon -fcr -P ${pidfile} -u ${telegraf_user} -o ${logfile} \
    go run main.go