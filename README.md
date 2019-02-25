# gate
Mango Web: Gate

Gate enables the use of subdomains and ssl at development time, and also acts as the gateway for all incoming requests.

## Run with Docker
*$ go build
*$ docker build -t avosa/gate:dev .
*$ docker rm gateDEV
*$ docker run -d --network host --name gateDEV -v `pwd`/certs/:/certs avosa/gate:dev
  ## Windows
docker run -d --network host --name gateDEV -v D:/Projects/Sandbox/stuff/gate/certs/:/certs avosa/gate:dev
*$ docker logs gateDEVdoc