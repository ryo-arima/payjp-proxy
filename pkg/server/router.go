package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/server/controller"
	"github.com/ryo-arima/payjp-proxy/pkg/server/middleware"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {

	commonRepository := repository.NewCommonRepository(conf)
	commonControllerForPublic := controller.NewCommonControllerForPublic(commonRepository)
	commonControllerForInternal := controller.NewCommonControllerForInternal(commonRepository)
	commonControllerForPrivate := controller.NewCommonControllerForPrivate(commonRepository)

	eventRepository := repository.NewEventRepository(conf)
	eventControllerForInternal := controller.NewEventControllerForInternal(eventRepository)
	eventControllerForPrivate := controller.NewEventControllerForPrivate(eventRepository)

	chargeRepository := repository.NewChargeRepository(conf)
	chargeControllerForInternal := controller.NewChargeControllerForInternal(chargeRepository)
	chargeControllerForPrivate := controller.NewChargeControllerForPrivate(chargeRepository)

	customerRepository := repository.NewCustomerRepository(conf)
	customerControllerForInternal := controller.NewCustomerControllerForInternal(customerRepository)
	customerControllerForPrivate := controller.NewCustomerControllerForPrivate(customerRepository)

	planRepository := repository.NewPlanRepository(conf)
	planControllerForPublic := controller.NewPlanControllerForPublic(planRepository)
	planControllerForInternal := controller.NewPlanControllerForInternal(planRepository)
	planControllerForPrivate := controller.NewPlanControllerForPrivate(planRepository)

	statementRepository := repository.NewStatementRepository(conf)
	statementControllerForInternal := controller.NewStatementControllerForInternal(statementRepository)
	statementControllerForPrivate := controller.NewStatementControllerForPrivate(statementRepository)

	subscriptionRepository := repository.NewSubscriptionRepository(conf)
	subscriptionControllerForPublic := controller.NewSubscriptionControllerForPublic(subscriptionRepository)
	subscriptionControllerForInternal := controller.NewSubscriptionControllerForInternal(subscriptionRepository)
	subscriptionControllerForPrivate := controller.NewSubscriptionControllerForPrivate(subscriptionRepository)

	termRepository := repository.NewTermRepository(conf)
	termControllerForInternal := controller.NewTermControllerForInternal(termRepository)
	termControllerForPrivate := controller.NewTermControllerForPrivate(termRepository)

	balanceRepository := repository.NewBalanceRepository(conf)
	balanceControllerForInternal := controller.NewBalanceControllerForInternal(balanceRepository)
	balanceControllerForPrivate := controller.NewBalanceControllerForPrivate(balanceRepository)

	router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	//common
	publicAPI.GET("/commons", commonControllerForPublic.GetCommons)
	internalAPI.GET("/commons", commonControllerForInternal.GetCommons)
	privateAPI.GET("/commons", commonControllerForPrivate.GetCommons)

	//event
	internalAPI.GET("/events", eventControllerForInternal.GetEvents)
	//internalAPI.POST("/event", eventControllerForInternal.CreateEvent)
	//internalAPI.PUT("/event", eventControllerForInternal.UpdateEvent)
	//internalAPI.DELETE("/event", eventControllerForInternal.DeleteEvent)
	privateAPI.GET("/events", eventControllerForPrivate.GetEvents)
	//privateAPI.POST("/event", eventControllerForPrivate.CreateEvent)
	//privateAPI.PUT("/event", eventControllerForPrivate.UpdateEvent)
	//privateAPI.DELETE("/event", eventControllerForPrivate.DeleteEvent)

	//charge
	internalAPI.GET("/charges", chargeControllerForInternal.GetCharges)
	internalAPI.POST("/charge", chargeControllerForInternal.CreateCharge)
	internalAPI.PUT("/charge", chargeControllerForInternal.UpdateCharge)
	internalAPI.DELETE("/charge", chargeControllerForInternal.DeleteCharge)
	privateAPI.GET("/charges", chargeControllerForPrivate.GetCharges)
	privateAPI.POST("/charge", chargeControllerForPrivate.CreateCharge)
	privateAPI.PUT("/charge", chargeControllerForPrivate.UpdateCharge)
	privateAPI.DELETE("/charge", chargeControllerForPrivate.DeleteCharge)

	//customer
	internalAPI.GET("/customers", customerControllerForInternal.GetCustomers)
	internalAPI.POST("/customer", customerControllerForInternal.CreateCustomer)
	internalAPI.PUT("/customer", customerControllerForInternal.UpdateCustomer)
	internalAPI.DELETE("/customer", customerControllerForInternal.DeleteCustomer)
	privateAPI.GET("/customers", customerControllerForPrivate.GetCustomers)
	privateAPI.POST("/customer", customerControllerForPrivate.CreateCustomer)
	privateAPI.PUT("/customer", customerControllerForPrivate.UpdateCustomer)
	privateAPI.DELETE("/customer", customerControllerForPrivate.DeleteCustomer)

	//plan
	publicAPI.GET("/plans", planControllerForPublic.GetPlans)
	internalAPI.GET("/plans", planControllerForInternal.GetPlans)
	internalAPI.POST("/plan", planControllerForInternal.CreatePlan)
	internalAPI.PUT("/plan", planControllerForInternal.UpdatePlan)
	internalAPI.DELETE("/plan", planControllerForInternal.DeletePlan)
	privateAPI.GET("/plans", planControllerForPrivate.GetPlans)
	privateAPI.POST("/plan", planControllerForPrivate.CreatePlan)
	privateAPI.PUT("/plan", planControllerForPrivate.UpdatePlan)
	privateAPI.DELETE("/plan", planControllerForPrivate.DeletePlan)

	//statement
	internalAPI.GET("/statements", statementControllerForInternal.GetStatements)
	internalAPI.POST("/statement", statementControllerForInternal.CreateStatement)
	internalAPI.PUT("/statement", statementControllerForInternal.UpdateStatement)
	internalAPI.DELETE("/statement", statementControllerForInternal.DeleteStatement)
	privateAPI.GET("/statements", statementControllerForPrivate.GetStatements)
	privateAPI.POST("/statement", statementControllerForPrivate.CreateStatement)
	privateAPI.PUT("/statement", statementControllerForPrivate.UpdateStatement)
	privateAPI.DELETE("/statement", statementControllerForPrivate.DeleteStatement)

	//subscription
	publicAPI.GET("/subscriptions", subscriptionControllerForPublic.GetSubscriptions)
	internalAPI.GET("/subscriptions", subscriptionControllerForInternal.GetSubscriptions)
	internalAPI.POST("/subscription", subscriptionControllerForInternal.CreateSubscription)
	internalAPI.PUT("/subscription", subscriptionControllerForInternal.UpdateSubscription)
	internalAPI.DELETE("/subscription", subscriptionControllerForInternal.DeleteSubscription)
	privateAPI.GET("/subscriptions", subscriptionControllerForPrivate.GetSubscriptions)
	privateAPI.POST("/subscription", subscriptionControllerForPrivate.CreateSubscription)
	privateAPI.PUT("/subscription", subscriptionControllerForPrivate.UpdateSubscription)
	privateAPI.DELETE("/subscription", subscriptionControllerForPrivate.DeleteSubscription)

	//term
	internalAPI.GET("/terms", termControllerForInternal.GetTerms)
	internalAPI.POST("/term", termControllerForInternal.CreateTerm)
	internalAPI.PUT("/term", termControllerForInternal.UpdateTerm)
	internalAPI.DELETE("/term", termControllerForInternal.DeleteTerm)
	privateAPI.GET("/terms", termControllerForPrivate.GetTerms)
	privateAPI.POST("/term", termControllerForPrivate.CreateTerm)
	privateAPI.PUT("/term", termControllerForPrivate.UpdateTerm)
	privateAPI.DELETE("/term", termControllerForPrivate.DeleteTerm)

	//balance
	internalAPI.GET("/balances", balanceControllerForInternal.GetBalances)
	internalAPI.POST("/balance", balanceControllerForInternal.CreateBalance)
	internalAPI.PUT("/balance", balanceControllerForInternal.UpdateBalance)
	internalAPI.DELETE("/balance", balanceControllerForInternal.DeleteBalance)
	privateAPI.GET("/balances", balanceControllerForPrivate.GetBalances)
	privateAPI.POST("/balance", balanceControllerForPrivate.CreateBalance)
	privateAPI.PUT("/balance", balanceControllerForPrivate.UpdateBalance)
	privateAPI.DELETE("/balance", balanceControllerForPrivate.DeleteBalance)

	return router
}
