package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapCommonCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapCommonCmd := &cobra.Command{
		Use:   "common",
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
	bootstrapCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapCommonCmd
}

func InitCreateCommonCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createCommonCmd := &cobra.Command{
		Use:   "common",
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
	createCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return createCommonCmd
}

func InitCreateCommonCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createCommonCmd := &cobra.Command{
		Use:   "common",
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
	createCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return createCommonCmd
}

func InitCreateCommonCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createCommonCmd := &cobra.Command{
		Use:   "common",
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
	createCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return createCommonCmd
}

func InitGetCommonCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getCommonCmd := &cobra.Command{
		Use:   "common",
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
	getCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return getCommonCmd
}

func InitGetCommonCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getCommonCmd := &cobra.Command{
		Use:   "common",
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
	getCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return getCommonCmd
}

func InitGetCommonCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getCommonCmd := &cobra.Command{
		Use:   "common",
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
	getCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return getCommonCmd
}

func InitUpdateCommonCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateCommonCmd := &cobra.Command{
		Use:   "common",
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
	updateCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCommonCmd
}

func InitUpdateCommonCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateCommonCmd := &cobra.Command{
		Use:   "common",
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
	updateCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCommonCmd
}

func InitUpdateCommonCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateCommonCmd := &cobra.Command{
		Use:   "common",
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
	updateCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return updateCommonCmd
}

func InitDeleteCommonCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteCommonCmd := &cobra.Command{
		Use:   "common",
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
	deleteCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCommonCmd
}

func InitDeleteCommonCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteCommonCmd := &cobra.Command{
		Use:   "common",
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
	deleteCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCommonCmd
}

func InitDeleteCommonCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteCommonCmd := &cobra.Command{
		Use:   "common",
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
	deleteCommonCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteCommonCmd
}