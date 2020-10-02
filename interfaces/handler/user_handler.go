package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/nurmanhabib/go-rest-api/domain/model"
    "github.com/nurmanhabib/go-rest-api/interfaces/response"
    "net/http"
)

func GetUsers(c *gin.Context) {
    res := []model.User{
        {
            Name: "Habib",
        },
        {
            Name: "Heru",
        },
    }

    c.Status(http.StatusNotFound)

    response.Success(c, res, "").Result()
}

func StoreUsers(c *gin.Context) {
    var user *model.User

    err := c.ShouldBind(&user)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(user)

    res := user

    response.Success(c, res, "User is created").Result()
}
