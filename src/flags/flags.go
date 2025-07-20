package flags

import "flag"

var (
	// HelpFlag prints the usage.
	HelpFlag = flag.Bool("help", false, "Print usage.")

	// ProtosDirectoryFlag is the directory where the configuration files are located.
	ProtosDirectoryFlag = flag.String("protos-directory", "./protos", "The directory where the configuration files are located.")

	// OutputDirectoryFlag is the directory where the output files will be located.
	OutputDirectoryFlag = flag.String("output-directory", "./out", "The directory where the output files will be located.")

	// RecurseFlag indicates whether to recurse into subdirectories.
	RecurseFlag = flag.Bool("recurse", true, "Recurse into subdirectories.")

	// CreateSolution indicates whether a solution file should be produced, dotnet must be present on the system.
	CreateSolution = flag.Bool("create-solution", true, "Whether a solution file should be produced, dotnet must be present on the system.")
)

const FlagsUsageString string = `
	[-h|--help]
	[--protos-directory <directory>] [--output-directory <directory>]
	[--recurse] [--create-solution]`
