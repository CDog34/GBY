package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
	. "github.com/CDog34/GBY/server/services"
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
