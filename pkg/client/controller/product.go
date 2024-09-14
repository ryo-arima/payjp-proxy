package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapProductCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapProductCmd := &cobra.Command{
		Use:   "product",
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
	bootstrapProductCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapProductCmd
}

func InitCreateProductCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createProductCmd := &cobra.Command{
		Use:   "product",
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
	createProductCmd.Flags().StringP("key", "k", "", "cache key")
	return createProductCmd
}

func InitCreateProductCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createProductCmd := &cobra.Command{
		Use:   "product",
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
	createProductCmd.Flags().StringP("key", "k", "", "cache key")
	return createProductCmd
}

func InitCreateProductCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createProductCmd := &cobra.Command{
		Use:   "product",
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
	createProductCmd.Flags().StringP("key", "k", "", "cache key")
	return createProductCmd
}

func InitGetProductCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getProductCmd := &cobra.Command{
		Use:   "product",
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
	getProductCmd.Flags().StringP("key", "k", "", "cache key")
	return getProductCmd
}

func InitGetProductCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getProductCmd := &cobra.Command{
		Use:   "product",
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
	getProductCmd.Flags().StringP("key", "k", "", "cache key")
	return getProductCmd
}

func InitGetProductCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getProductCmd := &cobra.Command{
		Use:   "product",
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
	getProductCmd.Flags().StringP("key", "k", "", "cache key")
	return getProductCmd
}

func InitUpdateProductCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateProductCmd := &cobra.Command{
		Use:   "product",
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
	updateProductCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProductCmd
}

func InitUpdateProductCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateProductCmd := &cobra.Command{
		Use:   "product",
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
	updateProductCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProductCmd
}

func InitUpdateProductCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateProductCmd := &cobra.Command{
		Use:   "product",
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
	updateProductCmd.Flags().StringP("key", "k", "", "cache key")
	return updateProductCmd
}

func InitDeleteProductCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteProductCmd := &cobra.Command{
		Use:   "product",
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
	deleteProductCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProductCmd
}

func InitDeleteProductCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteProductCmd := &cobra.Command{
		Use:   "product",
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
	deleteProductCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProductCmd
}

func InitDeleteProductCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteProductCmd := &cobra.Command{
		Use:   "product",
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
	deleteProductCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteProductCmd
}