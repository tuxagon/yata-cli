package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Config allows the management of any yata configuration
func Config(ctx *cli.Context) error {
	manager := yata.NewConfigManager()

	showKeys := ctx.Bool("show-keys")
	if showKeys {
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
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}

	for k, v := range configMap {
		if v != "" {
			yata.Printf("%s=%s\n", k, v)
		}
	}
	return nil
}

func displayConfigValue(key string, manager *yata.ConfigManager) error {
	value, err := manager.GetValueForKey(key)
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}

	if value != "" {
		yata.Printf("%s\n", value)
	}
	return nil
}

func displayConfigKeys(manager *yata.ConfigManager) error {
	keys, err := manager.GetKeys()
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}
	for _, k := range keys {
		yata.Println(k)
	}
	return nil
}

func changeConfigValue(key, newValue string, manager *yata.ConfigManager) error {
	err := manager.SetKey(key, newValue)
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}
	return nil
}
