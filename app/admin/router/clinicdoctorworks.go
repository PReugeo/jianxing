package router

import (
    "github.com/gin-gonic/gin"

    "go-admin/app/admin/middleware"
    "go-admin/app/admin/models"
    "go-admin/app/admin/service/dto"
    "go-admin/common/actions"
    jwt "go-admin/pkg/jwtauth"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerClinicDoctorWorksRouter)
}

// 需认证的路由代码
func registerClinicDoctorWorksRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/clinicdoctorworks").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.ClinicDoctorWorks{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.ClinicDoctorWorksSearch), func() interface{} {
            list := make([]models.ClinicDoctorWorks, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.ClinicDoctorWorksById), nil))
        r.POST("", actions.CreateAction(new(dto.ClinicDoctorWorksControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.ClinicDoctorWorksControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.ClinicDoctorWorksById)))
    }
}
