package response

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Success(c *gin.Context, data interface{}, message string) *Context {
    if len(message) < 1 {
        message = http.StatusText(c.Writer.Status())
    }

    response := NewResponseContext(c)
    response.Response = Response{
        Code:    c.Writer.Status(),
        Data:    data,
        Message: message,
    }

    return response
}
