package utils

import "context"

// Define the context key type.
type pageContext string

// Create a context key for the theme.
var pageContextKey pageContext = "pageContext"

type PageContextValue struct {
	AreaName string
	PageName string
}

func WithPageContext(ctx context.Context, value PageContextValue) context.Context {
	return context.WithValue(ctx, pageContextKey, value)
}

func GetPageContextValue(ctx context.Context) PageContextValue {
	val := ctx.Value(pageContextKey).(PageContextValue)
	return val
}
