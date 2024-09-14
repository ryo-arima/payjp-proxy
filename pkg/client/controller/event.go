package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapEventCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapEventCmd := &cobra.Command{
		Use:   "event",
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
	bootstrapEventCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapEventCmd
}

func InitCreateEventCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createEventCmd := &cobra.Command{
		Use:   "event",
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
	createEventCmd.Flags().StringP("key", "k", "", "cache key")
	return createEventCmd
}

func InitCreateEventCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createEventCmd := &cobra.Command{
		Use:   "event",
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
	createEventCmd.Flags().StringP("key", "k", "", "cache key")
	return createEventCmd
}

func InitCreateEventCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createEventCmd := &cobra.Command{
		Use:   "event",
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
	createEventCmd.Flags().StringP("key", "k", "", "cache key")
	return createEventCmd
}

func InitGetEventCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getEventCmd := &cobra.Command{
		Use:   "event",
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
	getEventCmd.Flags().StringP("key", "k", "", "cache key")
	return getEventCmd
}

func InitGetEventCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getEventCmd := &cobra.Command{
		Use:   "event",
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
	getEventCmd.Flags().StringP("key", "k", "", "cache key")
	return getEventCmd
}

func InitGetEventCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getEventCmd := &cobra.Command{
		Use:   "event",
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
	getEventCmd.Flags().StringP("key", "k", "", "cache key")
	return getEventCmd
}

func InitUpdateEventCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateEventCmd := &cobra.Command{
		Use:   "event",
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
	updateEventCmd.Flags().StringP("key", "k", "", "cache key")
	return updateEventCmd
}

func InitUpdateEventCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateEventCmd := &cobra.Command{
		Use:   "event",
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
	updateEventCmd.Flags().StringP("key", "k", "", "cache key")
	return updateEventCmd
}

func InitUpdateEventCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateEventCmd := &cobra.Command{
		Use:   "event",
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
	updateEventCmd.Flags().StringP("key", "k", "", "cache key")
	return updateEventCmd
}

func InitDeleteEventCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteEventCmd := &cobra.Command{
		Use:   "event",
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
	deleteEventCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteEventCmd
}

func InitDeleteEventCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteEventCmd := &cobra.Command{
		Use:   "event",
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
	deleteEventCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteEventCmd
}

func InitDeleteEventCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteEventCmd := &cobra.Command{
		Use:   "event",
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
	deleteEventCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteEventCmd
}