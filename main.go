package main

import "net/http"

func init() {
	ConnecToDB()
}

func Test(w http.ResponseWriter, r *http.Request) {
	Insert()
}

func main() {

	http.HandleFunc("/", Test)

	http.ListenAndServe(":8080", nil)

}
