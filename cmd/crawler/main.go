package main
import(
	"fmt"
	"bufio"
	"strings"
	"os"
	"web_crawler/internal/crawler"
)
func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Please enter seed URL:")
	input,_:=reader.ReadString('\n')
	seedURL:=strings.TrimSpace(input)
	c:=crawler.New(seedURL,7)
	c.Start()
	
}