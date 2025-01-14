package routes

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/controllers/payment"
	"power-wechat-tutorial/controllers/payment/merchant"
	"power-wechat-tutorial/controllers/payment/merchantService"
	"power-wechat-tutorial/controllers/payment/partner"
	"power-wechat-tutorial/controllers/payment/paymentScore"
	"power-wechat-tutorial/controllers/payment/redpack"
	"power-wechat-tutorial/controllers/payment/refund"
	"power-wechat-tutorial/controllers/payment/tax"
	"power-wechat-tutorial/controllers/payment/transfer"
)

func InitPaymentAPIRoutes(r *gin.Engine) {

	r.Static("/wx/payment", "./templates")
	r.POST("/wx/notify", payment.CallbackWXNotify)
	r.POST("/v3/pay/transactions/out-trade-no/5519778939773395659222199111/close", payment.APIMockCloseOrderResponse)

	apiRouterPayment := r.Group("/payment")
	{
		// Handle the pay route
		apiRouterPayment.GET("/order/make", payment.APIMakeOrder)
		apiRouterPayment.GET("/order/make/native", payment.APIMakeOrderNative)
		apiRouterPayment.GET("/order/make/app", payment.APIMakeOrderApp)
		apiRouterPayment.GET("/order/query", payment.APIQueryOrder)
		apiRouterPayment.GET("/order/close", payment.APICloseOrder)
		apiRouterPayment.GET("/order/refund", payment.APIRefundOrder)
		apiRouterPayment.GET("/order/revertOrderByOutTradeNumber", payment.APIRevertOrderByOutTradeNumber)

		apiRouterPayment.GET("/order/refund/query", refund.APIQueryRefundOrder)

		// Handle the partner pay route
		apiRouterPayment.GET("/partner/make", partner.APIMakeOrder)
		apiRouterPayment.GET("/partner/make/native", partner.APIMakeOrderNative)
		apiRouterPayment.GET("/partner/make/app", partner.APIMakeOrderApp)
		apiRouterPayment.GET("/partner/query", partner.APIQueryOrder)
		apiRouterPayment.GET("/partner/close", partner.APICloseOrder)

		// Handle the bill route
		apiRouterPayment.GET("/bill/downloadURL", payment.APIBillDownloadURL)

		apiRouterPayment.GET("/merchant/uploadImg", merchant.APIUploadImg)
		apiRouterPayment.GET("/merchantService/complaints", merchantService.APIComplaints)

		// Handle payment route
		apiRouterPayment.GET("redpack/sendNormal", redpack.APISendNormal)
		apiRouterPayment.GET("redpack/info", redpack.APIQueryRedPack)
		apiRouterPayment.GET("work/sendworkwxredpack", redpack.APIWorkSendWXRedpack)
		apiRouterPayment.GET("transfer/toBalance", transfer.APIToTransfer)
		apiRouterPayment.GET("transfer/toBankCard", transfer.APIToBankCard)
		apiRouterPayment.GET("transfer/queryBalanceOrder", transfer.APIQueryBalanceOrder)
		apiRouterPayment.GET("transfer/batch/batchTransfer", transfer.APIBatchTransfer)
		apiRouterPayment.GET("transfer/batch/queryBatchOrder", transfer.APIQueryBatchOrder)
		apiRouterPayment.GET("transfer/batch/queryBatchOrderDetail", transfer.APIQueryBatchOrderDetail)

		// Handle security route
		apiRouterPayment.GET("security/getRSAPublicKey", payment.APIGetRSAPublicKey)
		apiRouterPayment.POST("security/decryptCertificate", payment.APIDecryptCertificate)
		apiRouterPayment.GET("security/getCertificates", payment.APIGetCertificates)

		apiRouterPayment.GET("tax/applyForCardTemplate", tax.APIApplyForCardTemplate)

		apiRouterPayment.GET("paymentScore/serviceOrder", paymentScore.APIServiceOrder)

		// Handle profitSharing route
		apiRouterPayment.GET("profitSharing/orders", payment.APIOrders)
		apiRouterPayment.GET("profitSharing/addReceiver", payment.APIAddReceiver)

	}

}
