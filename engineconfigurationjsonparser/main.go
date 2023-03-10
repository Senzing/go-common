// The engineconfigurationjsonparser package helps parse the _ENGINE_CONFIGURATION_JSON.
package engineconfigurationjsonparser

import "context"

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

type EngineConfigurationJsonParser interface {
	GetConfigPath(ctx context.Context) (string, error)
	GetDatabaseUrls(ctx context.Context) ([]string, error)
	GetResourcePath(ctx context.Context) (string, error)
	GetSupportPath(ctx context.Context) (string, error)
	RedactedJson(ctx context.Context) (string, error)
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type EngineConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type EngineConfigurationSql struct {
	Backend    string `json:"BACKEND"`
	Connection string `json:"CONNECTION"`
}

type EngineConfiguration struct {
	Pipeline EngineConfigurationPipeline `json:"PIPELINE"`
	Sql      EngineConfigurationSql      `json:"SQL"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6401xxxx".
const ProductId = 6401
