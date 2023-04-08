docker pull sameersbn/bind:latest

docker run --name bind -d --restart=always \
  --publish 53:53/tcp \
  --publish 53:53/udp \
  --publish 10000:10000/tcp \
  --volume /home/go/software/dns-bind:/data \
  --env='WEBMIN_INIT_SSL_ENABLED=false' \
  --env='ROOT_PASSWORD=fox233'  sameersbn/bind:latest

docker run -d --name=bind \
  --restart=always \
  --dns=114.114.114.114 \
  --publish 53:53 \
  --publish 10000:10000 \
  --volume /home/go/software/dns-bind:/data \
  --env='ROOT_PASSWORD=fox233' \
  sameersbn/bind:latest


-------------------------------------
docker build -t dnsx:v1 .
docker run -itd --name=dnsx \
--restart=always    \
--publish 53:53/tcp     \
--publish 53:53/udp     \
dnsx:v1
----------------------
docker run -itd --name=dnsx \
--publish 53:53/tcp     \
--publish 53:53/udp     \
dnsx:v1
-----------------------
docker logs dnsx
docker exec -it dnsx /bin/sh

docker update --restart=no < container_id>

docker update --restart=no dnsx
docker stop dnsx
docker rm dnsx && docker rmi dnsx:v1 



