package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type BoardTemplate struct {
	NewBoard Board
}
type Coordinates struct {
	x int
	y int
}

func readXYFromRequest(r *http.Request) (int, int) {
	x, err := strconv.Atoi(r.PostFormValue("x"))
	if err != nil {
		log.Fatal("failed to conv str 1 ")
	}
	y, err := strconv.Atoi(r.PostFormValue("y"))
	if err != nil {
		log.Fatal("failed to conv str 2 ")
	}
	return x, y
}

func readfile() string {
	file, err := os.ReadFile("./static/index.html")
	if err != nil {
		fmt.Println(err)
	}
	return string(file[:])
}

func main() {
	/* to ez
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	*/

	newBoard := InitBoard()
	var PreviousCoordinates Coordinates

	newBoard.ReadPlayerBoard()

	funcMap := template.FuncMap{
		"isTwo": func(i int) bool {
			if i == 2 || i == 5 {
				return true
			}
			return false
		},
		"isThree": func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/style.css" {

			fmt.Println("css called")
			http.ServeFile(w, r, "./static/style.css")
		} else {
			//remember the prefilght request

			//templ := template.Must(template.ParseFiles("./static/index.html"))
			templ, err := template.New("board").Funcs(funcMap).Parse(readfile())
			if err != nil {
				fmt.Println(err)
			}

			templ.Execute(w, BoardTemplate{NewBoard: newBoard})

			fmt.Println("index called")

		}
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {

		templ, err := template.New("board").Funcs(funcMap).Parse(readfile())
		if err != nil {
			fmt.Println(err)
		}

		x, y := readXYFromRequest(r)
		newBoard.ClearBoard()
		newBoard.SelectBoard(x, y, &PreviousCoordinates)
		templ.Execute(w, BoardTemplate{NewBoard: newBoard})
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {

		templ, err := template.New("board").Funcs(funcMap).Parse(readfile())
		if err != nil {
			fmt.Println(err)
		}

		x, y := readXYFromRequest(r)
		newBoard.MovePiece(x, y, &PreviousCoordinates)
		newBoard.CheckForKing()
		newBoard.EnemyMove()
		EnemyCount, PlayerCount := newBoard.CountPieces()

		if PlayerCount== 0 {
			http.ServeFile(w, r, "./static/enemyWin.html")
		}


		if EnemyCount == 0 {
			http.ServeFile(w, r, "./static/userWin.html")
		}

		templ.Execute(w, BoardTemplate{NewBoard: newBoard})

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/turn", h2)
	http.HandleFunc("/move", h3)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

/*

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		   log.Print("HTMX requested");
		   //this header signifies that the request is coming from htmx
		   log.Print(r.Header.Get("HX-Request"));
    time.Sleep(time.Second)
		title := r.PostFormValue("title")
		direcotr := r.PostFormValue("dir")
		templ := template.Must(template.(helpers).ParseFiles("./static/index.html"))
    templ.ExecuteTemplate(w , "film-element" , Film{Title: title, Director: direcotr})
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film", handler2)


*/
