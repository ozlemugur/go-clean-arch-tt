![go-clean-arch-tt]

# go-clean-arch-tt

- Don’t reinvent the wheel.

- Polymorphism gives you the ability to create one module calling another and yet have the compile time dependency point against the flow of control instead of with the flow of control.
you have absolute control over your depedency structure you can avoid wiriting fragile rigid and non reusable modules.  (Robert C. Martin)

- this repo


## Content
- [Quick start](#quick-start)
- [Local Debugging Cheat Sheet](#local-debugging-cheat-sheet)


## Quick Start

```sh
# prepare the environment
$ make prepare

# create swagger files
$ make swag-v1

# To bring everything up together, use the compose-up command.  Postgres, App
$ make compose-up

# Run app with migrations
$ make run

```

 Check if it is alive ? : 
 ```sh
 curl http://localhost:8080/swagger/index.html
```


## Local Debugging Cheat Sheet

## Launch.json (Visual studio code example)

```sh
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug API",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/app",// The main directory of the application
            "cwd": "${workspaceFolder}",
            "env": {
                "GIN_MODE": "debug" // Environment variables
            },
            "args": [] // Add any arguments required for your API to run here
        }
    ]
}
```

Without application container:

```sh
$ make compose-up-without-app
```


## Postgresql

To see what is happening in the database:

```sh
# 
$ docker exec -it postgres psql -U user -d postgres
# to see tables: 
$ \dt
# to see the content of the table :
$ \d messages
# to quit psql: 
$ \q 
```



## message mocks

```sh
{
  "content": " Diego, libre dans sa tête Song by Johnny Hallyday https://open.spotify.com/track/0qJW9XIdyvr4yQrlUFP8xq?si=34cbcd7be5b34ea7",
  "recipient_phone": "05057136986"
}

{
  "content": " Legends never die Song by League of legends, Against The Current https://open.spotify.com/track/0TI3HDmlvuD0rCwHe5m2wD?si=c8ff85eca986442c",
  "recipient_phone": "05057136986"
}

{
  "content": " Canta Per me by Yoki kajiura  https://open.spotify.com/track/0TI3HDmlvuD0rCwHe5m2wD?si=c8ff85eca986442c",
  "recipient_phone": "05057136988"
}

{
  "content": " king Song by Florence and the machine https://open.spotify.com/track/1VSngtLdJhrlfHkLxTyOXK?si=d9292df2504e4da0",
  "recipient_phone": "05057136986"
}

{
  "content": " mother nature Song by The Hu and LP https://open.spotify.com/track/35SoEGEXsaNnfi8PsT8xEC?si=4347af98187349fa",
  "recipient_phone": "05057136986"
}

{
  "content": " winding river by yu-peng chen https://open.spotify.com/track/04WnFdVesT0VLu1Fc57VoI?si=c12404e1bbb34564",
  "recipient_phone": "05057136986"
}

{
  "content": " guizhong's lullaby by jordy chandra, beside bed https://open.spotify.com/track/0n2sLg3mtyxqGZMtGf0Uow?si=0617f0ebd3b148b2",
  "recipient_phone": "05057136986"
}

{
  "content": " nana para mi song by clara peya silvia perez cruz https://open.spotify.com/track/5IY5cuo1nQbcJvzE8h2YvF?si=8049c9c2cc0b4d4c",
  "recipient_phone": "05057136986"
}
```




## Start over from the beginning

```sh
$ make compose-down
$ make docker-rm-volume

```


## docker helper commands

## docker helper commands

Volume check:
```sh
$ docker volume ls  
$ docker volume ls -f dangling=true
$ docker volume rm "volumename"
$ docker volume inspect "volumename"
```

Docker main check commands:
```sh
$ docker ps
$ docker ps -a 
$ docker rm "containerid"
$ docker start/stop "containerid"
$ docker images
$ docker rmi "imageid"
$ docker exec -it postgres psql -U user -d postgres
```

docker log check:
```sh
$ docker logs -tail=all <containerid>
```


git commands:
```sh
$ git remote -v
$ pwd
$ git reset --hard origin
```


## Possible Error Amd Solutions

- "message":"app - Run - httpServer.Notify: listen tcp :8080: bind: address already in use"}

ozlemugur@cassandra automatic-message-sender % lsof -i :8080
COMMAND     PID      USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
app       42073 ozlemugur    9u  IPv6 0x9be0dbb3dbb48eb7      0t0  TCP *:http-alt (LISTEN)
Arc\x20He 82052 ozlemugur   34u  IPv6  0xf726809424a9b2f      0t0  TCP localhost:49923->localhost:http-alt (CLOSED)
ozlemugur@cassandra automatic-message-sender % kill -9 42073
ozlemugur@cassandra automatic-message-sender % 



## About logging

- log notes:
```sh
zerolog allows for logging at the following levels (from highest to lowest):
panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)
```

### we should consider 

- TODO: env setup should be reorganized.
- First relational database, if we will have time, we should consider others.
- we should consider bonus points later. (Redis)
- we should reconsider the naming, there is a really huge gap to put together everything in a one place.
- if time permits, we should consider to add JWT Authentication Middleware.
- we should reconsider database field name's struct adherence to go standards.
- we should consider pagination for the sentMessage endpoint
- if we have time we should consider to add message broker.

## Useful links
- [The Clean Architecture article](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Twelve factors](https://12factor.net/ru/)

## the fin
  (docs/img/purple.png)
