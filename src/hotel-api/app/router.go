package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	// Se ejecuta automáticamente al importar el paquete app.
	router = gin.Default()
	// Inicializa el router utilizando la configuración predeterminada de gin.
	router.Use(cors.Default())
	// Agrega el middleware CORS al router.
}

func StartRoute() {
	// inicia el enrutamiento de la aplicación.
	mapUrls()
	// mapear las rutas (URLs)
	log.Info("Starting server")
	// Registra un mensaje indicando que el servidor se está iniciando.
	router.Run(":8080")

}
