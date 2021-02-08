package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/patientrecords"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatientRecordsRouter)
}

// 需认证的路由代码
func registerPatientRecordsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &patientrecords.PatientRecords{}
	r := v1.Group("/patientrecords").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPatientRecordsList)
		r.GET("/:id", api.GetPatientRecords)
		//r.GET("/:id", api.GetPatientRecordsByPatientId)
		r.POST("", api.InsertPatientRecords)
		r.PUT("/:id", api.UpdatePatientRecords)
		r.DELETE("", api.DeletePatientRecords)
	}
}
