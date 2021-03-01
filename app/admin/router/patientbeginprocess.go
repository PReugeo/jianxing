package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/patientbeginprocess"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatientBeginProcessRouter)
}

// 需认证的路由代码
func registerPatientBeginProcessRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &patientbeginprocess.PatientBeginProcess{}
	r := v1.Group("/patientbeginprocess").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPatientBeginProcessList)
		r.GET("/:id", api.GetPatientBeginProcess)
		r.POST("", api.InsertPatientBeginProcess)
		r.PUT("/:id", api.UpdatePatientBeginProcess)
		r.DELETE("", api.DeletePatientBeginProcess)
	}
}
