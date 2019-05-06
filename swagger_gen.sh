#!/usr/bin/env bash
swagger generate spec -b . -o ./swagger/swagger.json -m

#scp ./swagger/swagger.json root@192.241.163.137:/home/docker/swagger/