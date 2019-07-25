package server

import (
	"github.com/CyrivlClth/snowserver/http/router"
)

func Run(addr string) error {
	r := router.New()
	return r.Run(addr)
}
