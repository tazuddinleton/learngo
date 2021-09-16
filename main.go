package main

import (
	"fmt"
	"io/ioutil"
	"learngo/tree"
	"net/http"
	"sync"
	"time"

	"mvdan.cc/xurls/v2"
)

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	s := Same(t1, t2)
	fmt.Println(s)
}

func webCrawlerWithMutex() {
	Crawl("https://golang.org/", 4, &UrlFetcher{
		mu:      sync.Mutex{},
		visited: map[string]*FetchResult{},
	})
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	_, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s \n", url)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type UrlFetcher struct {
	mu      sync.Mutex
	visited map[string]*FetchResult
}
type FetchResult struct {
	body string
	urls []string
}

func (f *UrlFetcher) Fetch(url string) (body string, urls []string, err error) {
	f.mu.Lock()
	_, ok := f.visited[url]
	f.mu.Unlock()
	if ok {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(b)
	resp.Body.Close()

	xurls := xurls.Strict()
	urls = xurls.FindAllString(body, -1)

	f.mu.Lock()
	f.visited[url] = &FetchResult{body: body, urls: urls}
	f.mu.Unlock()
	return
}

func binTreeExample() {
	t := tree.New(1)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Walk(t, ch)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	s1 := Collect(t1)
	s2 := Collect(t2)
	return s1 == s2
}

func Collect(t *tree.Tree) string {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Walk(t, ch)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	s := ""
	for v := range ch {
		s += fmt.Sprint(v)
	}
	return s
}

func runSelectExample() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd, even := make(chan int), make(chan int)

	go numberPicker(nums, odd, even)

	for {
		select {
		case v, ok := <-odd:
			if ok {
				fmt.Println("Odd number", v)
			} else {
				return
			}
		case v, ok := <-even:
			if ok {
				fmt.Println("Even number", v)
			} else {
				return
			}
		default:
			fmt.Println("Not ready yet!")
		}
	}
}

func numberPicker(nums []int, odd chan int, even chan int) {
	defer close(odd)
	defer close(even)

	for i := 1; i <= len(nums); i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
}

func runSumExample() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func shootBoomerang(speed int, name string) {
	fmt.Println(name, "Swaaaaa!")
	time.Sleep(time.Millisecond * time.Duration(speed))
	fmt.Println(name, "Swooop!")
}

func IsClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}
