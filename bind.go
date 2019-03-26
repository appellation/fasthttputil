package fasthttputil

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/appellation/fasthttputil/convert"
	"github.com/valyala/fasthttp"
)

// Bind binds the request context's content to the struct
func Bind(ctx *fasthttp.RequestCtx, d interface{}) error {
	if string(ctx.Request.Header.ContentType()) == "application/json" {
		return BindJSON(ctx, d)
	}

	return BindForm(ctx, d)
}

// BindJSON binds the request context's body to the struct
func BindJSON(ctx *fasthttp.RequestCtx, d interface{}) error {
	return json.Unmarshal(ctx.PostBody(), d)
}

// BindForm binds a request context's query parameters to the struct
func BindForm(ctx *fasthttp.RequestCtx, d interface{}) error {
	to := reflect.ValueOf(d)
	if !to.IsValid() || to.IsNil() {
		return ErrInvalidBinding
	}

	to = to.Elem()
	toType := to.Type()
	for i := 0; i < to.NumField(); i++ {
		field := to.Field(i)
		if !field.CanSet() {
			continue
		}

		toTypeField := toType.Field(i)
		key := toTypeField.Tag.Get("form")
		if key == "" {
			key = strings.ToLower(toTypeField.Name)
		}
		from := ctx.FormValue(key)

		conv, err := convert.FromString(string(from), field.Type())
		if err != nil {
			return ErrInvalidParameter
		}

		field.Set(conv)
	}

	return nil
}
