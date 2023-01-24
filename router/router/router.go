package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ricky-zhf/go-web/router/controller"
	"net/http"
)

//router
//用于处理路由

func SetupRouter() *gin.Engine {

	gin.ForceConsoleColor()
	r := gin.Default()

	//括号内的是变量，所以IndexHandler后不能加括号，在go中，当我们将函数本身赋值给某个变量的时候，是不能加括号的
	r.GET("/", controller.IndexHandler)
	//v1版本-使用路由组
	apiGroup := r.Group("api")
	{
		//blogGroup := apiGroup.Group("/blog.service")

		userGroup := apiGroup.Group("/blog.UserService")
		{
			controller.HandlerUserRouter(userGroup)
		}

		/*添加待办*/
		apiGroup.POST("/blog.Server", controller.BlogServiceRoute)
		//查看所有的待办事项
		apiGroup.GET("/todo", controller.GetTodoList)
		//更新待办
		apiGroup.PUT("/todo/:id", controller.UpdateATodo)
		//删除待办
		apiGroup.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}

func main() {

	// 创建路由
	r := gin.Default()
	// 绑定路由规则，执行的函数（gin.Context，封装了request和response）
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080，Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
