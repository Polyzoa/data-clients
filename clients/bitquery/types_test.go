package bitquery

import (
	"testing"

	"github.com/massigerardi/go-commons/commons"
)

func TestClient_StatsData(t *testing.T) {
	var data StatsData
	err := commons.LoadFromJson(MockTransferData, &data)
	if err != nil {
		t.Error(err)
	}
	if len(data.Transactions) == 0 || len(data.Senders) == 0 || len(data.Transfers) == 0 {
		t.Error("No data found")
	}
}

func TestClient_ContractData(t *testing.T) {
	var data ContractData
	err := commons.LoadFromJson(ContractDataResponse, &data)
	if err != nil {
		t.Error(err)
	}
	if len(data.Transactions) == 0 {
		t.Error("No data found")
	}
}

func TestClient_HoldersData(t *testing.T) {
	var data HoldersData
	err := commons.LoadFromJson(HoldersDataResponse, &data)
	if err != nil {
		t.Error(err)
	}
	if len(data.BurntBalance) == 0 || len(data.HighHolders) == 0 || len(data.Holders) == 0 || len(data.LowHolders) == 0 {
		t.Error("No data found")
	}
}
