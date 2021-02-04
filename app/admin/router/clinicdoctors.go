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
	routerCheckRole = append(routerCheckRole, registerClinicDoctorsRouter)
}

// 需认证的路由代码
func registerClinicDoctorsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/clinicdoctors").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.ClinicDoctors{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.ClinicDoctorsSearch), func() interface{} {
            list := make([]models.ClinicDoctors, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.ClinicDoctorsById), nil))
        r.POST("", actions.CreateAction(new(dto.ClinicDoctorsControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.ClinicDoctorsControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.ClinicDoctorsById)))
    }
}
