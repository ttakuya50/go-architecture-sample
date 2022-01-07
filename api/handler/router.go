package handler

//
//import "github.com/gin-gonic/gin"
//
//type Router struct {
//	engine      *gin.Engine
//	userHandler *UserHandler
//}
//
//func NewRouter(engine *gin.Engine, userHandler *UserHandler) *Router {
//	return &Router{
//		engine:      engine,
//		userHandler: userHandler,
//	}
//}
//
//func (r *Router) ApiRoutes() error {
//	r.engine.POST("/user", r.userHandler.Register)
//	r.engine.DELETE("/user", r.userHandler.Delete)
//	r.engine.POST("/user/list", r.userHandler.AddList)
//	//r.engine.POST("/user/list", openapi.UserListPost)
//	return r.engine.Run("127.0.0.1:8888")
//	return nil
//}
