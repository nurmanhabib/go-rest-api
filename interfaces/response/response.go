package response

import (
    "github.com/gin-gonic/gin"
)

type Context struct {
    Context     *gin.Context
    Response
}

type Response struct {
    Code        int             `json:"code"`
    Data        interface{}     `json:"data,omitempty"`
    Errors      interface{}     `json:"errors,omitempty"`
    Message     string          `json:"message,omitempty"`
    Meta        interface{}     `json:"meta,omitempty"`
}

func NewResponseContext(c *gin.Context) *Context {
    return &Context{
        Context: c,
    }
}

func (r *Context) Meta(meta interface{}) *Context {
    r.Response.Meta = meta

    return r
}

func (r *Context) Result() {
    r.Context.JSON(r.Code, r.Response)
}
