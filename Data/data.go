package Data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Funcion para insertar registro en la tabla TB_VerificacionADN
func AddADNRegistration(isMutant_dna bool, dna string) {
	db := obtenerBaseDeDatos()
	registrationDate_dna := time.Now()
	stmt, e := db.Prepare("INSERT INTO TB_VerificacionADN(isMutant_dna, input_dna, registrationDate_dna) VALUES (?,?,?)")
	ErrorCheck(e)
	defer stmt.Close()
	_, err := stmt.Exec(isMutant_dna, dna, registrationDate_dna)
	ErrorCheck(err)
}

type Record struct {
	count_mutant_dna int
	count_human_dna  int
	ratio            float64
}

//Funcion para insertar registro en la tabla TB_VerificacionADN
func GetRegistrationCount() (mutant_dna int, human_dna int, ratio float64) {
	db := obtenerBaseDeDatos()
	res, err := db.Query("select sum(case when isMutant_dna = '1' then 1 else 0 end) as mutant_dna, sum(case when isMutant_dna = '0' then 1 else 0 end) as human_dna  from TB_VerificacionADN")
	ErrorCheck(err)
	defer res.Close()
	var obj Record
	if res.Next() {
		err := res.Scan(&obj.count_mutant_dna, &obj.count_human_dna)
		if obj.count_human_dna != 0 {
			obj.ratio = float64(obj.count_mutant_dna) / float64(obj.count_human_dna)
		} else {
			obj.ratio = float64(obj.count_mutant_dna)
		}
		ErrorCheck(err)
	} else {
		fmt.Println("fallo consulta")
	}

	return obj.count_mutant_dna, obj.count_human_dna, obj.ratio
}

//funcion para capturar error y mostrar panic error
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//funcion para realizar conexion con mysql db DBMercadoLibre
func obtenerBaseDeDatos() (db *sql.DB) {
	Driver := "mysql"
	usuario := "admin"
	pass := "ADMIN123"
	host := "tcp(mercadolibredb.cif4vrop2tai.us-east-2.rds.amazonaws.com:3306)"
	nombreBaseDeDatos := "DBMercadoLibre"
	con, err := sql.Open(Driver, usuario+":"+pass+"@"+host+"/"+nombreBaseDeDatos)
	ErrorCheck(err)
	return con
}
