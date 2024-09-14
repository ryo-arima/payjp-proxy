package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapSubscriptionCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	bootstrapSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapSubscriptionCmd
}

func InitCreateSubscriptionCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	createSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return createSubscriptionCmd
}

func InitCreateSubscriptionCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	createSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return createSubscriptionCmd
}

func InitCreateSubscriptionCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	createSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return createSubscriptionCmd
}

func InitGetSubscriptionCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	getSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return getSubscriptionCmd
}

func InitGetSubscriptionCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	getSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return getSubscriptionCmd
}

func InitGetSubscriptionCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	getSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return getSubscriptionCmd
}

func InitUpdateSubscriptionCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	updateSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSubscriptionCmd
}

func InitUpdateSubscriptionCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	updateSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSubscriptionCmd
}

func InitUpdateSubscriptionCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	updateSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSubscriptionCmd
}

func InitDeleteSubscriptionCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	deleteSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSubscriptionCmd
}

func InitDeleteSubscriptionCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	deleteSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSubscriptionCmd
}

func InitDeleteSubscriptionCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteSubscriptionCmd := &cobra.Command{
		Use:   "subscription",
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
	deleteSubscriptionCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSubscriptionCmd
}