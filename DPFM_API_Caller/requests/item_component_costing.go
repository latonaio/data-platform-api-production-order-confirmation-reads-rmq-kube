package requests

type ItemComponentCosting struct {
	ProductionOrderConfirmation     int      `json:"ProductionOrderConfirmation"`
	ProductionOrderConfirmationItem int      `json:"ProductionOrderConfirmationItem"`
	BillOfMaterial      	 		int      `json:"BillOfMaterial"`
	BillOfMaterialItem			    int      `json:"BillOfMaterialItem"`
	Currency     			        string 	 `json:"Currency"`
	CostingAmount 			        float32	 `json:"CostingAmount"`
	CreationDate     			    string   `json:"CreationDate"`
	CreationTime  			        string   `json:"CreationTime"`
	LastChangeDate   			    string   `json:"LastChangeDate"`
	LastChangeTime   		        string   `json:"LastChangeTime"`
	IsReleased   			        *bool    `json:"IsReleased"`
	IsLocked     			        *bool    `json:"IsLocked"`
	IsCancelled  			        *bool    `json:"IsCancelled"`
	IsMarkedForDeletion			    *bool    `json:"IsMarkedForDeletion"`
}
