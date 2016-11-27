package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var ArticleRoutes = Routes{
	Route{
		Name:       "ArticleIndex",
		Method:     "GET",
		Pattern:    "/article",
		HandleFunc: ArticleIndex,
	},
	Route{
		Name:       "ArticleShow",
		Method:     "GET",
		Pattern:    "/article/{articleId}",
		HandleFunc: ArticleShow,
	},
}
