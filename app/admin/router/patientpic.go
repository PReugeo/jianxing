package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/patientpic"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatientPicRouter)
}

// 需认证的路由代码
func registerPatientPicRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &patientpic.PatientPic{}
	r := v1.Group("/patientpic").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPatientPicList)
		r.GET("/:id", api.GetPatientPic)
		r.POST("", api.InsertPatientPic)
		r.PUT("/:id", api.UpdatePatientPic)
		r.DELETE("", api.DeletePatientPic)
	}
}
