// +build js
package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"honnef.co/go/js/xhr"
)

type PageView struct {
	vecty.Core
	Tweets []*TweetView
}

type TweetView struct {
	vecty.Core
	Text       string
	Name       string
	Screenname string
	Avatar     string
}

func main() {
	pv := &PageView{}
	vecty.RenderBody(pv)
	pv.refresh()
}

func (pv *PageView) refresh() {
	req := xhr.NewRequest("GET", "/tweets")
	req.ResponseType = "json"
	err := req.Send(nil)
	if err != nil {
		println(err)
		return
	}

	if req.Status == 200 {
		println(req.Response)
		pv.Tweets = []*TweetView{}
		for _, tweet := range req.Response.Get("statuses").Interface().([]interface{}) {
			t := tweet.(map[string]interface{})
			u := t["user"].(map[string]interface{})

			pv.Tweets = append(pv.Tweets, &TweetView{
				Text:       t["text"].(string),
				Name:       u["name"].(string),              //tweet.Get("user").Get("name").String(),
				Screenname: u["screen_name"].(string),       //tweet.Get("user").Get("screen_name").String(),
				Avatar:     u["profile_image_url"].(string), //tweet.Get("user").Get("profile_image_url").String(),
			})
		}
		vecty.Rerender(pv)
	} else {
		println(req.Status)
		println(req.Response)
	}
}

func (pv *PageView) Render() vecty.ComponentOrHTML {
	var tweets vecty.List

	for _, tv := range pv.Tweets {
		tweets = append(tweets, tv)
	}

	return elem.Body(
		elem.Heading1(vecty.Text("#nebraskajs")),
		elem.UnorderedList(tweets),
	)
}

func (tv *TweetView) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Image(vecty.Markup(prop.Src(tv.Avatar))),
		elem.Div(
			vecty.Markup(vecty.Class("user")),
			elem.Span(
				vecty.Markup(vecty.Class("username")),
				vecty.Text(tv.Name),
			),
			elem.Span(
				vecty.Markup(vecty.Class("screenname")),
				vecty.Text("@"+tv.Screenname),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("tweet")),
			vecty.Text(tv.Text),
		),
	)
}
