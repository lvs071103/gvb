package api

import "gvb-server/api/settings_api"

type ApiGroup struct{
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)