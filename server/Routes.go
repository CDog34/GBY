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
	Route{
		Name:"ArticleUpdate",
		Method:"PUT",
		Pattern:"/article/{articleId}",
		HandleFunc:ArticleUpdate,
	},
	Route{
		Name:"ArticleDelete",
		Method:"DELETE",
		Pattern:"/article/{articleId}",
		HandleFunc:ArticleDelete,
	},
}