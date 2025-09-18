package methods

import(
	"strings"
	"unicode"
	"fmt"
	"github.com/PuerkitoBio/goquery"	
	"net/http"
)

type President struct{
	
	Number	string
	Name	string
	Party	string

}
func GetWithUa(client *http.Client ,url string) (*http.Response,error) {
	
	resp,err := http.NewRequest("GET", url, nil)
	if err !=nil{
		
		return nil,err
	}

	resp.Header.Set("User-Agent","Mozilla/5.0")
	
	return client.Do(resp)
	
}

func HtmlParsing(g *goquery.Document) {
	
	presidents := []President{}

	g.Find("tr").Each(func(row int, tr *goquery.Selection){
		
		
		president := President{}
		president.Number =tr.Find("th").Text()
		tr.Find("td").Each(func(col int, td *goquery.Selection){
			
			switch col {
				
			
			case 1:
				president.Name =Normalize(td.Text()) 

			case 4:
				president.Party = td.Text()

			}


		})
	presidents = append(presidents, president)
	})
	
	// remove the first element
	presidents = presidents[1:]	
	

	//remove the last element 
	// Need to remove this magic number
	
	
	excess_bad_data := 21

	presidents = presidents[:len(presidents) -excess_bad_data]	
		
	
	
	
	

	fmt.Printf("%+v",presidents)

}


func KeepLttersAndWhiteSpace (pattern string) string{
	
	result := make([]rune, 0, len(pattern))
	
	for _,r := range pattern{
		
		if unicode.IsSpace(r) || unicode.IsLetter(r){
			result = append(result, r)
		}

	}
	return string(result)
}

func Normalize(pattern string)string{
	
	pattern = KeepLttersAndWhiteSpace(pattern)
	    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1 // drop the rune
        }
        return unicode.ToLower(r)
    }, pattern)
	

} 



