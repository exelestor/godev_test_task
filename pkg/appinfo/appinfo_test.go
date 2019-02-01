package appinfo

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

type testpair struct {
	Html   string
	Name   string
	Icon   string
	Rating float64
}

const page = `<html>
<head>
<title>Sample Page</title>
</head>
<body>
    <div class="AHFaub"><span>Some Name</span></div>
	<div class="dQrBL"><img src="https://example.com/icon.png" class="T75of ujDFqe" aria-hidden="true" alt="Cover art" itemprop="image"></div>
	<div class="BHMmbe" aria-label="Средняя оценка: 4,8 из 5">4,8</div>
</body>
</html>`

var tests = testpair{
	Html:   page,
	Name:   "Some Name",
	Icon:   "https://example.com/icon.png",
	Rating: 4.8,
}

func makeHTMLNode(code string) (*html.Node, error) {
	r := strings.NewReader(code)
	dom, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return dom, nil
}

func TestFindName(t *testing.T) {
	dom, _ := makeHTMLNode(tests.Html)
	name := findName(dom)
	if name != tests.Name {
		t.Errorf(`name was incorrect, got: "%s", want: "%s".`, name, tests.Name)
	}
}

func TestFindIcon(t *testing.T) {
	dom, _ := makeHTMLNode(tests.Html)
	icon := findIcon(dom)
	if icon != tests.Icon {
		t.Errorf(`icon was incorrect, got: "%s", want: "%s".`, icon, tests.Icon)
	}
}

func TestFindRating(t *testing.T) {
	dom, _ := makeHTMLNode(tests.Html)
	rating := findRating(dom)
	if rating != tests.Rating {
		t.Errorf(`rating was incorrect, got: %f, want: %f.`, rating, tests.Rating)
	}
}
