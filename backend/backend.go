package backend

import (
	"space/backend/hooks"
	"space/lib/logger"
	"space/lib/server"
	"space/lib/server/ca"
	"space/lib/server/middlewares"

	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ip         = net.IPv4(127, 0, 0, 1)
	port       = 8000
	httpServer *http.Server
)

func init() {
	hooks.Bind()
	middlewares.SetOrigin(ip, port)
}

func Start() {
	log := logger.Log.WithOptions(zap.Fields(
		zap.String("ip", ip.String()),
	))

	log.Info("starting engine")

	// Sets default size only if not set
	gin.SetMode(gin.ReleaseMode)

	// Auto generate self signed certificate and private key
	ca.Setup(logger.Log)
	certFile := ca.GetCertificate()
	keyFile := ca.GetPrivateKey()
	log = log.WithOptions(zap.Fields(
		zap.String("certFile", certFile),
		zap.String("keyFile", keyFile),
	))

	// Start HTTPS server
	httpServer = server.Initialize(ip, port, server.Router)
	go func() {
		errorsStartingUp := 0
		var err error

		for errorsStartingUp < 5 {
			log.Info("attempting to start HTTPS server",
				zap.Int("port", port),
				zap.Int("errorsStartingUp", errorsStartingUp),
			)

			httpServer = server.Initialize(ip, port, server.Router)
			err = httpServer.ListenAndServeTLS(certFile, keyFile)

			if err != nil {
				log.Info("retrying to start HTTPS server, because of an error",
					zap.Int("port", port),
					zap.Int("errorsStartingUp", errorsStartingUp),
					zap.Error(err),
				)

				port++
				errorsStartingUp++
				continue
			}
			break
		}

		log.Panic("failed to start HTTPS server",
			zap.Int("port", port),
			zap.Int("errorsStartingUp", errorsStartingUp),
			zap.Error(err),
		)
	}()

	server.Wait(httpServer, log)
}
