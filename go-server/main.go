package main

import (
	"fmt"
	"log"
	"net/http"

)
func formhandler(w http.ResponseWriter , r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "PArseform() err: %v" ,err)
		return
	}
	fmt.Fprintf(w, "Post reuqest successful")
	Name := r.FormValue("name")
	Adress := r.FormValue("adress")
	fmt.Fprintf(w,"Name = %s\n",Name)
	
	fmt.Fprintf(w,"Adress = %s\n",Adress)

}


func hellohandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w, "404 not found" , http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"menthod is not found" , http.StatusNotFound)
		return
	}
	fmt.Fprint(w,"hello!")

}

func main(){ 
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",hellohandler)

	fmt.Printf("stating server at port 8000\n")
	if err := http.ListenAndServe(":8080",nil);err !=nil{
		log.Fatal(err)
	}

}

