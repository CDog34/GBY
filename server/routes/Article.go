package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var LinkRoutes = Routes{
	Route{
		Name:       "LinkList",
		Method:     "GET",
		Pattern:    "/link",
		HandleFunc: LinkList,
	},
	Route{
		Name:       "LinkShow",
		Method:     "GET",
		Pattern:    "/link/{linkId}",
		HandleFunc: LinkShow,
	},
}
