package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapAccountCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapAccountCmd := &cobra.Command{
		Use:   "account",
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
	bootstrapAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapAccountCmd
}

func InitCreateAccountCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createAccountCmd := &cobra.Command{
		Use:   "account",
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
	createAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return createAccountCmd
}

func InitCreateAccountCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createAccountCmd := &cobra.Command{
		Use:   "account",
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
	createAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return createAccountCmd
}

func InitCreateAccountCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createAccountCmd := &cobra.Command{
		Use:   "account",
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
	createAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return createAccountCmd
}

func InitGetAccountCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getAccountCmd := &cobra.Command{
		Use:   "account",
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
	getAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return getAccountCmd
}

func InitGetAccountCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getAccountCmd := &cobra.Command{
		Use:   "account",
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
	getAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return getAccountCmd
}

func InitGetAccountCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getAccountCmd := &cobra.Command{
		Use:   "account",
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
	getAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return getAccountCmd
}

func InitUpdateAccountCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateAccountCmd := &cobra.Command{
		Use:   "account",
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
	updateAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAccountCmd
}

func InitUpdateAccountCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateAccountCmd := &cobra.Command{
		Use:   "account",
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
	updateAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAccountCmd
}

func InitUpdateAccountCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateAccountCmd := &cobra.Command{
		Use:   "account",
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
	updateAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAccountCmd
}

func InitDeleteAccountCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteAccountCmd := &cobra.Command{
		Use:   "account",
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
	deleteAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAccountCmd
}

func InitDeleteAccountCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteAccountCmd := &cobra.Command{
		Use:   "account",
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
	deleteAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAccountCmd
}

func InitDeleteAccountCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteAccountCmd := &cobra.Command{
		Use:   "account",
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
	deleteAccountCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAccountCmd
}