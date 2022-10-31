package main

import (
    "net/http"
    "fmt"
    "os"
    "golang.org/x/net/html"
    "strings"
)
func main() {
   for _,arg := range os.Args[1:] {
        title(arg)
   }
}

func title(url string) error {
    resp,err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    ct := resp.Header.Get("Content-Type")
    if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
//         resp.Body.Close()
        return fmt.Errorf("%s has byte %s,not text/html", url, ct)
    }
    doc,err := html.Parse(resp.Body)
    if err != nil {
        return fmt.Errorf("parsing %s as HTML:%v",url, err)
    }
//     resp.Body.Close()
    if err != nil {
        return fmt.Errorf("parsing %s as HTML:%v", url, err)
    }
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
            fmt.Println(n.FirstChild.Data)
        }
    }
    forEachNode(doc, visitNode, nil)
    return nil

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)){
    if pre != nil {
        pre(n)
    }
    for c:=n.FirstChild; c!=nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}