package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"golang.org/x/crypto/bcrypt"
	//"github.com/googollee/go-socket.io"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type Player struct {
	Ship Ship `json: "ship"`
}

type Ship struct {
	Name    string `json: "name"`
	Defense int    `json: "defense"`
	Slots   int    `json: "slots"`
}

func main() {

	/*
		server, err := socketio.NewServer(nil)
		if err != nil {
			panic(err)
		}
	*/

	//server.On("connection", ProcessHistory)

	//http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/login", AccountLogin)
	http.HandleFunc("/register", AccountRegister)

	http.ListenAndServe(":3000", nil)
	fmt.Println("Servidor iniciado na porta :3000")
}

func AccountLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session := GetSession()
		defer session.Close()

		accountData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		account := Account{}
		json.Unmarshal(accountData, &account)

		collection := session.DB("galacticbay").C("Players")

		result := Account{}
		collection.Find(account).One(&result)

		fmt.Println(result)
		/*
			if (result != Account{}) {
				server.On("connection", ProcessHistory)
			}
		*/

	}
}

func AccountRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session := GetSession()
		defer session.Close()

		accountData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		account := Account{}
		json.Unmarshal(accountData, &account)

		//player.Password, _ = bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)

		fmt.Println(account)

		collection := session.DB("galacticbay").C("Players")

		collection.Insert(account)

		//w.Header().Set("Content-Type", "application/text")
	}
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}
