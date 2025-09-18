package methods

import(
	
	"unicode"
	"fmt"
	"github.com/PuerkitoBio/goquery"	
)


type President struct{
	
	Number	string
	Name	string
	Party	string

}

func HtmlParsing(g *goquery.Document) {
	
	var presidents []President
	
	g.Find("tr").Each(func(row int, tr *goquery.Selection){
		
		president := President{}
		
		president.Number  = tr.Find("th").Text()
		tr.Find("td").Each(func(col int, td *goquery.Selection){
			
			switch col {
				
			
			case 1:
				president.Name = KeepLttersAndWhiteSpace(td.Text())

			case 4:
				president.Party = td.Text()

			}


		})
	presidents = append(presidents, president)
	})
	
	// remove the first element
	presidents = presidents[1:]	
	

	//remove the last element 
	presidents = presidents[:len(presidents) -21]

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
