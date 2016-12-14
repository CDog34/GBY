package Route

import (
	. "github.com/CDog34/GBY/server/handlers"
)

var AuthRoutes = Routes{
	Route{
		Name:       "AuthLogin",
		Method:     "POST",
		Pattern:    "/auth/login",
		HandleFunc: AuthLogin,
	},
	Route{
		Name:       "AuthNewUser",
		Method:     "POST",
		Pattern:    "/auth/register",
		HandleFunc: AuthNewUser,
	},
	Route{
		Name:       "AuthLogout",
		Method:     "GET",
		Pattern:    "/auth/logout",
		HandleFunc: AuthLogout,
	},
	Route{
		Name:       "AuthValid",
		Method:     "Get",
		Pattern:    "/auth/valid",
		HandleFunc: AuthValid,
		Roles:      []int{100},
	},
}
