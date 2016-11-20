package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
	. "github.com/CDog34/GBY/server/services"
)

var IndexRoutes = Routes{
	Route{
		Name:"Index",
		Method:"GET",
		Pattern:"/",
		HandleFunc:Index,
	},
}

var RouteRules = append(
	IndexRoutes,
	ArticleRoutes...
)
