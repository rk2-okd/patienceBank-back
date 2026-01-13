package checkuser

import (
	"github.com/gin-gonic/gin"
)

func CheckUser(c *gin.Context) (int, bool) {
	v, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}
	id, ok := v.(int)
	return id, ok
}
