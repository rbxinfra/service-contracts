package models

// Proto is a config for a specific proto
type Proto struct {
	// Name represents the name of the proto
	// as well as the name of the csproj
	Name string `yaml:"name"`
	
	// Type defines what type of proto it is,
	// by default this is "messages"
	Type string `yaml:"type"`

	// Path is the relative path to the ProtosPath
	// where to find the proto file for this proto
	// e.g:
	// protos_path: roblox/test/v1
	// protos: [ { name: "test", path: "test.proto" } ] -> roblox/test/v1/test.proto
	Path string `yaml:"path"`

	// DependsOn is a list of protos that this proto
	// depends on within the generation context,
	// be that if another proto was loaded into this context
	// then it can use that.
	DependsOn []string `yaml:"depends_on"`

	// Version is used here just to forward version in this context
	// this is not exposed through Yaml
	Version string `yaml:"-"`

	// OutputPath is used to find where to put the csproj, it is based
	// on the actual path:
	// path: roblox/test/v1/test.proto -> /grpcclients/protos/roblox/test/v1/test.proto/bedev2/Roblox.Test.csproj
	OutputPath string `yaml:"-"`
}

// ContractConfiguration is a configuration.
type ContractConfiguration struct {
	// Version of the contract.
	Version string `yaml:"version"`

	// The base path to search for proto files
	ProtosPath string `yaml:"protos_path"`

	// The list of protos within this service.
	Protos []*Proto `yaml:"protos"`
}
