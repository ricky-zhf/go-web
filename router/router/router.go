package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricky-zhf/go-web/router/controller"
	"log"
)

func SetupRouter(r *gin.Engine) {

	//括号内的是变量，所以IndexHandler后不能加括号，在go中，当我们将函数本身赋值给某个变量的时候，是不能加括号的。
	r.GET("/", controller.IndexHandler)
	//v1版本-使用路由组
	apiGroup := r.Group("api")
	{
		blogGroup := apiGroup.Group("/blog.BlogService")
		{
			controller.HandlerBlogRouter(blogGroup)
		}

		//userGroup := apiGroup.Group("/blog.UserService")
		{
			//...
		}
	}
}

func InitRouter() error {
	// 创建路由
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(addUuid())
	// 绑定路由规则，执行的函数（gin.Context，封装了request和response）
	SetupRouter(r)

	// 3.监听端口，默认在8080，Run("里面不指定端口号默认为8080")
	if err := r.Run(":8000"); err != nil {
		log.Println("r.Run failed|err=", err)
		return err
	}
	return nil
}

func addUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), "traceId", s)
		c.Set("context", ctx)
		c.Writer.Header().Set("X-Request-Id", s)
		log.Println("get in addUuid...|uuid=", s)
		c.Next()
	}
}
