package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapTermCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapTermCmd := &cobra.Command{
		Use:   "term",
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
	bootstrapTermCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapTermCmd
}

func InitCreateTermCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createTermCmd := &cobra.Command{
		Use:   "term",
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
	createTermCmd.Flags().StringP("key", "k", "", "cache key")
	return createTermCmd
}

func InitCreateTermCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createTermCmd := &cobra.Command{
		Use:   "term",
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
	createTermCmd.Flags().StringP("key", "k", "", "cache key")
	return createTermCmd
}

func InitCreateTermCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createTermCmd := &cobra.Command{
		Use:   "term",
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
	createTermCmd.Flags().StringP("key", "k", "", "cache key")
	return createTermCmd
}

func InitGetTermCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getTermCmd := &cobra.Command{
		Use:   "term",
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
	getTermCmd.Flags().StringP("key", "k", "", "cache key")
	return getTermCmd
}

func InitGetTermCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getTermCmd := &cobra.Command{
		Use:   "term",
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
	getTermCmd.Flags().StringP("key", "k", "", "cache key")
	return getTermCmd
}

func InitGetTermCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getTermCmd := &cobra.Command{
		Use:   "term",
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
	getTermCmd.Flags().StringP("key", "k", "", "cache key")
	return getTermCmd
}

func InitUpdateTermCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateTermCmd := &cobra.Command{
		Use:   "term",
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
	updateTermCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTermCmd
}

func InitUpdateTermCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateTermCmd := &cobra.Command{
		Use:   "term",
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
	updateTermCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTermCmd
}

func InitUpdateTermCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateTermCmd := &cobra.Command{
		Use:   "term",
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
	updateTermCmd.Flags().StringP("key", "k", "", "cache key")
	return updateTermCmd
}

func InitDeleteTermCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteTermCmd := &cobra.Command{
		Use:   "term",
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
	deleteTermCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTermCmd
}

func InitDeleteTermCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteTermCmd := &cobra.Command{
		Use:   "term",
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
	deleteTermCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTermCmd
}

func InitDeleteTermCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteTermCmd := &cobra.Command{
		Use:   "term",
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
	deleteTermCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteTermCmd
}