package client

import (
	"github.com/ryo-arima/payjp-proxy/pkg/client/controller"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAppUser struct {
	Create *cobra.Command
	Get    *cobra.Command
	Update *cobra.Command
	Delete *cobra.Command
}

func InitRootCmdForAppUser() *cobra.Command {
	var rootCmdForAppUser = &cobra.Command{
		Use:   "payjp-proxy-app",
		Short: "'payjp-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAppUser
}

func InitBaseCmdForAppUser() BaseCmdForAppUser {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the value of a key",
		Long:  "update the value of a key",
	}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
	}
	baseCmdForAppUser := BaseCmdForAppUser{
		Create: createCmd,
		Get:    getCmd,
		Update: updateCmd,
		Delete: deleteCmd,
	}
	return baseCmdForAppUser
}

func ClientForAppUser(conf config.BaseConfig) {
	rootCmdForAppUser := InitRootCmdForAppUser()
	rootCmdForAppUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAppUser := InitBaseCmdForAppUser()

	//create
	createCommonCmdForAppUser := controller.InitCreateCommonCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createCommonCmdForAppUser)
	createEventCmdForAppUser := controller.InitCreateEventCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createEventCmdForAppUser)
	createChargeCmdForAppUser := controller.InitCreateChargeCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createChargeCmdForAppUser)
	createCustomerCmdForAppUser := controller.InitCreateCustomerCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createCustomerCmdForAppUser)
	createPlanCmdForAppUser := controller.InitCreatePlanCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createPlanCmdForAppUser)
	createStatementCmdForAppUser := controller.InitCreateStatementCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createStatementCmdForAppUser)
	createSubscriptionCmdForAppUser := controller.InitCreateSubscriptionCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createSubscriptionCmdForAppUser)
	createTransferCmdForAppUser := controller.InitCreateTransferCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createTransferCmdForAppUser)
	createTermCmdForAppUser := controller.InitCreateTermCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createTermCmdForAppUser)
	createBalanceCmdForAppUser := controller.InitCreateBalanceCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createBalanceCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Create)

	//get
	getCommonCmdForAppUser := controller.InitGetCommonCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getCommonCmdForAppUser)
	getEventCmdForAppUser := controller.InitGetEventCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getEventCmdForAppUser)
	getChargeCmdForAppUser := controller.InitGetChargeCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getChargeCmdForAppUser)
	getCustomerCmdForAppUser := controller.InitGetCustomerCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getCustomerCmdForAppUser)
	getPlanCmdForAppUser := controller.InitGetPlanCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getPlanCmdForAppUser)
	getStatementCmdForAppUser := controller.InitGetStatementCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getStatementCmdForAppUser)
	getSubscriptionCmdForAppUser := controller.InitGetSubscriptionCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getSubscriptionCmdForAppUser)
	getTransferCmdForAppUser := controller.InitGetTransferCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getTransferCmdForAppUser)
	getTermCmdForAppUser := controller.InitGetTermCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getTermCmdForAppUser)
	getBalanceCmdForAppUser := controller.InitGetBalanceCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getBalanceCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Get)

	//update
	updateCommonCmdForAppUser := controller.InitUpdateCommonCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateCommonCmdForAppUser)
	updateEventCmdForAppUser := controller.InitUpdateEventCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateEventCmdForAppUser)
	updateChargeCmdForAppUser := controller.InitUpdateChargeCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateChargeCmdForAppUser)
	updateCustomerCmdForAppUser := controller.InitUpdateCustomerCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateCustomerCmdForAppUser)
	updatePlanCmdForAppUser := controller.InitUpdatePlanCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updatePlanCmdForAppUser)
	updateStatementCmdForAppUser := controller.InitUpdateStatementCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateStatementCmdForAppUser)
	updateSubscriptionCmdForAppUser := controller.InitUpdateSubscriptionCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateSubscriptionCmdForAppUser)
	updateTransferCmdForAppUser := controller.InitUpdateTransferCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateTransferCmdForAppUser)
	updateTermCmdForAppUser := controller.InitUpdateTermCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateTermCmdForAppUser)
	updateBalanceCmdForAppUser := controller.InitUpdateBalanceCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateBalanceCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Update)

	//delete
	deleteCommonCmdForAppUser := controller.InitDeleteCommonCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteCommonCmdForAppUser)
	deleteEventCmdForAppUser := controller.InitDeleteEventCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteEventCmdForAppUser)
	deleteChargeCmdForAppUser := controller.InitDeleteChargeCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteChargeCmdForAppUser)
	deleteCustomerCmdForAppUser := controller.InitDeleteCustomerCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteCustomerCmdForAppUser)
	deletePlanCmdForAppUser := controller.InitDeletePlanCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deletePlanCmdForAppUser)
	deleteStatementCmdForAppUser := controller.InitDeleteStatementCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteStatementCmdForAppUser)
	deleteSubscriptionCmdForAppUser := controller.InitDeleteSubscriptionCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteSubscriptionCmdForAppUser)
	deleteTransferCmdForAppUser := controller.InitDeleteTransferCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteTransferCmdForAppUser)
	deleteTermCmdForAppUser := controller.InitDeleteTermCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteTermCmdForAppUser)
	deleteBalanceCmdForAppUser := controller.InitDeleteBalanceCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteBalanceCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Delete)

	rootCmdForAppUser.Execute()
}
