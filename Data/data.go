package Data

import (
	"database/sql"
	"fmt"
)

// Declaramos la variable global con la conexión a la BBDD
var db *sql.DB

func TestCon() {
	obtenerBaseDeDatos()
	fmt.Println("Prueba de acceso")
}
func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "admin"
	pass := "ADMIN123"
	host := "mercadolibredb.cif4vrop2tai.us-east-2.rds.amazonaws.com:3306"
	nombreBaseDeDatos := "DBMercadoLibre"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(db)
	return db, nil
}
