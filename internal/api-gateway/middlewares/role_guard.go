package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get user information from the context
		user, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract role from user information
		authUser, ok := user.(AuthenticatedUser)

		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		userRole := authUser.Role

		// Check if the user has any of the allowed roles
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				// If the user has any allowed role, proceed to the next middleware
				ctx.Next()
				return
			}
		}

		// If none of the allowed roles match, deny access
		ctx.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
	}
}
