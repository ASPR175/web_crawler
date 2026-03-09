package crawler
import(
	"fmt"
	"sync"
)
type Crawler struct {
    SeedURL string,
    Workers int,
    Queue   *URLQueue,
    Visited map[string]struct{},
	mu sync.Mutex,
}
func New(s string,w int)*Crawler{
	return &Crawler{
		SeedURl:s,
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
	for i:=1;i<=c.Workers;i++{
		go Crawling(i,c.Queue.URLs,c.&mu,&wg)
	}
	c.mu.Lock()
}

