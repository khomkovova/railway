package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

import b64 "encoding/base64"

var PrivateKey, err2 = rsa.GenerateKey(rand.Reader, 2048)
var PublicKey = &PrivateKey.PublicKey

type CredentialsSignin struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Commands struct {
	Speed int `json:"speed"`
	Direction int `json:"direction"`
}
type CommandsRailway struct {
	Firstswitch int `json:"firstswitch"`
	Secondswitch int `json:"secondswitch"`
}

type CredentialsRegistration struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email string `json:"email"`
}

var commands Commands
var commandsRailway  CommandsRailway


func Signin(w http.ResponseWriter, r *http.Request) {
	var creds CredentialsSignin

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var username string
	err = Db.QueryRow("SELECT username FROM users WHERE username=? AND password=?", creds.Username, creds.Password).Scan(&username)
	if err != nil || username == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	cookie, err3 := encodeCookie(creds.Username)
	if err3 == false{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func Signup(w http.ResponseWriter, r *http.Request)  {
	var creds CredentialsRegistration
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(creds)
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := Db.Query("INSERT INTO users (username,password)" + "VALUES ('" + creds.Username + "','" + creds.Password + "')" )
	fmt.Println(id)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Db.Close()
	//cookie, err3 := encodeCookie(creds.Username)
	//if err3 == false{
	//	fmt.Println("error")
	//	return
	//}
	//http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/signin", http.StatusSeeOther)

}


func Admin(w http.ResponseWriter, r *http.Request)  {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user != "admin"{
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}

	jd := json.NewDecoder(r.Body)

	err := jd.Decode(&commandsRailway)
	jd.DisallowUnknownFields()
	err2 := jd.Decode(&commands)
	fmt.Println(commandsRailway)
	fmt.Println(commands)
	if err != nil && err2 !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("command=", commands)
	w.WriteHeader(http.StatusOK)



}

func Welcome(w http.ResponseWriter, r *http.Request) {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var username string
	err := Db.QueryRow("SELECT username FROM users WHERE username=?", user).Scan(&username)
	if err != nil || username == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte("{'speed':'" + string(commands.Speed) + "', 'direction':'" + string(commands.Direction) +"'}"))



}

func GetTrainCommand(w http.ResponseWriter, r *http.Request){
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token != "train1234567890"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(string(commands.Speed + '0')+string(commands.Direction + '0')))
}

func GetRailwayCommand(w http.ResponseWriter, r *http.Request){
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token != "railway1234567890"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(string(commandsRailway.Firstswitch + '0')+string(commandsRailway.Secondswitch + '0')))


}

func decodeCookie(r *http.Request) string {
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token == ""{
		fmt.Println("error")
		return ""
	}
	sDec, _ := b64.StdEncoding.DecodeString(token)
	label := []byte("")
	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, PrivateKey, sDec, label)
	if err != nil{
		fmt.Println("error")
		return ""
	}
	user := string(plainText)
	fmt.Println("token=",token)
	return user
}

func encodeCookie(user string) (http.Cookie, bool)  {
	message := []byte(user)
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, PublicKey, message, label)
	if err != nil{
		return http.Cookie{}, false
	}
	sEnc := b64.StdEncoding.EncodeToString(ciphertext)
	cookie := http.Cookie{Name:"token", Value:sEnc}
	return cookie, true

}

