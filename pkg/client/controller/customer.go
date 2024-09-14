package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapCustomerCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	bootstrapCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapCustomerCmd
}

func InitCreateCustomerCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	createCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return createCustomerCmd
}

func InitCreateCustomerCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	createCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return createCustomerCmd
}

func InitCreateCustomerCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	createCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return createCustomerCmd
}

func InitGetCustomerCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	getCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return getCustomerCmd
}

func InitGetCustomerCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	getCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return getCustomerCmd
}

func InitGetCustomerCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	getCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return getCustomerCmd
}

func InitUpdateCustomerCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	updateCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCustomerCmd
}

func InitUpdateCustomerCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	updateCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCustomerCmd
}

func InitUpdateCustomerCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	updateCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCustomerCmd
}

func InitDeleteCustomerCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	deleteCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCustomerCmd
}

func InitDeleteCustomerCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	deleteCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCustomerCmd
}

func InitDeleteCustomerCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteCustomerCmd := &cobra.Command{
		Use:   "customer",
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
	deleteCustomerCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCustomerCmd
}