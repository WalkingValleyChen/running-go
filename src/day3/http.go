package day3

import (
	"net/http"
	"fmt"
	"log"
)

func HelloWorld(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	for k,v:=range r.Form{
		fmt.Println("param",k,"->",v)
	}

	fmt.Fprint(w,"HelloWorld Web!")
}

func WebHandler()  {
	http.HandleFunc("/",HelloWorld)
	err:=http.ListenAndServe("localhost:9090",nil)
	if err!=nil{
		log.Fatal("error:",err)
	}
}