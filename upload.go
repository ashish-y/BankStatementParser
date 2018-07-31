package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ledongthuc/pdf"
)

type ParsePageVariables struct {
	Title     string
	PDFText   string
	PageTitle string
}

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
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("C:/Users/ashis/Documents/workspace/statementCheck/uploaded_files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		//Parse the PDF here
		fmt.Println("Parsing the PDF:", handler.Filename)
		content, err := readPdf("C:/Users/ashis/Documents/workspace/statementCheck/uploaded_files/" + handler.Filename) // Read local pdf file
		if err != nil {
			panic(err)
		}
		fmt.Println(content)
		t, err := template.ParseFiles("html/parsepage.html") //parse the html file homepage.html
		if err != nil {                                      // if there is an error
			log.Print("template parsing error: ", err) // log it
		}
		ParsePageVariables := ParsePageVariables{
			Title:     "::::::::::::::PDF TEXT::::::::::",
			PDFText:   content,
			PageTitle: "Page Parser",
		}

		err = t.Execute(w, ParsePageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil {                        // if there is an error
			log.Print("template executing error: ", err) //log it
		}
	}
}
func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
