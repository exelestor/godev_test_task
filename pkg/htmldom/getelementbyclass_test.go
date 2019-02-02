package htmldom

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

const page = `<html>
<head>
<title>Sample Title</title>
</head>
<body>
    <div>
        <div class="wrong">Wrong Text</div>
		<div class="ok" style="color: red">Right Text</div>
    </div>
	<p class="ok_too">Also Right Text</p>
</body>
</html>`

func TestGetFirstElementByClass(t *testing.T) {
	dom, err := html.Parse(strings.NewReader(page))
	if err != nil {
		t.Errorf("cannot parse html")
	}

	elem := GetFirstElementByClass(dom, "ok")
	if elem.FirstChild.Data != "Right Text" {
		t.Errorf(`got incorrect element by search id "ok", returned text: %s`, elem.FirstChild.Data)
	}

	elem = GetFirstElementByClass(dom, "ok_too")
	if elem.FirstChild.Data != "Also Right Text" {
		t.Errorf(`got incorrect element by search id "ok_too", returned text: %s`, elem.FirstChild.Data)
	}
}
