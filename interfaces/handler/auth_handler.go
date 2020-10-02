package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/nurmanhabib/go-rest-api/domain/model"
    "github.com/nurmanhabib/go-rest-api/interfaces/response"
)

type Authenticate struct {}

func NewAuthenticate() *Authenticate {
    return &Authenticate{}
}

func (auth *Authenticate) Login(c *gin.Context) {
    var user model.User

    _ = c.ShouldBind(&user)

    user.ValidateLogin(c)

    messages := response.NewValidationErrors()

    messages.Add(response.ErrorField{
        Field:    "email",
        Messages: []string{"Required"},
    })

    messages.Add(response.ErrorField{
        Field:    "email",
        Messages: []string{"Must be email format"},
    })

    response.Error(c, messages).Result()

    //response.Success(c, "data", "Authenticated").Result()
}
