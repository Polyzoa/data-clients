package coingecko

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Assets(t *testing.T) {
	var assets Assets
	err := json.Unmarshal([]byte(MockPartialAssets), &assets)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 10, len(assets))
	assert.Equal(t, "ethereum", assets[0].Id)
	assert.Equal(t, 1, *assets[0].ChainIdentifier)
	assert.Equal(t, "onus", assets[9].Id)
	assert.Nil(t, assets[9].ChainIdentifier)
}

func Test_Coin(t *testing.T) {
	var coin Coin
	err := json.Unmarshal([]byte(Coin4Trump), &coin)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "4trump", coin.Id)
	assert.Equal(t, "4TRUMPJwguiFjfY6PLpwT6SZ5BZmCuzfqWWeJAdR6xP3", coin.Platforms["solana"])
}
