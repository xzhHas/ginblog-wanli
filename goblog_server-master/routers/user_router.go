package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gvb_server/api"
	"gvb_server/middleware"
)

var store = cookie.NewStore([]byte("Hb65sd7fk76ul8jg34523"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.POST("qq_login", app.QQLoginView)
	router.GET("qq_login_path", app.QQLoginLinkView) // QQ登录的跳转地址
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemove)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("user_register", app.UserRegisterView) // 用户注册
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)
	router.PUT("user_info", middleware.JwtAuth(), app.UserUpdateNickName)
}
