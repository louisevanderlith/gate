# gate
Mango Web: Gate

Gate enables the use of subdomains and ssl at development time, and also acts as the gateway for all incoming requests.
Access_Tokens are controlled by the Gate. We create a cookie which is transformed into a Authorization header for each request.

The Gate application acts as the main entry point for all applications, as we can assign names to every registered application.
So, when debugging we can call 'https://comms.localhost(:80)/' instead of 'http://localhost:8085' for the 'Comms.API'.
This avoids the need to remember the specific port for every application running and also removes the need to store URLs within every application.
The front-end also relies on the Gate to determine the correct environment's URL when calling services.

The gate application will try to load the default WWW website when a subdomain can't be found.
We require 1(one) instance of Gate running for every environment we have.

## Run with Docker
```
>$ docker build -t avosa/gate:dev .
>$ docker rm GateAPI
>$ docker run -d -e KEYPATH=/certs/ -e PUBLICKEY=fullchain.pem -e PRIVATEKEY=privkey.pem -p 80:80 -p 443:443 --network mango_net --name GateAPI -v `pwd`/certs/:/certs avosa/gate:dev
>$ docker logs GateAPI
```

## Note on Symlinks
When running your container on a production environment, you may come across symlinks
You should then rather mount the fullchain.pem and privkey.pem individually using;
-v $(readlink -f `pwd`/certs/fullchain.pem):/certs/fullchain.pem -v $(readlink -f `pwd`/certs/privkey.pem):/certs/privkey.pem