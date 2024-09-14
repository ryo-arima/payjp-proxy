package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapStatementCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapStatementCmd := &cobra.Command{
		Use:   "statement",
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
	bootstrapStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapStatementCmd
}

func InitCreateStatementCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createStatementCmd := &cobra.Command{
		Use:   "statement",
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
	createStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return createStatementCmd
}

func InitCreateStatementCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createStatementCmd := &cobra.Command{
		Use:   "statement",
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
	createStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return createStatementCmd
}

func InitCreateStatementCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createStatementCmd := &cobra.Command{
		Use:   "statement",
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
	createStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return createStatementCmd
}

func InitGetStatementCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getStatementCmd := &cobra.Command{
		Use:   "statement",
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
	getStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return getStatementCmd
}

func InitGetStatementCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getStatementCmd := &cobra.Command{
		Use:   "statement",
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
	getStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return getStatementCmd
}

func InitGetStatementCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getStatementCmd := &cobra.Command{
		Use:   "statement",
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
	getStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return getStatementCmd
}

func InitUpdateStatementCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateStatementCmd := &cobra.Command{
		Use:   "statement",
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
	updateStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return updateStatementCmd
}

func InitUpdateStatementCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateStatementCmd := &cobra.Command{
		Use:   "statement",
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
	updateStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return updateStatementCmd
}

func InitUpdateStatementCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateStatementCmd := &cobra.Command{
		Use:   "statement",
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
	updateStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return updateStatementCmd
}

func InitDeleteStatementCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteStatementCmd := &cobra.Command{
		Use:   "statement",
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
	deleteStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteStatementCmd
}

func InitDeleteStatementCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteStatementCmd := &cobra.Command{
		Use:   "statement",
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
	deleteStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteStatementCmd
}

func InitDeleteStatementCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteStatementCmd := &cobra.Command{
		Use:   "statement",
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
	deleteStatementCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteStatementCmd
}