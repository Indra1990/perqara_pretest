package usecase

type (
	VendingMachine struct {
		ID    string `json:"vandingMachineId"`
		Item  string `json:"item"`
		Price string `json:"price"`
	}
)
