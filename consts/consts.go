package consts

const (
	HTTPGet                   string = "GET"
	HTTPPost                  string = "POST"
	HTTPPut                   string = "PUT"
	HTTPPatch                 string = "PATCH"
	HTTPDelete                string = "DELETE"
	RoutePing                 string = "/ping"
	RouteGetAllVendors        string = "/vendors"
	RouteCreateShopper        string = "/shoppers"
	RouteGetShopper           string = "/shoppers/%s"
	RouteGetAllShoppers       string = "/shoppers"
	RoutePatchShopper         string = "/shoppers/%s"
	RouteGetWalletToken       string = "/shoppers/%s/walletToken"
	RouteCreateFiatPayment    string = "/payments"
	RouteCreateRelayedPayment string = "/payments/relayed"
	RouteGetPayment           string = "/payments/%s"
	RouteGetAllPayments       string = "/payments"
	RouteSendInvoice          string = "/payments/%s/invoice"
	RouteGetInvoice           string = "/payments/%s/invoice/preview"
	RouteGetReceipt           string = "/payments/%s/receipt"
	RouteGetSignatureMetadata string = "/payments/metadata?shopperId=%s"
)
