package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var IndexRoutes = Routes{
	Route{
		Name:       "Index",
		Method:     "GET",
		Pattern:    "/",
		HandleFunc: Index,
	},
}

func mergeRouteRules(args ...Routes) Routes {
	var result Routes
	for _, val := range args {
		result = append(result, val...)
	}
	return result
}

var RouteRules = mergeRouteRules(
	IndexRoutes,
	ArticleRoutes,
	AuthRoutes,
	AdminRoutes,
	LinkRoutes,
)
