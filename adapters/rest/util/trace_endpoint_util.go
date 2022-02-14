package util

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/core/tools/providers/tracer"
)

func TraceRestEndpoint(echoCtx echo.Context, identificator string) (tracer.Span, context.Context) {
	jeagerTracer := tracer.New("rest")
	span := jeagerTracer.StartSpanFromRequest(echoCtx.Request(), identificator)
	ctx := jeagerTracer.ContextWithSpan(echoCtx.Request().Context(), span)
	return span, ctx
}
