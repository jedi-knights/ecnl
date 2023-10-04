# MongoDB

To get started get the follwing packages

```sh
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

```go
package main

import (
    "context"
    "log"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

Here we are adding the mongo and options packages, which the MongoDB Go driver provides.

Running MongoDB locally

```sh
docker run -it --rm -p 27017:27017 --name mongodb mongo
```

Running MongoDB locally with persistent data

```sh
docker run -it --rm --name mongodb -v ~/mongo/data:/data/db -p 27017:27017 mongo:latest
```

This maps the internal /data/db to your local ~/mongo/data directory.


List running containers
  
```sh
docker container ls
```

List all containers

```sh
docker container ls -a
```

Removing a container

```sh
docker container rm <container id>
```
