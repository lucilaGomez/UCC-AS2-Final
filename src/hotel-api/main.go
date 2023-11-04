package main

import (
	"hotel-api/app"
	"hotel-api/db"
	"hotel-api/queue"
)

func main() {

	db.Init_db()      // Inicialización de la Base de Datos
	queue.InitQueue() // Inicialización de la Cola de Mensajes
	app.StartRoute()  // Inicialización de las Rutas de la Aplicación
}
