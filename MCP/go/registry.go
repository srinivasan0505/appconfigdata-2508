package main

import (
	"github.com/aws-appconfig-data/mcp-server/config"
	"github.com/aws-appconfig-data/mcp-server/models"
	tools_configuration_configuration_token "github.com/aws-appconfig-data/mcp-server/tools/configuration_configuration_token"
	tools_configurationsessions "github.com/aws-appconfig-data/mcp-server/tools/configurationsessions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_configuration_configuration_token.CreateGetlatestconfigurationTool(cfg),
		tools_configurationsessions.CreateStartconfigurationsessionTool(cfg),
	}
}
