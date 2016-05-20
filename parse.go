 package getJson
 
 
import ( 
    "net/http"
    "io/ioutil"
    "fmt"
    
    "github.com/onefiftynine/pargego"
         )


func main(){
    resp, err := http.Get("http://www.omdbapi.com/?t=sds&y=&plot=short&r=json")
       if err != nil{
           
        } 
        
        body, err := ioutil.ReadAll(resp.Body)
        _ = body
        defer resp.Body.Close()
         if err != nil {
            fmt.Printf("%s", err)
         }
        fmt.Printf("%s\n", string(body))
 }