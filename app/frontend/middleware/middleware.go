package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func Reister(h *server.Hertz) {
	h.Use(GlobalAuth())

}