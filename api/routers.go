package api

import (
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/", UserRegistration)
}

func UserRetrieve(router *gin.RouterGroup) {
	router.GET("/:name", UserRetrieval)
}

func UserUpdate(router *gin.RouterGroup) {
	router.PUT("/", UserModification)
}

func UserDelete(router *gin.RouterGroup) {
	router.DELETE("/:name", UserDeletion)
}

func SlotCreate(router *gin.RouterGroup) {
	router.POST("/", SlotCreation)
}

func SlotRetrieve(router *gin.RouterGroup) {
	router.GET("/:name", SlotRetrieval)
}

func SlotRetrieveAll(router *gin.RouterGroup) {
	router.GET("/", SlotRetrievalAll)
}

func SlotDelete(router *gin.RouterGroup) {
	router.DELETE("/:name", SlotDeletion)
}

func ScheduleRetrieve(router *gin.RouterGroup) {
	router.GET("/", ScheduleRetrieval)
}

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api")

	UserRegister(v1.Group("/user"))
	UserRetrieve(v1.Group("/user"))
	UserUpdate(v1.Group("/user"))
	UserDelete(v1.Group("/user"))

	SlotCreate(v1.Group("/slots"))
	SlotRetrieve(v1.Group("/slots"))
	SlotRetrieveAll(v1.Group("/slots"))
	SlotDelete(v1.Group("/slots"))

	ScheduleRetrieve(v1.Group("/schedule"))

	return r
}
