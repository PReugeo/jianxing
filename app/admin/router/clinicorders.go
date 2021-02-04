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
	routerCheckRole = append(routerCheckRole, registerClinicOrdersRouter)
}

// 需认证的路由代码
func registerClinicOrdersRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/clinicorders").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.ClinicOrders{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.ClinicOrdersSearch), func() interface{} {
            list := make([]models.ClinicOrders, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.ClinicOrdersById), nil))
        r.POST("", actions.CreateAction(new(dto.ClinicOrdersControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.ClinicOrdersControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.ClinicOrdersById)))
    }
}
