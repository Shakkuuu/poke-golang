package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
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
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", viewhandler)
	http.ListenAndServe(":8080", nil)
}

func viewhandler(w http.ResponseWriter, r *http.Request) {
	var page *Page
	page = &Page{Pika: pika, Rotto: rotto}
	switch r.Method {
	case "GET":
		// fmt.Printf("GET: %d", randomu())
		tmpl, err := template.ParseFiles("index.html") // ParseFilesを使う
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, page)
		if err != nil {
			panic(err)
		}
	case "POST":
		// fmt.Printf("POST: %d", randomu())
		jyu := r.FormValue("jyu")
		aian := r.FormValue("aian")
		switch {
		case jyu != "":
			page.Rotto.HP -= page.Pika.Waza[0].ATK * page.Pika.ATK
			fmt.Println("ピカチュウの10まんボルト!")
		case aian != "":
			page.Rotto.HP -= page.Pika.Waza[1].ATK * page.Pika.ATK
			fmt.Println("ピカチュウのアイアンテール!")
		}
		switch randomu() {
		case 0:
			page.Pika.HP -= page.Rotto.Waza[0].ATK * page.Rotto.ATK
			fmt.Println("オーロットのウッドホーン!")
		case 1:
			page.Pika.HP -= page.Rotto.Waza[1].ATK * page.Rotto.ATK
			fmt.Println("オーロットののろい!")
		case 2:
			fmt.Println("オーロットは攻撃してこなかった")
		}
		http.Redirect(w, r, "/", 301)
	}
}

func randomu() int {
	ran := rand.Intn(3)
	// fmt.Printf("kan: %d", ran)
	return ran
}
