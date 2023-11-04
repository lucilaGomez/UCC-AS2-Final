package app

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// se ejecuta automaticamente cuando se importa el paquete 'app'.
	log.SetOutput(os.Stdout)
	// Configura la salida del registro para que se envíe a os.Stdout.
	log.SetLevel(log.DebugLevel)
	// Establece el nivel de registro en DebugLevel.
	log.Info("Starting logger system")
	// Registra un mensaje informativo indicando que el sistema de registro se está iniciando.
}
