package commands

import (
	"io/ioutil"

	"github.com/docker/app/internal"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewrootCmd returns the base root command.
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Short: "Docker Application Packages",
		Long:  `Build and deploy Docker Application Packages.`,
	}
}

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, dockerCli command.Cli) {
	cmd.AddCommand(
		installCmd(dockerCli),
		upgradeCmd(dockerCli),
		uninstallCmd(dockerCli),
		statusCmd(dockerCli),
		initCmd(),
		inspectCmd(dockerCli),
		mergeCmd(dockerCli),
		pushCmd(),
		renderCmd(dockerCli),
		splitCmd(),
		validateCmd(),
		versionCmd(dockerCli),
		completionCmd(dockerCli, cmd),
		bundleCmd(dockerCli),
	)
	if internal.Experimental == "on" {
		cmd.AddCommand(
			pullCmd(),
		)
	}
}

func firstOrEmpty(list []string) string {
	if len(list) != 0 {
		return list[0]
	}
	return ""
}

func muteDockerCli(dockerCli command.Cli) func() {
	stdout := dockerCli.Out()
	stderr := dockerCli.Err()
	dockerCli.Apply(command.WithCombinedStreams(ioutil.Discard))
	return func() {
		dockerCli.Apply(command.WithOutputStream(stdout), command.WithErrorStream(stderr))
	}
}

type parametersOptions struct {
	parametersFiles []string
	overrides       []string
}

func (o *parametersOptions) addFlags(flags *pflag.FlagSet) {
	flags.StringArrayVarP(&o.parametersFiles, "parameters-files", "f", []string{}, "Override parameters files")
	flags.StringArrayVarP(&o.overrides, "set", "s", []string{}, "Override parameters values")
}

type credentialOptions struct {
	targetContext  string
	credentialsets []string
}

func (o *credentialOptions) addFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.targetContext, "target-context", "", "Context on which the application is executed")
	flags.StringArrayVarP(&o.credentialsets, "credential-set", "c", []string{}, "Use a duffle credentialset (either a YAML file, or a credential set present in the duffle credential store)")
}
