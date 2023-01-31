package server

import (
	"context"
	"net/http"
	"todos/config"
)

func Run(conf config.Config, handler http.Handler) error {
	server := new(http.Server)

	server.Addr = conf.Server.Host + ":" + conf.Server.Port
	server.Handler = handler

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func ShutDown(server *http.Server, ctx context.Context) error {
	return server.Shutdown(ctx)
}
