package requests

type ItemDoc struct {
	ProductionOrderConfirmation          int     `json:"ProductionOrderConfirmation"`
	ProductionOrderConfirmationItem      int     `json:"ProductionOrderConfirmationItem"`
	DocType                  		     string  `json:"DocType"`
	DocVersionID             			 int     `json:"DocVersionID"`
	DocID                   			 string  `json:"DocID"`
	FileExtension        			     string  `json:"FileExtension"`
	FileName          			         *string `json:"FileName"`
	FilePath           			         *string `json:"FilePath"`
	DocIssuerBusinessPartner 			 *int    `json:"DocIssuerBusinessPartner"`
}
