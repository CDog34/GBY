package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var AdminRoutes = Routes{
	Route{
		Name:       "ArticleListAll",
		Method:     "GET",
		Pattern:    "/admin/article",
		HandleFunc: ArticleListAll,
		Roles:      []int{100},
	},
	Route{
		Name:       "ArticleCreate",
		Method:     "POST",
		Pattern:    "/admin/article",
		HandleFunc: ArticleCreate,
		Roles:      []int{100},
	},
	Route{
		Name:       "ArticleDelete",
		Method:     "DELETE",
		Pattern:    "/admin/article/{articleId}",
		HandleFunc: ArticleDelete,
		Roles:      []int{100},
	},
	Route{
		Name:       "ArticleUpdate",
		Method:     "PUT",
		Pattern:    "/admin/article/{articleId}",
		HandleFunc: ArticleUpdate,
		Roles:      []int{100},
	},
	Route{
		Name:       "ArticleRecover",
		Method:     "GET",
		Pattern:    "/admin/article/{articleId}/recover",
		HandleFunc: ArticleRecover,
		Roles:      []int{100},
	},
}
