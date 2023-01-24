package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/dme86/goaws/goaws"
)

func init() {
	RootCmd.AddCommand(newRDSCmd())
}

func newRDSCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rds",
		Short: "Manage RDS resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help() // nolint: errcheck
		},
	}

	cmd.AddCommand(
		newRDSLsCmd(),
	)

	return cmd
}

func newRDSLsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List RDS instances",
		RunE:  runRDSLsCmd,
	}

	flags := cmd.Flags()
	flags.BoolP("quiet", "q", false, "Only display DBInstanceIdentifier")
	flags.StringP("fields", "F", "DBInstanceClass Engine AllocatedStorage StorageTypeIops InstanceCreateTime DBInstanceIdentifier ReadReplicaSource", "Output fields list separated by space")

	viper.BindPFlag("rds.ls.quiet", flags.Lookup("quiet"))   // nolint: errcheck
	viper.BindPFlag("rds.ls.fields", flags.Lookup("fields")) // nolint: errcheck

	return cmd
}

func runRDSLsCmd(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	options := goaws.RDSLsOptions{
		Quiet:  viper.GetBool("rds.ls.quiet"),
		Fields: viper.GetStringSlice("rds.ls.fields"),
	}

	return client.RDSLs(options)
}
