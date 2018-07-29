package main

import (
	"html/template"
	"net/http"
	"time"
	"fmt"
	"os"
	"io"
	"crypto/md5"
	"strconv"
	"log"
  )

  
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		t, err := template.ParseFiles("html/parsepage.html") //parse the html file homepage.html
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it
		  }
		  err = t.Execute(w, "test") //execute the template and pass it the HomePageVars struct to fill in the gaps
		  if err != nil { // if there is an error
			  log.Print("template executing error: ", err) //log it
			}
		f, err := os.OpenFile("C:/Users/ashis/Documents/workspace/statementCheck/uploaded_files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}