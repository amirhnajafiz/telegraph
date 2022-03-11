# Telegraph

Telegraph is a web-communicator for sending and receiving messages via **nats** server.

## Nats server

## How to use ?
#### Join/Register to a room to begin the communications
url:
```shell
/api/join
```

method:
```shell
POST
```

request:
```json
{
  "username": "[your name]",
  "password": "[your password]"
}
```

response:
```json
{
  "token": "[JWT token]"
}
```

#### Send a message 
url:
```shell
/api/publish
```

method:
```shell
POST
```

header:
```json
{
  "jwt-token": "[the jwt token you got from join]"
}
```

request:
```json
{
  "sender": "[your application name]",
  "message": "[message you want to send]"
}
```

response:
```json
{
  "id": "[message id]",
  "sender": "[your name]",
  "message": "[message you send]",
  "time": "[local time of message sending]"
}
```

#### Get previous messages of a person
url:
```shell
/api/suppress
```

method:
```shell
GET
```

header:
```json
{
  "jwt-token": "[the jwt token you got from join]"
}
```

request (form value):
```json
{
  "sender": "[sender name of messages]"
}
```

response:
```json
{
  "data": [
    {
      "id": "[message id]",
      "sender": "[your name]",
      "message": "[message you send]",
      "time": "[local time of message sending]"
    }, ...
  ]
}
```

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

It also provides the _application.yaml_, so you can set for
cluster **ArgoCD**.