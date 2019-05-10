# gate
Mango Web: Gate

Gate enables the use of subdomains and ssl at development time, and also acts as the gateway for all incoming requests.
Access_Tokens are controlled by the Gate. We create a cookie which is transformed into a Authorization header for each request.

## Run with Docker
```
>$ docker build -t avosa/gate:dev .
>$ docker rm GateDEV
>$ docker run -d -e RUNMODE=DEV -p 80:80 -p 443:443 --network mango_net --name GateDEV -v `pwd`/certs/:/certs avosa/gate:dev
>$ docker logs GateDEV
```

## Note on Symlinks
When running your container on a production environment, you may come across symlinks
You should then rather mount the fullchain.pem and privkey.pem individually using;
-v $(readlink -f `pwd`/certs/fullchain.pem):/certs/fullchain.pem -v $(readlink -f `pwd`/certs/privkey.pem):/certs/privkey.pem