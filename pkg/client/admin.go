package client

import (
	"github.com/ryo-arima/payjp-proxy/pkg/client/controller"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAdminUser struct {
	Bootstrap *cobra.Command
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAdminUser() *cobra.Command {
	var rootCmdForAdminUser = &cobra.Command{
		Use:   "payjp-proxy-admin",
		Short: "'payjp-proxy' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAdminUser
}

func InitBaseCmdForAdminUser() BaseCmdForAdminUser {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
	}
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
	baseCmdForAdminUser := BaseCmdForAdminUser{
		Bootstrap: bootstrapCmd,
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAdminUser
}

func ClientForAdminUser(conf config.BaseConfig) {
	rootCmdForAdminUser := InitRootCmdForAdminUser()
	rootCmdForAdminUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAdminUser := InitBaseCmdForAdminUser()

	//bootstrap
	bootstrapCommonCmdForAdminUser := controller.InitBootstrapCommonCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapCommonCmdForAdminUser)
	bootstrapEventCmdForAdminUser := controller.InitBootstrapEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapEventCmdForAdminUser)
	bootstrapChargeCmdForAdminUser := controller.InitBootstrapChargeCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapChargeCmdForAdminUser)
	bootstrapCustomerCmdForAdminUser := controller.InitBootstrapCustomerCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapCustomerCmdForAdminUser)
	bootstrapPlanCmdForAdminUser := controller.InitBootstrapPlanCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapPlanCmdForAdminUser)
	bootstrapStatementCmdForAdminUser := controller.InitBootstrapStatementCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapStatementCmdForAdminUser)
	bootstrapSubscriptionCmdForAdminUser := controller.InitBootstrapSubscriptionCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapSubscriptionCmdForAdminUser)
	bootstrapTransferCmdForAdminUser := controller.InitBootstrapTransferCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapTransferCmdForAdminUser)
	bootstrapTermCmdForAdminUser := controller.InitBootstrapTermCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapTermCmdForAdminUser)
	bootstrapBalanceCmdForAdminUser := controller.InitBootstrapBalanceCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapBalanceCmdForAdminUser)
	bootstrapProductCmdForAdminUser := controller.InitBootstrapProductCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapProductCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Bootstrap)

	//create
	createCommonCmdForAdminUser := controller.InitCreateCommonCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createCommonCmdForAdminUser)
	createEventCmdForAdminUser := controller.InitCreateEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createEventCmdForAdminUser)
	createChargeCmdForAdminUser := controller.InitCreateChargeCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createChargeCmdForAdminUser)
	createCustomerCmdForAdminUser := controller.InitCreateCustomerCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createCustomerCmdForAdminUser)
	createPlanCmdForAdminUser := controller.InitCreatePlanCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createPlanCmdForAdminUser)
	createStatementCmdForAdminUser := controller.InitCreateStatementCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createStatementCmdForAdminUser)
	createSubscriptionCmdForAdminUser := controller.InitCreateSubscriptionCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createSubscriptionCmdForAdminUser)
	createTransferCmdForAdminUser := controller.InitCreateTransferCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createTransferCmdForAdminUser)
	createTermCmdForAdminUser := controller.InitCreateTermCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createTermCmdForAdminUser)
	createBalanceCmdForAdminUser := controller.InitCreateBalanceCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createBalanceCmdForAdminUser)
	createProductCmdForAdminUser := controller.InitCreateProductCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createProductCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Create)

	//get
	getCommonCmdForAdminUser := controller.InitGetCommonCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getCommonCmdForAdminUser)
	getEventCmdForAdminUser := controller.InitGetEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getEventCmdForAdminUser)
	getChargeCmdForAdminUser := controller.InitGetChargeCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getChargeCmdForAdminUser)
	getCustomerCmdForAdminUser := controller.InitGetCustomerCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getCustomerCmdForAdminUser)
	getPlanCmdForAdminUser := controller.InitGetPlanCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getPlanCmdForAdminUser)
	getStatementCmdForAdminUser := controller.InitGetStatementCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getStatementCmdForAdminUser)
	getSubscriptionCmdForAdminUser := controller.InitGetSubscriptionCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getSubscriptionCmdForAdminUser)
	getTransferCmdForAdminUser := controller.InitGetTransferCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getTransferCmdForAdminUser)
	getTermCmdForAdminUser := controller.InitGetTermCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getTermCmdForAdminUser)
	getBalanceCmdForAdminUser := controller.InitGetBalanceCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getBalanceCmdForAdminUser)
	getProductCmdForAdminUser := controller.InitGetProductCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getProductCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Get)

	//update
	updateCommonCmdForAdminUser := controller.InitUpdateCommonCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateCommonCmdForAdminUser)
	updateEventCmdForAdminUser := controller.InitUpdateEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateEventCmdForAdminUser)
	updateChargeCmdForAdminUser := controller.InitUpdateChargeCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateChargeCmdForAdminUser)
	updateCustomerCmdForAdminUser := controller.InitUpdateCustomerCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateCustomerCmdForAdminUser)
	updatePlanCmdForAdminUser := controller.InitUpdatePlanCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updatePlanCmdForAdminUser)
	updateStatementCmdForAdminUser := controller.InitUpdateStatementCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateStatementCmdForAdminUser)
	updateSubscriptionCmdForAdminUser := controller.InitUpdateSubscriptionCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateSubscriptionCmdForAdminUser)
	updateTransferCmdForAdminUser := controller.InitUpdateTransferCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateTransferCmdForAdminUser)
	updateTermCmdForAdminUser := controller.InitUpdateTermCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateTermCmdForAdminUser)
	updateBalanceCmdForAdminUser := controller.InitUpdateBalanceCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateBalanceCmdForAdminUser)
	updateProductCmdForAdminUser := controller.InitUpdateProductCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateProductCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Update)

	//delete
	deleteCommonCmdForAdminUser := controller.InitDeleteCommonCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteCommonCmdForAdminUser)
	deleteEventCmdForAdminUser := controller.InitDeleteEventCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteEventCmdForAdminUser)
	deleteChargeCmdForAdminUser := controller.InitDeleteChargeCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteChargeCmdForAdminUser)
	deleteCustomerCmdForAdminUser := controller.InitDeleteCustomerCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteCustomerCmdForAdminUser)
	deletePlanCmdForAdminUser := controller.InitDeletePlanCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deletePlanCmdForAdminUser)
	deleteStatementCmdForAdminUser := controller.InitDeleteStatementCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteStatementCmdForAdminUser)
	deleteSubscriptionCmdForAdminUser := controller.InitDeleteSubscriptionCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteSubscriptionCmdForAdminUser)
	deleteTransferCmdForAdminUser := controller.InitDeleteTransferCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteTransferCmdForAdminUser)
	deleteTermCmdForAdminUser := controller.InitDeleteTermCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteTermCmdForAdminUser)
	deleteBalanceCmdForAdminUser := controller.InitDeleteBalanceCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteBalanceCmdForAdminUser)
	deleteProductCmdForAdminUser := controller.InitDeleteProductCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteProductCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Delete)

	rootCmdForAdminUser.Execute()
}
