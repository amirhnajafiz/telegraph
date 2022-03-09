# Telegraph

Telegraph is a web-communicator for sending and receiving messages via **nats** server.

## Nats server

## How to use ?

## Docker
Use the following command to run the project on docker:
```shell
docker compose up -d
```

## Deploy
To deploy the project on a kubernetes cluster, use the following helm commands:<br />
First install all dependencies charts:
```shell
helm dep up ./deploy/telegraph
```

Then launch the project by the following command:
```shell
helm install ./deploy/telegraph
```