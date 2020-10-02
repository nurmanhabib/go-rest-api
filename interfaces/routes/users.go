package routes

import "github.com/nurmanhabib/go-rest-api/interfaces/handler"

func userRoute(r *Router) {
    v1 := r.e.Group("api/v1")

    v1.GET("users", handler.GetUsers)
    v1.POST("users", handler.StoreUsers)
}
