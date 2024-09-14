package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapPlanCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapPlanCmd := &cobra.Command{
		Use:   "plan",
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
	bootstrapPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapPlanCmd
}

func InitCreatePlanCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createPlanCmd := &cobra.Command{
		Use:   "plan",
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
	createPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return createPlanCmd
}

func InitCreatePlanCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createPlanCmd := &cobra.Command{
		Use:   "plan",
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
	createPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return createPlanCmd
}

func InitCreatePlanCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createPlanCmd := &cobra.Command{
		Use:   "plan",
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
	createPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return createPlanCmd
}

func InitGetPlanCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getPlanCmd := &cobra.Command{
		Use:   "plan",
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
	getPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return getPlanCmd
}

func InitGetPlanCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getPlanCmd := &cobra.Command{
		Use:   "plan",
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
	getPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return getPlanCmd
}

func InitGetPlanCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getPlanCmd := &cobra.Command{
		Use:   "plan",
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
	getPlanCmd.Flags().StringP("key", "k", "", "cache key")
	return getPlanCmd
}

func InitUpdatePlanCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updatePlanCmd := &cobra.Command{
		Use:   "plan",
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
	updatePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePlanCmd
}

func InitUpdatePlanCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updatePlanCmd := &cobra.Command{
		Use:   "plan",
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
	updatePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePlanCmd
}

func InitUpdatePlanCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updatePlanCmd := &cobra.Command{
		Use:   "plan",
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
	updatePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return updatePlanCmd
}

func InitDeletePlanCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deletePlanCmd := &cobra.Command{
		Use:   "plan",
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
	deletePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePlanCmd
}

func InitDeletePlanCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deletePlanCmd := &cobra.Command{
		Use:   "plan",
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
	deletePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePlanCmd
}

func InitDeletePlanCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deletePlanCmd := &cobra.Command{
		Use:   "plan",
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
	deletePlanCmd.Flags().StringP("key", "k", "", "cache key")
	return deletePlanCmd
}