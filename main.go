package main

import (
	"Goverment/Methods"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	"log"
)



func main(){


	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	

	resp,err := methods.GetWithUa(client, "https://en.wikipedia.org/wiki/List_of_presidents_of_the_United_States")
	
	if err != nil{
		
		fmt.Printf("Error %s \n", err )
		return
	}
	defer func(r io.ReadCloser){
		
		_,_ = io.Copy(io.Discard, r)
		_ = r.Close()
	}(resp.Body)



	
	
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	
	if err != nil{
		log.Fatal("Failed to parse body ",err)
	}
	
	methods.HtmlParsing(doc)
	
	
	


}
