package model

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gopkg.in/go-playground/validator.v9"
    "reflect"
)

type User struct {
    Name string `json:"name" form:"name"`
    Email string `json:"email" form:"email" validate:"required|email"`
    Password string `json:"password" form:"password" validate:"required"`
}

func (u *User) ValidateLogin(c *gin.Context) {
    var validate *validator.Validate

    validate = validator.New()

    err := validate.Struct(u)

    if err != nil {
        fmt.Println(err.(validator.ValidationErrors))
        fmt.Println()

        for _, err := range err.(validator.ValidationErrors) {

            fmt.Println(err.Namespace())
            fmt.Println(err.Field())
            fmt.Println(err.StructNamespace())
            fmt.Println(err.StructField())
            fmt.Println(err.Tag())
            fmt.Println(err.ActualTag())
            fmt.Println(err.Kind())
            fmt.Println(err.Type())
            fmt.Println(err.Value())
            fmt.Println(err.Param())
            fmt.Println()
            fmt.Println()
        }
    }

    x := reflect.TypeOf(u)
    fmt.Println(x.Elem().NumField())
    for i:=0;i<x.Elem().NumField(); i++ {
        fmt.Println(x.Elem().Field(i).Tag.Get("ok"))
    }
}
