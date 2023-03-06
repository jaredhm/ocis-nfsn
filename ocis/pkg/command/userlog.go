package command

import (
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/command/helper"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/services/userlog/pkg/command"
	"github.com/urfave/cli/v2"
)

// UserlogCommand is the entrypoint for the userlog command.
func UserlogCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     cfg.Userlog.Service.Name,
		Usage:    helper.SubcommandDescription(cfg.Userlog.Service.Name),
		Category: "services",
		Before: func(c *cli.Context) error {
			configlog.Error(parser.ParseConfig(cfg, true))
			cfg.Userlog.Commons = cfg.Commons
			return nil
		},
		Subcommands: command.GetCommands(cfg.Userlog),
	}
}

func init() {
	register.AddCommand(UserlogCommand)
}