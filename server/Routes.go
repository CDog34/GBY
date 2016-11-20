package main

import (
	. "github.com/CDog34/GBY/server/handlers"
	. "github.com/CDog34/GBY/server/services"
)

var routes = Routes{
	Route{
		Name:"Index",
		Method:"GET",
		Pattern:"/",
		HandleFunc:Index,
	},
	Route{
		Name:"ArticleIndex",
		Method:"GET",
		Pattern:"/article",
		HandleFunc:ArticleIndex,
	},
	Route{
		Name:"ArticleShow",
		Method:"GET",
		Pattern:"/article/{articleId}",
		HandleFunc:ArticleShow,
	},
	Route{
		Name:"ArticleCreate",
		Method:"POST",
		Pattern:"/article",
		HandleFunc:ArticleCreate,
	},
}