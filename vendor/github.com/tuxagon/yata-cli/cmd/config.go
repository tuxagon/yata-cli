package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type configArgs struct {
	showKeys bool
}

func (a *configArgs) Parse(ctx *cli.Context) {
	a.showKeys = ctx.Bool("show-keys")
}

// Config allows for editing flattened config values
func Config(ctx *cli.Context) error {
	cargs := &configArgs{}
	cargs.Parse(ctx)

	manager := yata.NewConfigManager()

	if cargs.showKeys {
		return displayConfigKeys(manager)
	}

	args := ctx.Args()
	switch {
	case len(args) == 0:
		return displayConfigAll(manager)
	case len(args) == 1:
		return displayConfigValue(args[0], manager)
	case len(args) > 1:
		return changeConfigValue(args[0], args[1], manager)
	}

	return nil
}

func displayConfigAll(manager *yata.ConfigManager) error {
	configMap, err := manager.GetAll()
	handleError(err)

	for k, v := range configMap {
		if v != "" {
			yata.Printf("%s=%s\n", k, v)
		}
	}

	return nil
}

func displayConfigValue(key string, manager *yata.ConfigManager) error {
	value, err := manager.GetValueForKey(key)
	handleError(err)

	if value != "" {
		yata.Printf("%s\n", value)
	}

	return nil
}

func displayConfigKeys(manager *yata.ConfigManager) error {
	keys, err := manager.GetKeys()
	handleError(err)

	for _, k := range keys {
		yata.Println(k)
	}

	return nil
}

func changeConfigValue(key, newValue string, manager *yata.ConfigManager) error {
	err := manager.SetKey(key, newValue)
	handleError(err)
	return nil
}
