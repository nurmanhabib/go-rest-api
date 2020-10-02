package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/nurmanhabib/go-rest-api/interfaces/container"
)

type Router struct {
    e *gin.Engine
    container *container.Container
}

func NewRouter(c *container.Container) *Router {
    return &Router{
        e: gin.Default(),
        container: c,
    }
}

func (r *Router) Boot() *Router {
    pingRoute(r)
    authRoute(r)
    userRoute(r)

    return r
}

func (r *Router) Run() error {
    return r.e.Run(":" + r.container.Config.Port)
}
