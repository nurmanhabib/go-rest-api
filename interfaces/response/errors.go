package response

import (
    "github.com/gin-gonic/gin"
)

type Errors struct {
    Code int
    Message string
    Fields []ErrorField
}

type ErrorField struct {
    Field string `json:"field"`
    Messages []string `json:"messages"`
}

func NewErrors(message string, code int) *Errors {
    return &Errors{
        code,
        message,
        []ErrorField{},
    }
}

func NewValidationErrors() *Errors {
    return NewErrors("Validation(s) Error", 422)
}

func Error(c *gin.Context, errors *Errors) *Context {
    response := NewResponseContext(c)
    response.Response = Response{
        Code: errors.Code,
        Errors: errors.Fields,
        Message: errors.Message,
    }

    return response
}

func (e *Errors) Add(f ErrorField) *Errors {
    var pos int
    var ok bool

    pos, ok = e.Find(f)

    if ok {
        for _, m := range f.Messages {
            e.Fields[pos].Messages = append(e.Fields[pos].Messages, m)
        }
    } else {
        e.Fields = append(e.Fields, f)
    }

    return e
}

func (e *Errors) Find(f ErrorField) (int, bool) {
    for i, field := range e.Fields {
        if field.Field == f.Field {
           return i, true
        }
    }

    return -1, false
}
