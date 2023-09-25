package routers

import (
	"api/internal/domain/dosen"
	"api/internal/domain/fakultas"
	"api/internal/domain/matakuliah"
	"api/internal/domain/mengajar"
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

	dosenRepository := dosen.NewRepository(db)
	dosenService := dosen.Newservice(dosenRepository)
	dosenHandler := handler.NewDosenHandler(dosenService)

	matakuliahRepository := matakuliah.NewRepository(db)
	matakuliahService := matakuliah.Newservice(matakuliahRepository)
	matakuliahHandler := handler.NewMatakuliahHandler(matakuliahService)

	mengajarRepository := mengajar.NewRepository(db)
	mengajarService := mengajar.Newservice(mengajarRepository)
	mengajarHandler := handler.NewMengajarHandler(mengajarService)

	// users ROUTE
	r.api.POST("/users", userHandler.RegisterUser)
	r.api.PUT("/users", userHandler.Update)
	r.api.GET("/users", userHandler.GetAllData)
	r.api.GET("/users-fakultas", userHandler.GetUsersFakultas)

	// fakultas ROUTE
	r.api.POST("/fakultas", fakultasHandler.CreateFakultas)
	r.api.GET("/fakultas", fakultasHandler.GetAllData)
	r.api.GET("/fakultas/:id", fakultasHandler.GetDataById)
	r.api.PUT("/fakultas/:id", fakultasHandler.Update)
	r.api.DELETE("/fakultas/:id", fakultasHandler.Delete)
	r.api.GET("/fakultas/total", fakultasHandler.GetTotal)

	// dosen ROUTE
	r.api.POST("/dosen", dosenHandler.Create)
	r.api.PUT("/dosen/:id", dosenHandler.Update)
	r.api.GET("/dosen", dosenHandler.GetAllData)
	r.api.GET("/dosen/:id", dosenHandler.GetDataById)
	r.api.DELETE("/dosen/:id", dosenHandler.Delete)

	// matakuliah ROUTE
	r.api.POST("/matakuliah", matakuliahHandler.Create)
	r.api.PUT("/matakuliah", matakuliahHandler.Update)
	r.api.GET("/matakuliah", matakuliahHandler.GetAllData)
	r.api.DELETE("/matakuliah/:kode_mk", matakuliahHandler.Delete)

	// mengajar ROUTE
	r.api.POST("/mengajar", mengajarHandler.Create)
	r.api.PUT("/mengajar", mengajarHandler.Update)
	r.api.GET("/mengajar", mengajarHandler.GetAllData)
	r.api.DELETE("/mengajar/:id_dosen/:kode_mk", mengajarHandler.Delete)

}
