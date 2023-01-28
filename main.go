package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Pika  *Poke
	Rotto *Poke
}

type WAZA struct {
	Name string
	ATK  int
}

var jyuman = WAZA{Name: "10まんボルト", ATK: 5}
var aianteru = WAZA{Name: "アイアンテール", ATK: 2}

var uddohoon = WAZA{Name: "ウッドホーン", ATK: 3}
var noroi = WAZA{Name: "のろい", ATK: 2}

type Poke struct {
	Name string
	HP   int
	ATK  int
	Waza []WAZA
}

var pika = &Poke{Name: "ピカチュー", HP: 10, ATK: 2, Waza: []WAZA{jyuman, aianteru}}
var rotto = &Poke{Name: "オーロット", HP: 30, ATK: 1, Waza: []WAZA{uddohoon, noroi}}

func main() {
	http.HandleFunc("/", viewhandler)
	http.ListenAndServe(":8080", nil)
}

func viewhandler(w http.ResponseWriter, r *http.Request) {
	var page *Page
	page = &Page{Pika: pika, Rotto: rotto}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("index.html") // ParseFilesを使う
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, page)
		if err != nil {
			panic(err)
		}
	case "POST":
		jyu := r.FormValue("jyu")
		aian := r.FormValue("aian")
		switch {
		case jyu != "":
			page.Rotto.HP -= page.Pika.Waza[0].ATK
		case aian != "":
			page.Rotto.HP -= page.Pika.Waza[1].ATK
		}
		http.Redirect(w, r, "/", 301)
	}
}
