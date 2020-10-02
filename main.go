package main

import (
    "github.com/nurmanhabib/go-rest-api/interfaces/container"
    "github.com/nurmanhabib/go-rest-api/interfaces/routes"
    "log"
)

func main() {
    c := container.NewContainer()

    router := routes.NewRouter(c).Boot()

    log.Fatal(router.Run())
}
