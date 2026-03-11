package crawler
import(
	"fmt"
	"sync"
)
type Crawler struct {
    SeedURL string
    Workers int
    Queue   *URLQueue
    Visited map[string]struct{}
	mu sync.Mutex
}
func New(s string,w int)*Crawler{
	return &Crawler{
		SeedURL:s,
		Workers:w,
	}
}

func (c *Crawler)Start(){
	c.Queue = &URLQueue{
    URLs: make(chan string,100),
    }
	c.Visited = make(map[string]struct{})
	var wg sync.WaitGroup

	c.Visited[c.SeedURL] = struct{}{}
	for i:=0;i<c.Workers;i++{
		wg.Add(1)
		go Crawling(i,c.Queue.URLs,&c.mu,&wg)
	}
	
		c.Queue.URLs<-c.SeedURL
	
   wg.Wait()
	
}

