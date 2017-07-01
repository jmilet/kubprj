#!/bin/bash

docker run -d --rm --name countersrv1 jmilet/countersrv
docker run -d --rm --name countersrv2 jmilet/countersrv
docker run -d -it --link countersrv1 --link countersrv2 --name nginx --rm -p 8080:80 -v $HOME/kubprj/arch/nginx.cf:/etc/nginx/nginx.conf nginx
