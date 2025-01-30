package goplus

import (
	"fmt"
)

var MockSecurityInfoResults = fmt.Sprintf(`{
    "code": 1,
    "message": "OK",
    "result": {
        "0x4e7e15a6b83718ba79674cc4ce09f7d4a631ec2c": %v
		}
}`, MockSecurityInfo)

const MockSecurityInfo = `{
		"anti_whale_modifiable": "0",
		"buy_tax": "0",
		"can_take_back_ownership": "0",
		"cannot_buy": "0",
		"cannot_sell_all": "0",
		"creator_address": "0x57bb4723e1c9f2a79b7e6462955bceee6be1ce3c",
		"creator_balance": "0",
		"creator_percent": "0.000000",
		"dex": [
				{
						"liquidity_type": "UniV2",
						"name": "UniswapV2",
						"liquidity": "0.00152942",
						"pair": "0x020bbfdedb16d290dfbc1b6f1cbe08e27a152a6e"
				}
		],
		"external_call": "0",
		"hidden_owner": "1",
		"holder_count": "40",
		"holders": [
				{
						"address": "0x020bbfdedb16d290dfbc1b6f1cbe08e27a152a6e",
						"tag": "UniswapV2",
						"is_contract": 1,
						"balance": "1769801017821791.9629249",
						"percent": "1769801.017821791962924900",
						"is_locked": 0
				},
				{
						"address": "0x45034043bb16659e6e354914c31d4cbb59a4e026",
						"tag": "",
						"is_contract": 0,
						"balance": "324697464.290920998",
						"percent": "0.324697464290920998",
						"is_locked": 0
				},
				{
						"address": "0x41c447198813ab44a4e408abfc991975bc15e758",
						"tag": "",
						"is_contract": 0,
						"balance": "105578472.668006244",
						"percent": "0.105578472668006244",
						"is_locked": 0
				},
				{
						"address": "0x4049a7411bec231702b074adfb438d5cf24b700e",
						"tag": "",
						"is_contract": 0,
						"balance": "59278719.719598057",
						"percent": "0.059278719719598057",
						"is_locked": 0
				},
				{
						"address": "0x435f44d53b0844fe54ca2330e30fbfd57ce8741c",
						"tag": "",
						"is_contract": 0,
						"balance": "26386619.164451236",
						"percent": "0.026386619164451236",
						"is_locked": 0
				},
				{
						"address": "0xb1d9e5549d75fe15a7a0eec1e985302cb90fa6b8",
						"tag": "",
						"is_contract": 0,
						"balance": "18999094.674286307",
						"percent": "0.018999094674286307",
						"is_locked": 0
				},
				{
						"address": "0x9e020cf8d0a66cc8a1a072ca523ca65915ef42db",
						"tag": "",
						"is_contract": 0,
						"balance": "10268018.995620591",
						"percent": "0.010268018995620591",
						"is_locked": 0
				},
				{
						"address": "0x174c4cc68ab4d41832ca10b3ef6d3417ece7a9ac",
						"tag": "",
						"is_contract": 0,
						"balance": "8557650.198433829",
						"percent": "0.008557650198433829",
						"is_locked": 0
				},
				{
						"address": "0x1330d5d176a60daf807f0e378dbd5dff56f18449",
						"tag": "",
						"is_contract": 0,
						"balance": "8121041.76974438",
						"percent": "0.008121041769744380",
						"is_locked": 0
				},
				{
						"address": "0x7465b656391f17cb0a631774ac3917512d893c51",
						"tag": "",
						"is_contract": 0,
						"balance": "4488754.119179285",
						"percent": "0.004488754119179285",
						"is_locked": 0
				}
		],
		"honeypot_with_same_creator": "1",
		"is_anti_whale": "0",
		"is_blacklisted": "0",
		"is_honeypot": "1",
		"is_in_dex": "1",
		"is_mintable": "1",
		"is_open_source": "1",
		"is_proxy": "0",
		"is_whitelisted": "0",
		"lp_holder_count": "2",
		"lp_holders": [
				{
						"address": "0x000000000000000000000000000000000000dead",
						"tag": "",
						"value": null,
						"is_contract": 0,
						"balance": "0.921954445729287731",
						"percent": "0.999999999999998915",
						"NFT_list": null,
						"is_locked": 1
				},
				{
						"address": "0x0000000000000000000000000000000000000000",
						"tag": "Null Address",
						"value": null,
						"is_contract": 0,
						"balance": "0.000000000000001",
						"percent": "0.000000000000001084",
						"NFT_list": null,
						"is_locked": 1
				}
		],
		"lp_total_supply": "0.921954445729288731",
		"owner_address": "0x0000000000000000000000000000000000000000",
		"owner_balance": "0",
		"owner_change_balance": "1",
		"owner_percent": "0",
		"personal_slippage_modifiable": "0",
		"selfdestruct": "0",
		"sell_tax": "0",
		"slippage_modifiable": "0",
		"token_name": "KOKODI",
		"token_symbol": "KOKO",
		"total_supply": "1000000000",
		"trading_cooldown": "0",
		"transfer_pausable": "0"
}`
