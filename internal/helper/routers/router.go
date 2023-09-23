package routers

import (
	"api/internal/domain/fakultas"
	"api/internal/domain/user"
	"api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routers struct {
	api *gin.RouterGroup
}

func RoutersInit(api *gin.RouterGroup) *Routers {
	return &Routers{api: api}
}

func (r *Routers) ExecRouters(db *sqlx.DB) {
	userRepository := user.NewRepository(db)
	userService := user.Newservice(userRepository)
	userHandler := handler.NewUserHandler(userService)

	fakultasRepostory := fakultas.NewRepository(db)
	fakultasService := fakultas.NewService(fakultasRepostory)
	fakultasHandler := handler.NewFakultasHandler(fakultasService)

	// users ROUTE
	r.api.POST("/users", userHandler.RegisterUser)
	r.api.GET("/users", userHandler.GetAllData)
	r.api.GET("/users-fakultas", userHandler.GetUsersFakultas)

	// fakultas ROUTE
	r.api.POST("/fakultas", fakultasHandler.CreateFakultas)
	r.api.GET("/fakultas", fakultasHandler.GetAllData)
	r.api.GET("/fakultas/:id", fakultasHandler.GetDataById)
	r.api.PUT("/fakultas/:id", fakultasHandler.Update)
	r.api.DELETE("/fakultas/:id", fakultasHandler.Delete)
	r.api.GET("/fakultas/total", fakultasHandler.GetTotal)

}
