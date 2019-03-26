package fasthttputil

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// SendJSON sends a struct as JSON
func SendJSON(ctx *fasthttp.RequestCtx, d interface{}) error {
	ctx.SetContentType("application/json")
	return json.NewEncoder(ctx).Encode(d)
}
