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
		Roles:      []int{100},
	},
	Route{
		Name:       "ArticleShow",
		Method:     "GET",
		Pattern:    "/article/{articleId}",
		HandleFunc: ArticleShow,
	},
	Route{
		Name:       "ArticleCreate",
		Method:     "POST",
		Pattern:    "/article",
		HandleFunc: ArticleCreate,
	},
	Route{
		Name:       "ArticleDelete",
		Method:     "DELETE",
		Pattern:    "/article/{articleId}",
		HandleFunc: ArticleDelete,
	},
	Route{
		Name:       "ArticleUpdate",
		Method:     "PUT",
		Pattern:    "/article/{articleId}",
		HandleFunc: ArticleUpdate,
	},
}
