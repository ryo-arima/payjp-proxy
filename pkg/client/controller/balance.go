package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapBalanceCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	bootstrapBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapBalanceCmd
}

func InitCreateBalanceCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	createBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return createBalanceCmd
}

func InitCreateBalanceCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	createBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return createBalanceCmd
}

func InitCreateBalanceCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	createBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return createBalanceCmd
}

func InitGetBalanceCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	getBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return getBalanceCmd
}

func InitGetBalanceCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	getBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return getBalanceCmd
}

func InitGetBalanceCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	getBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return getBalanceCmd
}

func InitUpdateBalanceCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	updateBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return updateBalanceCmd
}

func InitUpdateBalanceCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	updateBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return updateBalanceCmd
}

func InitUpdateBalanceCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	updateBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return updateBalanceCmd
}

func InitDeleteBalanceCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	deleteBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteBalanceCmd
}

func InitDeleteBalanceCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	deleteBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteBalanceCmd
}

func InitDeleteBalanceCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteBalanceCmd := &cobra.Command{
		Use:   "balance",
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
	deleteBalanceCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteBalanceCmd
}