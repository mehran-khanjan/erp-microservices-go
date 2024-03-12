package middlewares

import (
	"net/http"
	"strconv"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		// before request

		c.Next()
	}
}
