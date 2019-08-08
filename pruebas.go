package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func pruebas() {
	// var nombre string
	// var apellido string
	// fmt.Print("ingresar nombre: ")
	// fmt.Scanf("%s", &nombre)
	// fmt.Print("ingresar apellido: ")
	// fmt.Scanf("%s", &apellido)
	// fmt.Printf("hola %s %s, bienvenido\n", nombre, apellido)
	// fmt.Println("hola"[0])
	// fmt.Println(string(64))

	// banquito := muebles.Sillas{
	// 	Name:  "Banquito",
	// 	Price: 200,
	// }
	// banquito.Tela = "Cuero"
	// banquito.SetMaterial("Cuerina")
	// ratona := muebles.Mesas{
	// 	Name:  "Ratona",
	// 	Price: 300,
	// }
	// fmt.Println(banquito)
	// fmt.Println(ratona)
	// // banquito.SetMaterial("Cuerina")
	// muebles.CallGetPrice(banquito)
	// muebles.CallGetPrice(ratona)

	// a := 123
	// fmt.Println(a.(interface))

	// var client = &http.Client{Timeout: 10 * time.Second}
	// url := "http://jazmeendeco.com.ar.ci5.toservers.com/api/?date=07/07/2019&limit=10"
	// resp, err := client.Get(url)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// var factura []Factura
	// err = json.NewDecoder(resp.Body).Decode(&factura)

	// if err != nil {
	// 	panic(err.Error)
	// }

	// fmt.Println(factura[0].Comentario)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, factura[0].Comentario)
	// })
	// http.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "HOME")
	// })
	// http.ListenAndServe(":8000", nil)

	fmt.Println("iniciando")

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/jazmeendeco")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("conectado")

	results, err := db.Query("SELECT id, RazonSocial FROM datos_nuevos ORDER BY id DESC LIMIT 10")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var factura Factura
		err = results.Scan(
			&factura.ID,
			&factura.RazonSocial,
		)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ID:", factura.ID, "RS:", factura.RazonSocial)
	}

}

// Factura factura
type Factura struct {
	ID                  int    `json:"id"`
	NroBoleta           string `json:"NroBoleta"`
	RazonSocial         string `json:"RazonSocial"`
	Telefono            string `json:"Telefono"`
	EMAIL               string `json:"EMAIL"`
	FechaIng            string `json:"FechaIng"`
	FechaEst            string `json:"FechaEst"`
	FechaConf           string `json:"FechaConf"`
	Caracteristicas     string `json:"Caracteristicas"`
	Descripcion         string `json:"Descripcion"`
	Mentrega            string `json:"mentrega"`
	Comentario          string `json:"Comentario"`
	Estado              string `json:"Estado"`
	Localidad           string `json:"localidad"`
	DatosInc            string `json:"datosInc"`
	Expreso             string `json:"expreso"`
	Pago                string `json:"pago"`
	ComboComent         string `json:"comboComent"`
	EstructuraTipo      string `json:"estructura_tipo"`
	EstructuraFechaPed  string `json:"estructura_FechaPed"`
	EstructuraFechaEnt  string `json:"estructura_FechaEnt"`
	CarpinteriaTipo     string `json:"carpinteria_tipo"`
	CarpinteriaFechaPed string `json:"carpinteria_FechaPed"`
	CarpinteriaFechaEnt string `json:"carpinteria_FechaEnt"`
	IDOld               string `json:"idOld"`
}
