package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	m "github.com/ytobing/scrapper/models"
)

type Src struct {
}

func InitService() Src {

	return Src{}
}

func (s *Src) ReadData(URL string) (res m.Data, err error) {
	res.URL = URL
	URL = checkURL(URL)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(string) {
		defer wg.Done()
		response := getURL(URL)
		res.HTMLVersion, err = checkHTMLVersion(response)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}(URL)

	response := getURL(URL)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func(*goquery.Document) {
		defer wg.Done()
		res.LinkDetails = checkLink(doc, URL)
		return
	}(doc)

	wg.Add(1)
	go func(*goquery.Document) {
		defer wg.Done()
		res.PageTitle = readTitle(doc)
		return
	}(doc)

	wg.Add(1)
	go func(*goquery.Document) {
		defer wg.Done()
		res.IsContainLoginForm = findLogin(doc)
		return
	}(doc)

	wg.Add(1)
	go func(*goquery.Document) {
		defer wg.Done()
		res.HeadingsCount = countHeadings(doc)
		return
	}(doc)

	wg.Wait()

	fmt.Println(res)
	return
}

func checkHTMLVersion(response *http.Response) (string, error) {

	content, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if strings.Contains(string(content), "<!DOCTYPE html>") {
		return "5", nil
	} else {
		return "<5", nil
	}
}

func getURL(URL string) (response *http.Response) {
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}

func readTitle(doc *goquery.Document) (pageTitle string) {
	fmt.Println(doc)
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		pageTitle = s.Find("title").Text()

	})

	return
}

func checkLink(doc *goquery.Document, URL string) (linkCount m.LinkDetails) {

	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		href, ok := s.Attr("href")
		if ok {
			u, err := url.Parse(href)
			if err != nil {
				fmt.Println(err)
				return
			}
			uReq, err := url.Parse(URL)
			if err != nil {
				fmt.Println(err)
				return
			}

			hostReq := uReq.Hostname()
			_ = hostReq

			if u.Hostname() == "" {
				linkCount.InternalLink += 1
				resp, err := http.Get(fmt.Sprintf("%s://%s%s", uReq.Scheme, hostReq, u.Path))
				if err != nil {
					fmt.Println(err)
					return
				}
				// fmt.Println(resp.StatusCode)
				if err != nil {
					linkCount.InaccessibleLink += 1

				} else if resp.StatusCode != http.StatusOK {
					linkCount.InaccessibleLink += 1
				}

			} else {
				// fmt.Println(u.Hostname())
				linkCount.ExternalLink += 1

				resp, err := http.Get(u.String())
				if err != nil {
					linkCount.InaccessibleLink += 1

				} else if resp.StatusCode != http.StatusOK {
					linkCount.InaccessibleLink += 1
				}

			}
		}

	})

	return
}

func findLogin(doc *goquery.Document) (isLoginExist bool) {
	doc.Find("input").Each(func(i int, s *goquery.Selection) {

		t, ok := s.Attr("type")
		if ok && t == "password" {
			isLoginExist = true
		}

	})
	return
}

func countHeadings(doc *goquery.Document) (headingCount map[string]int) {
	headingCount = make(map[string]int)
	for i := 1; i <= 6; i++ {
		tag := fmt.Sprintf("h%d", i)
		key := fmt.Sprintf("heading_%d", i)
		doc.Find("body").Find(tag).Each(func(i int, s *goquery.Selection) {
			headingCount[key] += 1

		})
	}

	return
}

func checkURL(URL string) string {
	if !strings.Contains(URL, "http://") && !strings.Contains(URL, "https://") {
		return "http://" + URL
	}
	return URL

}
