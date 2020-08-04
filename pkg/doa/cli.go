package doa

import (
	"log"
	"os"
	"path"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CLI holds specific information to run the CLI
type CLI struct {
	// Configuration is a Viper instance
	Configuration *viper.Viper
	ConfigFile    string
}

// New generates a new CLI structure and applies the configuration defaults such as:
//    * the default configuration name (~/.doa.yaml)
//    * the default configuration type (yaml)
//    * the default configuration paths (~/)
//    * the default configuration filesystems permissions (0600)
// New also generates a single default application entry (terraform) so that the config file template can be exposed to the user outside of documentation
//
// Arguments:
//     None
//
// Returns:
//     *CLI: A pointer to a new instance of the CLI structure
func New() *CLI {
	udir := func() string {
		d, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		return d
	}()

	v := viper.New()
	v.SetConfigName(".doa")
	v.SetConfigType("yaml")
	v.AddConfigPath(udir)
	v.SetConfigPermissions(0600)

	v.SetDefault("doa_api_version", "v1")
	v.SetDefault("install_location", path.Join(udir, "doa", "bin"))
	v.SetDefault("spec.tools", []map[string]interface{}{
		map[string]interface{}{
			"name":         "terraform",
			"src_location": "https://releases.hashicorp.com/terraform",
			"version":      "latest",
		},
	})

	return &CLI{
		Configuration: v,
	}
}

// Run will run the CLI.
// If the CLI fails to run for any reason, the app will quit with an error message.
//
// Arguments:
//     None
//
// Returns:
//     None
func (C *CLI) Run() {
	var (
		rootCMD = &cobra.Command{
			Use:     "doa",
			Version: "0.1.0",
			Short:   "DOA - DevOps Assistant",
			Long:    "The DevOps Assistant is a tool designed to make getting all of the tooling that you use as a DevOps Engineer and/or an SRE",
			Args:    cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				if C.ConfigFile != "~/.doa.yaml" {
					C.Configuration.SetConfigFile(C.ConfigFile)
				}
				if err := C.Configuration.ReadInConfig(); err != nil {
					log.Fatalf("error reading the configuration file: %+v\nIf you didn't expect to have to provide a configuration file, you may have wanted the \"install\" or \"remove\" commands.", err)
				}
				log.Print(C.Configuration)
			},
		}
		initCMD = &cobra.Command{
			Use:   "init",
			Short: "Initialize DOA",
			Args:  cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				log.Println("Initializing DOA...")
				err := C.Configuration.SafeWriteConfig()
				if reflect.TypeOf(err) == reflect.TypeOf(viper.ConfigFileAlreadyExistsError("")) {
					log.Fatalf("Skipping initialization: %+v", err)
				} else if err != nil {
					log.Fatalf("error writing the DOA configuration file: %+v", err)
				}
			},
		}
		installCMD = &cobra.Command{
			Use:   "install",
			Short: "Install tools in an AdHoc fashion",
			Run: func(ccmd *cobra.Command, args []string) {
				log.Fatalln("Not yet implimented")
			},
		}
		removeCMD = &cobra.Command{
			Use:   "remove",
			Short: "Remove tools in an AdHoc fashion",
			Run: func(ccmd *cobra.Command, args []string) {
				log.Fatalln("Not yet implimented")
			},
		}
	)

	rootCMD.AddCommand(initCMD, installCMD, removeCMD)

	rootCMD.Flags().StringVarP(&C.ConfigFile, "config", "f", "~/.doa.yaml", "The configuration file to read from")

	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
