package configuration

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.rbx.com/roblox/service-contracts-generator/flags"
	"github.rbx.com/roblox/service-contracts-generator/models"
	"gopkg.in/yaml.v2"
)

func parseYAMLFile(fileName string) (*models.ContractConfiguration, error) {
	config := &models.ContractConfiguration{}

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	yamlParser := yaml.NewDecoder(yamlFile)
	if err = yamlParser.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func parseFile(fileName string) (*models.ContractConfiguration, error) {
	if fileName == "" {
		return nil, nil
	}

	fileExtension := path.Ext(fileName)

	switch fileExtension {
	case ".yaml":
		return parseYAMLFile(fileName)
	case ".yml":
		return parseYAMLFile(fileName)
	default:
		return nil, nil
	}
}

var (
	ErrProtosDirectoryNotSpecified   = errors.New("protos directory not specified")
	ErrNilContract				     = errors.New("the config couldn't be parsed")
	ErrContractNoVersionSpecified    = errors.New("version must be specified")
	ErrContractNoProtosPathSpecified = errors.New("protos_path must be specified")
	ErrContractNoProtosSpecified     = errors.New("protos must include at least one entry")
	ErrProtoNameNotSpecified         = errors.New("proto must define a name")
	ErrProtoPathNotSpecified         = errors.New("proto must define a path")
)

func validateServiceContract(configuration *models.ContractConfiguration) error {
	if configuration == nil {
		return ErrNilContract
	}

	if configuration.Version == "" {
		return ErrContractNoVersionSpecified
	}

	if len(configuration.Protos) == 0 {
		return ErrContractNoProtosSpecified
	}

	for _, proto := range configuration.Protos {
		if proto.Name == "" {
			return ErrProtoNameNotSpecified
		}

		if proto.Type == "" {
			proto.Type = "messages"
		}

		proto.Type = strings.ToLower(proto.Type)

		if proto.Path == "" {
			return ErrProtoPathNotSpecified
		}

		proto.Path = path.Join(configuration.ProtosPath, proto.Path)
		proto.OutputPath = proto.Path // tpl parser adds the rest to this
		proto.Version = configuration.Version
	}

	return nil
}

// Parse parses the configuration files.
// And produces the full configuration model.
func Parse() (map[string]*models.Proto, error) {
	if *flags.ProtosDirectoryFlag == "" {
		return nil, ErrProtosDirectoryNotSpecified
	}

	var configurations []*models.ContractConfiguration

	err := filepath.Walk(*flags.ProtosDirectoryFlag, func(fileName string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && !*flags.RecurseFlag {
			return filepath.SkipDir
		}

		configuration, err := parseFile(fileName)
		if err != nil {
			return err
		}

		if configuration != nil {
			configurations = append(configurations, configuration)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	configurationMap := make(map[string]*models.Proto)

	for _, entry := range configurations {
		if err := validateServiceContract(entry); err != nil {
			return nil, err
		}

		for _, proto := range entry.Protos {
			if _, ok := configurationMap[proto.Name]; ok != true {
				configurationMap[proto.Name] = proto
			}
		}
	}

	return configurationMap, nil
}
