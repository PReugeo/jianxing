package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/middleware"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatientProfileRouter)
}

// 需认证的路由代码
func registerPatientProfileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/patientprofile").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		model := &models.PatientProfile{}
		r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.PatientProfileSearch), func() interface{} {
			list := make([]models.PatientProfile, 0)
			return &list
		}))
		r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.PatientProfileById), nil))
		r.POST("", actions.CreateAction(new(dto.PatientProfileControl)))
		r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.PatientProfileControl)))
		r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.PatientProfileById)))
	}
}
