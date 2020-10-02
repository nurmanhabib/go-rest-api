package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/nurmanhabib/go-rest-api/interfaces/response"
    "time"
)

func pingRoute(r *Router) {
    r.e.HEAD("/ping", pong)
    r.e.GET("/ping", pong)
}

func pong(c *gin.Context) {
    data := map[string]interface{}{
        "timestamp": time.Now(),
    }

    response.Success(c, data, "PONG").Result()
}
