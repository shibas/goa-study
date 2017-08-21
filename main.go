//go:generate goagen bootstrap -d goa-study/design

package main

import (
	"compress/gzip"
	"flag"
	"goa-study/app"
	"goa-study/controllers"
	"goa-study/env"

	"github.com/deadcheat/goacors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	gm "github.com/goadesign/goa/middleware/gzip"
)

func main() {
	// Start your Engine!!
	env.IgniteDBConnection()

	// Create service
	service := goa.New("goa-study-api")
	// cors
	service.Use(goacors.WithConfig(service, &env.CorsConf))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(gm.Middleware(gzip.BestSpeed))

	// Mount beer controller
	b := controllers.NewBeerController(service)
	app.MountBeerController(service, b)

	// Mount tap controller
	t := controllers.NewTapController(service)
	app.MountTapController(service, t)

	// Only On DevelopmentMode
	if env.OnDevelopment {
		swg := controllers.NewSwaggerController(service)
		controllers.MountSwaggerController(service, swg)
	}

	port := flag.Int("p", env.Server.PortNum, "port number. default set on config")
	docPort := flag.Int("dp", env.Server.DocPort, "port number for doc. default set on config")
	host := flag.String("h", env.Server.HostName, "name of server host. default set on config")
	docHost := flag.String("dh", env.Server.DocHostName, "name of server host. default set on config")
	flag.Parse()

	env.Server.HostName = *host
	env.Server.PortNum = *port

	env.Server.DocHostName = *docHost
	env.Server.DocPort = *docPort

	// Start service
	if err := service.ListenAndServe(env.Server.APIHostString()); err != nil {
		service.LogError("startup", "err", err)
	}
}
