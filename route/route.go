package route

import (
	"cygo_iris/controller/v1/user"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	api := app.Party("/").AllowMethods()
	{
		v1 := api.Party("/api/v1")
		{
			v1.PartyFunc("/user", func(p iris.Party) {
				p.Post("/register", user.Register)
				p.Post("/login", user.Login)
				p.Delete("/logout", user.Logout)
			})
		}
	}
}
