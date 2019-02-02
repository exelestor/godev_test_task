package appinfo

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/exelestor/godev_test_task/pkg/htmldom"
)

type Result struct {
	ID     string   `json:"id"`
	Name   string   `json:"name,omitempty"`
	Icon   string   `json:"icon,omitempty"`
	Rating *float64 `json:"rating,omitempty"`
	Error  string   `json:"error,omitempty"`
}

const (
	APIURL   = "https://play.google.com/store/apps/details"
	IDNAME   = "AHFaub"
	IDICON   = "dQrBL"
	IDRATING = "BHMmbe"

	ERRAPPNOTFOUND = "application not found"
)

var (
	client = http.Client{
		Timeout: time.Second * 5,
	}
)

func Get(id, lang string) (Result, error) {
	dom, err := makeRequest(id, lang)
	if err != nil {
		if err.Error() == ERRAPPNOTFOUND {
			return Result{ID: id, Error: ERRAPPNOTFOUND}, nil
		}
		return Result{}, err
	}

	result := Result{
		ID:   id,
		Name: findName(dom),
		Icon: findIcon(dom),
	}

	rating := findRating(dom)
	if rating != -1 {
		result.Rating = &rating
	}

	return result, nil
}

func SetTimeout(duration time.Duration) {
	client = http.Client{
		Timeout: duration,
	}
}

func makeRequest(id, lang string) (*html.Node, error) {
	var url string
	if lang != "" {
		url = fmt.Sprintf("%s?id=%s&hl=%s", APIURL, id, lang)
	} else {
		url = fmt.Sprintf("%s?id=%s", APIURL, id)
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf(ERRAPPNOTFOUND)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("returned non 200 code: %d", resp.StatusCode)
	}

	dom, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return dom, nil
}

func findName(node *html.Node) string {
	element := htmldom.GetFirstElementByClass(node, IDNAME)
	if element == nil {
		return ""
	}

	return element.FirstChild.FirstChild.Data
}

func findIcon(node *html.Node) string {
	element := htmldom.GetFirstElementByClass(node, IDICON)
	if element == nil {
		return ""
	}

	src, ok := htmldom.GetAttribute(element.FirstChild, "src")
	if !ok {
		return ""
	}

	return src
}

func findRating(node *html.Node) float64 {
	element := htmldom.GetFirstElementByClass(node, IDRATING)
	if element == nil {
		return -1
	}

	rating := strings.Replace(element.FirstChild.Data, ",", ".", -1)
	result, err := strconv.ParseFloat(rating, 10)
	if err != nil {
		return -1
	}

	return result
}
