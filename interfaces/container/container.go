package container

import "github.com/nurmanhabib/go-rest-api/infrastructure/config"

type Container struct {
    Config *config.Config
}

func NewContainer() *Container {
    return &Container{
        Config: config.New(),
    }
}
