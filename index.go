package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mooc/db"
	"net/http"
	// "os"
	"regexp"
)

func main() {
	webSites := make([]db.WebSite, 0)
	webSites = db.Query("select * from website")
	fmt.Println(webSites)
	for _, value := range webSites {
		fmt.Println(value.Href)
		getUrlFromDomain(value.Href)
	}
	// getUrlFromDomain("http://www.imooc.com")
}
func getUrlFromDomain(url string) int {
	resp, err := http.Get(url)
	checkErr(err)
	fmt.Println(resp.Status, resp.StatusCode)
	myReader := bufio.NewReader(resp.Body)
	// writeFile, err := os.Create("imooc.html")
	// checkErr(err)
	// myWriter := bufio.NewWriter(writeFile)
	i := 1
	for ; ; i++ {
		line, err := myReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		re, _ := regexp.Compile("<style([\\S\\s]*)</style>")
		src := re.ReplaceAllString(line, "")
		re, _ = regexp.Compile("<script([\\S\\s]*)</script>")
		src = re.ReplaceAllString(src, "")
		// len, err := writeFile.WriteString(src)
		checkErr(err)
		// fmt.Printf("%d times input file lens %d\n", i, len)
		// re, _ = regexp.Compile("<a href=.*</a>")
		// str := re.FindAllString(src, -1)
		// for _, value := range str {
		// 	len, err := writeFile.WriteString(value)
		// 	checkErr(err)
		// 	fmt.Printf("%d times input file lens %d\n", i, len)
		// }
		reg := regexp.MustCompile(`<a\b[^>]+\bhref="([^"]*)".*[^>]*>([\s\S]*?)</a>`)
		result := reg.FindAllStringSubmatch(line, -1)
		if result != nil {
			// fmt.Println(result[0])
			if result[0][1][0] != '/' {
				continue
			}
			thisUrl := "http://www.imooc.com" + result[0][1]
			isExit := db.Query("select * from website where href='" + thisUrl + "'")
			// fmt.Println(thisUrl, isExit)
			if len(isExit) > 0 {
				continue
			}
			name := result[0][2]
			affectRow := db.Insert("insert into website(name,href) values(?,?)", name, thisUrl)
			fmt.Printf("insert %d row\n", affectRow)
		}
	}
	return i
	// myWriter.Flush()
	// defer writeFile.Close()
}

// func regexHtml(line string) {

// }

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
