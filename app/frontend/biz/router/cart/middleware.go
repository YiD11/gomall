// Code generated by hertz generator.

package cart

import (
	"github.com/YiD11/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Auth(),
	}
}

func _cartMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcartMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addcartitemMw() []app.HandlerFunc {
	// your code...
	return nil
}
