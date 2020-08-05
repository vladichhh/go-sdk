package types

// Ping -
type Ping struct {
	Status string `json:"status"`
}

// Vendor -
type Vendor struct {
	ID                    string               `json:"_id"`
	Name                  string               `json:"name"`
	Email                 string               `json:"email"`
	Country               string               `json:"country"`
	City                  string               `json:"city"`
	FirstName             string               `json:"firstName"`
	LastName              string               `json:"lastName"`
	Address               string               `json:"address"`
	Phone                 string               `json:"phone"`
	Zip                   string               `json:"zip"`
	State                 string               `json:"state"`
	VatID                 string               `json:"vatId"`
	TaxID                 string               `json:"taxId"`
	PayoutInfo            []PayoutInfo         `json:"payoutInfo"`
	VendorPrincipal       VendorPrincipal      `json:"vendorPrincipal"`
	ReceiptEmail          string               `json:"receiptEmail"`
	EmailSetup            EmailSetup           `json:"emailSetup"`
	DefaultPayoutCurrency string               `json:"defaultPayoutCurrency"`
	AutoReceipt           bool                 `json:"autoReceipt"`
	AutoInvoice           bool                 `json:"autoInvoice"`
	Frequency             string               `json:"frequency"`
	RawLogo               string               `json:"rawLogo"`
	DocumentNumberFormat  DocumentNumberFormat `json:"documentNumberFormat"`
	Invoice               InvoiceMain          `json:"invoice"`
	CreationDate          string               `json:"creationDate"`
}

// VendorPrincipal -
type VendorPrincipal struct {
	FirstName                    string `json:"firstName"`
	LastName                     string `json:"lastName"`
	Address                      string `json:"address"`
	City                         string `json:"city"`
	Country                      string `json:"country"`
	Zip                          string `json:"zip"`
	Dob                          string `json:"dob"`
	PersonalIdentificationNumber string `json:"personalIdentificationNumber"`
	DriverLicenseNumber          string `json:"driverLicenseNumber"`
	PassportNumber               string `json:"passportNumber"`
	Email                        string `json:"email"`
}

// DocumentNumberFormat -
type DocumentNumberFormat struct {
	Prefix      string `json:"prefix"`
	Placeholder string `json:"placeholder"`
}

// InvoiceMain -
type InvoiceMain struct {
	DocumentFooter string `json:"documentFooter"`
}

// EmailSetup -
type EmailSetup struct {
	Invoice Invoice
	Receipt Receipt
}

// Invoice -
type Invoice struct {
	Subject        string `json:"subject"`
	BodyHeaderText string `json:"bodyHeaderText"`
}

// Receipt -
type Receipt struct {
	Subject        string `json:"subject"`
	BodyHeaderText string `json:"bodyHeaderText"`
}

// PayoutInfo -
type PayoutInfo struct {
	ID                  string `json:"_id"`
	BankID              string `json:"bankId"`
	BankName            string `json:"bankName"`
	Country             string `json:"country"`
	City                string `json:"city"`
	Address             string `json:"address"`
	Zip                 string `json:"zip"`
	BankAccountID       int    `json:"bankAccountId"`
	Iban                string `json:"iban"`
	BankAccountType     string `json:"bankAccountType"`
	BankAccountClass    string `json:"bankAccountClass"`
	NameOnAccount       string `json:"nameOnAccount"`
	PayoutType          string `json:"payoutType"`
	BaseCurrency        string `json:"baseCurrency"`
	MinimalPayoutAmount int    `json:"minimalPayoutAmount"`
	State               string `json:"state"`
	SwiftBic            string `json:"swiftBic"`
}

// Shopper -
type Shopper struct {
	ID                string `json:"_id"`
	Vendor            string `json:"vendor"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	WalletAddress     string `json:"walletAddress"`
	MaliciousAttempts int    `json:"maliciousAttempts"`
	UseLimePayWallet  bool   `json:"useLimePayWallet"`
}

// WalletToken -
type WalletToken struct {
	WalletToken string `json:"walletToken"`
}

// Payment -
type Payment struct {
	ID                  string               `json:"_id"`
	Status              string               `json:"status"`
	Date                string               `json:"date"`
	Currency            string               `json:"currency"`
	Shopper             string               `json:"shopper"`
	Vendor              string               `json:"vendor"`
	Items               []Item               `json:"items"`
	FundTxData          FundTxData           `json:"fundTxData"`
	GenericTransactions []GenericTransaction `json:"genericTransactions"`
	PaymentDetails      PaymentDetails       `json:"paymentDetails"`
	Type                string               `json:"type"`
	LimeToken           string               `json:"limeToken"`
}

// Item -
type Item struct {
	ID          string  `json:"_id"`
	Description string  `json:"description"`
	LineAmount  float32 `json:"lineAmount"`
	Quantity    int     `json:"quantity"`
}

// FundTxData -
type FundTxData struct {
	WeiAmount              string `json:"weiAmount"`
	TokenAmount            string `json:"tokenAmount"`
	AuthorizationSignature string `json:"authorizationSignature"`
	TransactionHash        string `json:"transactionHash"`
	Status                 string `json:"status"`
	Nonce                  string `json:"nonce"`
}

// GenericTransaction -
type GenericTransaction struct {
	ID                string           `json:"_id"`
	To                string           `json:"to"`
	FunctionName      string           `json:"functionName"`
	GasPrice          string           `json:"gasPrice"`
	GasLimit          int              `json:"gasLimit"`
	SignedTransaction string           `json:"signedTransaction"`
	Status            string           `json:"status"`
	TransactionHash   string           `json:"transactionHash"`
	FunctionParams    []FunctionParams `json:"functionParams"`
}

// FunctionParams -
type FunctionParams struct {
	ID    string      `json:"_id"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// PaymentDetails -
type PaymentDetails struct {
	TaxRate     float32    `json:"taxRate"`
	TaxAmount   float32    `json:"taxAmount"`
	BaseAmount  float32    `json:"baseAmount"`
	TotalAmount float32    `json:"totalAmount"`
	CardHolder  CardHolder `json:"cardHolder"`
}

// CardHolder -
type CardHolder struct {
	VatNumber string `json:"vatNumber"`
	Name      string `json:"name"`
	IsCompany bool   `json:"isCompany"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
	Street    string `json:"street"`
}

// SignatureMetadata -
type SignatureMetadata struct {
	Nonce          string `json:"nonce"`
	ShopperAddress string `json:"shopperAddress"`
	EscrowAddress  string `json:"escrowAddress"`
}
