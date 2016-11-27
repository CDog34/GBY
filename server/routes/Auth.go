package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var AuthRoutes = Routes{
	Route{
		Name:       "AuthLogin",
		Method:     "POST",
		Pattern:    "/login",
		HandleFunc: AuthLogin,
	},
	Route{
		Name:       "AuthNewUser",
		Method:     "POST",
		Pattern:    "/register",
		HandleFunc: AuthNewUser,
	},
}
