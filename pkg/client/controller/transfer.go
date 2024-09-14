package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapTransferCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	bootstrapTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapTransferCmd
}

func InitCreateTransferCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	createTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return createTransferCmd
}

func InitCreateTransferCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	createTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return createTransferCmd
}

func InitCreateTransferCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	createTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return createTransferCmd
}

func InitGetTransferCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	getTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return getTransferCmd
}

func InitGetTransferCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	getTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return getTransferCmd
}

func InitGetTransferCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	getTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return getTransferCmd
}

func InitUpdateTransferCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	updateTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTransferCmd
}

func InitUpdateTransferCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	updateTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTransferCmd
}

func InitUpdateTransferCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	updateTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTransferCmd
}

func InitDeleteTransferCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	deleteTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTransferCmd
}

func InitDeleteTransferCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	deleteTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTransferCmd
}

func InitDeleteTransferCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteTransferCmd := &cobra.Command{
		Use:   "transfer",
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
	deleteTransferCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTransferCmd
}