package commands

import (
	"aahframe.work"
	"aahframe.work/config"
	"aahframe.work/console"
	"fmt"
	"path/filepath"
)

var CfgTest = console.Command{
	Name:    "cfgtest",
	Usage:   "This is a config loading test",
	Flags: []console.Flag{
		console.StringFlag{
			Name:  "envprofile, e",                                       // long and short posix flag name
			Value: aah.App().Config().StringDefault("env.active", "dev"), // default flag value
			Usage: "Environment profile name to activate (e.g: dev, prod)",
		},
        console.StringFlag{
            Name:  "config, c",                                       // long and short posix flag name
            Usage: "Config file",
        },
	},
	Action: func(ctx *console.Context) error {
		cfg := aah.App().Config()

		aah.App().Log().Info("from aah.conf")
		aah.App().Log().Info(aah.App().Config().StringDefault("object.key", "default-aah.conf"))

		if envCfg, found := cfg.GetSubConfig(fmt.Sprintf("env.%s", ctx.String("envprofile"))); found {
			if err := cfg.Merge(envCfg); err != nil {
				return fmt.Errorf("Unable to env config %s", err)
			}
		}
		aah.App().Log().Infof("from %s profile", ctx.String("envprofile"))
		aah.App().Log().Info(aah.App().Config().StringDefault("object.key", fmt.Sprintf("default-%s", ctx.String("envprofile"))))

		extCfgFile := ctx.String("config")
		if len(extCfgFile) > 0 {
			extCfgPath, err := filepath.Abs(extCfgFile)
			if err != nil {
				return fmt.Errorf("Unable to resolve external config: %s", extCfgFile)
			}
			extCfg, err := config.LoadFile(extCfgPath)
			if err != nil {
				return fmt.Errorf("Unable to load external config, error: %s", err)
			}
			if err = cfg.Merge(extCfg); err != nil {
				return fmt.Errorf("Unable to merge external config %s", err)
			}
		}

		aah.App().Log().Info("final value")
		aah.App().Log().Info(aah.App().Config().StringDefault("object.key", "default-final"))

		return nil
	},
}
