# Service Contracts

Repository that generates C# projects for Roblox service contracts

## Building

Ensure you have [Go 1.20+](https://go.dev/dl/)

1. Clone the repository via `git`:

    ```txt
    git clone git@github.rbx.com:Roblox/service-contracts.git
    cd service-contracts
    ```

2. Build via [make](https://www.gnu.org/software/make/)

    ```txt
    make build-debug
    ```

## Usage

`cd src && go run main.go --help` (use the build binary found in the bin directory if you downloaded a prebuilt or built it yourself)

```txt
Usage: service-contracts-generator
Build Mode: 
Commit:  
        [-h|--help]
        [--protos-directory <directory>] [--output-directory <directory>]
        [--recurse] [--create-solution]

  -create-solution
        Whether a solution file should be produced, dotnet must be present on the system. (default true)
  -help
        Print usage.
  -output-directory string
        The directory where the output files will be located. (default "./out")
  -protos-directory string
        The directory where the configuration files are located. (default "./protos")
  -recurse
        Recurse into subdirectories. (default true)
```
# Notice

## Usage of Roblox, or any of its assets.

# ***This project is not affiliated with Roblox Corporation.***

The usage of the name Roblox and any of its assets is purely for the purpose of providing a clear understanding of the project's purpose and functionality. This project is not endorsed by Roblox Corporation, and is not intended to be u
sed for any commercial purposes.

Any code in this project was soley produced with or without the assistance of error traces and/or behaviour analysis of public facing APIs.
