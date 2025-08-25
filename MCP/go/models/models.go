package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetLatestConfigurationRequest represents the GetLatestConfigurationRequest schema from the OpenAPI specification
type GetLatestConfigurationRequest struct {
}

// StartConfigurationSessionrequest represents the StartConfigurationSessionrequest schema from the OpenAPI specification
type StartConfigurationSessionrequest struct {
	Requiredminimumpollintervalinseconds int `json:"RequiredMinimumPollIntervalInSeconds,omitempty"` // Sets a constraint on a session. If you specify a value of, for example, 60 seconds, then the client that established the session can't call <a>GetLatestConfiguration</a> more frequently than every 60 seconds.
	Applicationidentifier string `json:"ApplicationIdentifier"` // The application ID or the application name.
	Configurationprofileidentifier string `json:"ConfigurationProfileIdentifier"` // The configuration profile ID or the configuration profile name.
	Environmentidentifier string `json:"EnvironmentIdentifier"` // The environment ID or the environment name.
}

// GetLatestConfigurationResponse represents the GetLatestConfigurationResponse schema from the OpenAPI specification
type GetLatestConfigurationResponse struct {
	Configuration interface{} `json:"Configuration,omitempty"`
}

// StartConfigurationSessionRequest represents the StartConfigurationSessionRequest schema from the OpenAPI specification
type StartConfigurationSessionRequest struct {
	Applicationidentifier interface{} `json:"ApplicationIdentifier"`
	Configurationprofileidentifier interface{} `json:"ConfigurationProfileIdentifier"`
	Environmentidentifier interface{} `json:"EnvironmentIdentifier"`
	Requiredminimumpollintervalinseconds interface{} `json:"RequiredMinimumPollIntervalInSeconds,omitempty"`
}

// StartConfigurationSessionResponse represents the StartConfigurationSessionResponse schema from the OpenAPI specification
type StartConfigurationSessionResponse struct {
	Initialconfigurationtoken interface{} `json:"InitialConfigurationToken,omitempty"`
}
