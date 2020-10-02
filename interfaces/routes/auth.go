package routes

import "github.com/nurmanhabib/go-rest-api/interfaces/handler"

func authRoute(r *Router) {
    v1 := r.e.Group("api/v1")

    authenticate := handler.NewAuthenticate()

    v1.POST("login", authenticate.Login)
}
