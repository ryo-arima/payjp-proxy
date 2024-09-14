package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapChargeCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	bootstrapChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapChargeCmd
}

func InitCreateChargeCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return createChargeCmd
}

func InitCreateChargeCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return createChargeCmd
}

func InitCreateChargeCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return createChargeCmd
}

func InitGetChargeCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return getChargeCmd
}

func InitGetChargeCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return getChargeCmd
}

func InitGetChargeCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return getChargeCmd
}

func InitUpdateChargeCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return updateChargeCmd
}

func InitUpdateChargeCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return updateChargeCmd
}

func InitUpdateChargeCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "update the value of a key",
		Long:  "update the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return updateChargeCmd
}

func InitDeleteChargeCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteChargeCmd
}

func InitDeleteChargeCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteChargeCmd
}

func InitDeleteChargeCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteChargeCmd := &cobra.Command{
		Use:   "charge",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteChargeCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteChargeCmd
}