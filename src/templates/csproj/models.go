package csproj

import "github.rbx.com/roblox/service-contracts-generator/models"

// ProjectModel is the template model for the Project template
type ProjectModel struct {
	// Version is the version of the project
	Version string

	// Product is the product field of the MSBuild project
	Product string

	// Description is the description field of the MSBuild project.
	Description string

	// Type is the type of project, one of messages or api.
	Type string

	// References is the references built based on the depends_on
	// section of the protobuf config.
	References string

	// Protos is the constructed protobuf references field.
	Protos string
}

// BuildModel builds the model based on the proto model.
func BuildModel(proto *models.Proto) *ProjectModel {
	model := &ProjectModel{
		Version: proto.Version,
		Type: proto.Type,
	}

	return model
}
