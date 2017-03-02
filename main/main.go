package main

import (

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	_"github.com/lib/pq"
	"net/http"

//	"database/sql"
)


var people  map[int]*person

type person struct{
	Id int
	Name    string
	Vorname string
	Tel     int
	Email   string
	ORT     int
	Region  string
}

func indexHandler(rdr render.Render)  {
	rdr.HTML(200,"index", people)
}

func newHandler(rdr render.Render, r *http.Request)  {

	var info [7]string

	nameM := r.FormValue("name")
	vornameM := r.FormValue("vorname")
	telM := r.FormValue("tel")
	emailM := r.FormValue("email")
	ortM := r.FormValue("ort")
	regionM := r.FormValue("region")

	info[1] = nameM
	info[2] = vornameM
	info[3] = telM
	info[4] = emailM
	info[5] = ortM
	info[6] = regionM

	rdr.HTML(200,"new", info)

}

func endHandler(rdr render.Render, r *http.Request)  {/*

	nameM := r.FormValue("name")
	vornameM := r.FormValue("vorname")
	telM := r.FormValue("tel")
	emailM := r.FormValue("email")
	ortM := r.FormValue("ort")
	regionM := r.FormValue("region")



	db,err := sql.Open("postgres", "user=oosy dbname=gorm1 password=oo sslmode=disable")

	res, err := db.Exec("INSERT INTO person VALUES (default, name, vorname, tel, email, ort, region)",
		nameM, vornameM, telM, emailM, ortM, regionM)

	if err != nil {
		return 0
	}
	 return res.RowsAffected()

	rdr.HTML(200, "ende", people)*/
}


func main()  {



	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "html",
		Extensions:[]string{".tmpl",".html"},
		Charset: "UTF-8",
		IndentJSON:true,
	}))

	m.Get("/",indexHandler)
	m.Get("/new", newHandler)
	m.Get("/ende", endHandler)


	m.Run()
}


