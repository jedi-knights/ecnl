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
