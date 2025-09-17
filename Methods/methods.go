package methods

import (
	"net/http"
)

func GetWithUa(client *http.Client ,url string) (*http.Response,error) {
	
	resp,err := http.NewRequest("GET", url, nil)
	if err !=nil{
		
		return nil,err
	}

	resp.Header.Set("User-Agent","Mozilla/5.0")
	
	return client.Do(resp)
	
}
