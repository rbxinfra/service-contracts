package templates

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.rbx.com/roblox/service-contracts-generator/models"
	"github.rbx.com/roblox/service-contracts-generator/templates/csproj"
)

const (
	productMessagesText = "BEDEV2 gRPC Models %s"
	productClientText = "BEDEV2 gRPC API Client %s"
	productServiceText = "BEDEV2 gRPC API %s"

	baseDescriptionText = "Service-Contracts repository auto-generated this"
	messagesDescriptionText = baseDescriptionText + " messages library for %s."
	clientDescriptionText = baseDescriptionText + " service client for %s."
	serviceDescriptionText = baseDescriptionText + " service for %s."

	referenceText = "        <ProjectReference Include=\"$(RootDir)/%s\" />"
	itemGroupText = "<ItemGroup>%s\n    </ItemGroup>"

	protoReferenceText = "<Protobuf ProtoRoot=\"$(ProtoRoot)\" Include=\"$(ProtoRoot)/%s\" Link=\"%s\" GrpcServices=\"%s\" />"

	csprojExtension = ".csproj"
	clientCsprojExtension = ".Client.csproj"
)

func executeTemplateForCsproj(model *csproj.ProjectModel) (string, error) {
	tpl := template.New("csproj")

	var err error

	if tpl, err = tpl.Parse(csproj.ProjectTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func getProduct(proto *models.Proto, isService bool) string {
	if proto.Type == "api" {
		if isService {
			return fmt.Sprintf(productServiceText, proto.Name)
		}

		return fmt.Sprintf(productClientText, proto.Name)
	}

	return fmt.Sprintf(productMessagesText, proto.Name)
}

func getProductDescription(proto *models.Proto, isService bool) string {
	if proto.Type == "api" {
		if isService {
			return fmt.Sprintf(serviceDescriptionText, proto.Name)
		}

		return fmt.Sprintf(clientDescriptionText, proto.Name)
	}

	return fmt.Sprintf(messagesDescriptionText, proto.Name)
}

func buildReference(ref *models.Proto) string {
	return fmt.Sprintf(referenceText, path.Join(ref.OutputPath, ref.Name + csprojExtension))
}

func buildReferencesString(refs []*models.Proto) string {
	if len(refs) == 0 {
		return ""
	}

	references := make([]string, len(refs))

	for _, ref := range refs {
		references = append(references, buildReference(ref))
	}

	return fmt.Sprintf(itemGroupText, strings.Join(references, "\n"))
}

func buildProtoReference(proto *models.Proto, isForService bool) string {
	services := "None"

	if proto.Type == "api" {
		if isForService {
			services = "Service"
		} else {
			services = "Client"
		}
	}

	return fmt.Sprintf(protoReferenceText, proto.Path, proto.Path, services)
}

// ParseForConfiguration parses the templates for the configuration.
func ParseForConfiguration(configMap map[string]*models.Proto) (map[string]string, error) {
	files := make(map[string]string)

	for protoName, proto := range configMap {
		fmt.Printf("Building proto %s\n", protoName)

		protoDependencies := []*models.Proto{}
		if len(proto.DependsOn) > 0 {
			for _, dependencyName := range proto.DependsOn {
				fmt.Printf("Depends on: %s\n", dependencyName)

				if dependency, ok := configMap[dependencyName]; ok {
					protoDependencies = append(protoDependencies, dependency)
				}
			}
		}

		model := csproj.BuildModel(proto)
		model.References = buildReferencesString(protoDependencies)

		if proto.Type == "api" {
			// Gen for both client and service here.

			// service first.
			model.Product = getProduct(proto, true)
			model.Description = getProductDescription(proto, true)
			model.Protos = buildProtoReference(proto, true)

			var err error

			if files[path.Join(proto.OutputPath, "service", protoName + csprojExtension)], err = executeTemplateForCsproj(model); err != nil {
				fmt.Fprintf(os.Stderr, "Skipping rest of proto %s due to error: %+v\n", proto.Name, err)

				continue
			}

			model.Product = getProduct(proto, false)
			model.Description = getProductDescription(proto, false)
			model.Protos = buildProtoReference(proto, false)


			if files[path.Join(proto.OutputPath, "bedev2", protoName + clientCsprojExtension)], err = executeTemplateForCsproj(model); err != nil {
				fmt.Fprintf(os.Stderr, "Skipping rest of proto %s due to error: %+v\n", proto.Name, err)

				continue
			}
		} else {
			model.Product = getProduct(proto, false)
			model.Description = getProductDescription(proto, false)
			model.Protos = buildProtoReference(proto, false)

			var err error
			if files[path.Join(proto.OutputPath, protoName + csprojExtension)], err = executeTemplateForCsproj(model); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating proto %s: %+v\n", proto.Name, err)

				continue
			}
		}
	}

	return files, nil
}
