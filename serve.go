package main

import (
    "fmt"
    "io"
    "net/http"
    "io/ioutil"
)

func check(e error){
    if e!=nil{
        panic(e)
    }
}


func pagehandler(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.RequestURI)
    if r.RequestURI == "/" {
    files, _ := ioutil.ReadDir(".")
    io.WriteString(w,"<html><body>")
    for _, f := range files{
        io.WriteString(w,"<a href='"+f.Name()+"'>"+f.Name()+"</a><br>")
    }
    io.WriteString(w,"</body></html>")
    } else {
        files, _ := ioutil.ReadDir(".")
        for _, f := range files{
        if "/"+f.Name() == r.RequestURI{
            if f.IsDir(){
                rfiles, _ := ioutil.ReadDir(f.Name())
                io.WriteString(w,"<html><body>")
                for _, f := range rfiles{
                    io.WriteString(w,"<a href='"+f.Name()+"'>"+f.Name()+"</a><br>")
                }
                io.WriteString(w,"</body></html>")
            } else {
            data,err := ioutil.ReadFile(f.Name())
            check(err)
            io.WriteString(w,string(data))
            }
        }
    }
    }
}

func main(){
    fmt.Println("Running server")
    http.HandleFunc("/",pagehandler)
    http.ListenAndServe(":8000",nil)
}