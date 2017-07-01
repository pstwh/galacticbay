package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Login    string `json: "login"`
	Password string `json: "password"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/login", UserLogin)
	http.HandleFunc("/register", UserRegister)

	http.ListenAndServe(":3000", nil)
	fmt.Println("Servidor iniciado na porta :3000")
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session := GetSession()
		defer session.Close()

		playerData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		player := Player{}
		json.Unmarshal(playerData, &player)

		fmt.Println(player)

		collection := session.DB("galacticbay").C("Players")

		collection.UpsertId(player.Login, player)

		//w.Header().Set("Content-Type", "application/text")
	}
}

func UserRegister(w http.ResponseWriter, r *http.Request) {

}

func GetSession() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}
