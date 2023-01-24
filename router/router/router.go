package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ricky-zhf/go-web/router/controller"
	"log"
)

func SetupRouter(r *gin.Engine) {

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
		////查看所有的待办事项
		//apiGroup.GET("/todo", controller.GetTodoList)
		////更新待办
		//apiGroup.PUT("/todo/:id", controller.UpdateATodo)
		////删除待办
		//apiGroup.DELETE("/todo/:id", controller.DeleteATodo)
	}
}

func InitRouter() error {
	// 创建路由
	gin.ForceConsoleColor()
	r := gin.Default()

	// 绑定路由规则，执行的函数（gin.Context，封装了request和response）
	SetupRouter(r)

	// 3.监听端口，默认在8080，Run("里面不指定端口号默认为8080")
	if err := r.Run(":8000"); err != nil {
		log.Println("r.Run failed|err=", err)
		return err
	}
	return nil
}
