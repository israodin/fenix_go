package main

import (
	"fmt"
	"net/http"
)

var alive, ready bool

func main() {
	alive = true
	ready = true
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/liveness", livenessProbe)
	http.HandleFunc("/fail-liveness", failLivenessProbe)
	http.HandleFunc("/readines", readinesProbe)
	http.HandleFunc("/fail-readiness", failreadinesProbe)
	// config from Env
	// fmt.Println("FENIX_PORT:", os.Getenv("FENIX_PORT"))

	// config from  file
	// file, err := os.Open("/opt/fenix/conf/fenix.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Print(string(b))

	// file, err := os.Open("/opt/fenix/conf/creds/username.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))

	fmt.Println("serving trafic on  port  6767")
	http.ListenAndServe(":6767", nil)
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!!!")
}

func livenessProbe(w http.ResponseWriter, r *http.Request) {
	if alive {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func readinesProbe(w http.ResponseWriter, r *http.Request) {
	if ready {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func failLivenessProbe(w http.ResponseWriter, r *http.Request) {
	alive = false
}
func failreadinesProbe(w http.ResponseWriter, r *http.Request) {
	ready = false
}
