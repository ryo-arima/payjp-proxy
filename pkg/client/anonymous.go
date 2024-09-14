package client

import (
	"github.com/ryo-arima/payjp-proxy/pkg/client/controller"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get       *cobra.Command
}

func InitRootCmdForAnonymousUser() *cobra.Command {
	var rootCmdForAnonymousUser = &cobra.Command{
		Use:   "payjp-proxy-anonymous",
		Short: "'payjp-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAnonymousUser
}

func InitBaseCmdForAnonymousUser() BaseCmdForAnonymousUser {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:       getCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getCommonCmdForAnonymousUser := controller.InitGetCommonCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getCommonCmdForAnonymousUser)
	getEventCmdForAnonymousUser := controller.InitGetEventCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getEventCmdForAnonymousUser)
	getChargeCmdForAnonymousUser := controller.InitGetChargeCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getChargeCmdForAnonymousUser)
	getCustomerCmdForAnonymousUser := controller.InitGetCustomerCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getCustomerCmdForAnonymousUser)
	getPlanCmdForAnonymousUser := controller.InitGetPlanCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getPlanCmdForAnonymousUser)
	getStatementCmdForAnonymousUser := controller.InitGetStatementCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getStatementCmdForAnonymousUser)
	getSubscriptionCmdForAnonymousUser := controller.InitGetSubscriptionCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getSubscriptionCmdForAnonymousUser)
	getTransferCmdForAnonymousUser := controller.InitGetTransferCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getTransferCmdForAnonymousUser)
	getTermCmdForAnonymousUser := controller.InitGetTermCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getTermCmdForAnonymousUser)
	getBalanceCmdForAnonymousUser := controller.InitGetBalanceCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getBalanceCmdForAnonymousUser)
	getProductCmdForAnonymousUser := controller.InitGetProductCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getProductCmdForAnonymousUser)
	getAccountCmdForAnonymousUser := controller.InitGetAccountCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getAccountCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}