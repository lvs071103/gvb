package settings_api

import (
	"gvb-server/models/resp"

	"github.com/gin-gonic/gin"
)

func (SettingsApi)SettingsInfoView(c *gin.Context)  {
	// c.JSON(200, gin.H{"msg": "first example"})
	// resp.Ok(map[string]string{}, "this is ok test",c)
	// resp.OkwithData(map[string]string{"queryset":"hello world"},c)
	// resp.FailedWithCode(resp.SettingsError, c)
	resp.FailedWithCode(2, c)
}