package crawler
import(
    "fmt"
    "web_crawler/internal/fetcher"
    "web_crawler/internal/indexer"
    "web_crawler/internal/parser"
)
func Crawling(id int, queue chan string, mu *sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()
    for url := range queue{
        fmt.Println("Worker ",id,"Processing ",url)
        body,err:=Fetcher(url)
        if err!=nil{
            continue
        }
        links,err:=ParseLinks(body)
        if err!=nil{
            fmt.Sprintln("Failed to the Parse the link")
        }

        Index(url,links)
        for _,link:=range links{
            mu.Lock()
             
            mu.Unlock()
        }
    }
}