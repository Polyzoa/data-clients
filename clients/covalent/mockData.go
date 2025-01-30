package covalent

const MockSummary = `{
		"updated_at": "2024-12-08T09:40:01.777472701Z",
		"address": "0x2abf0cfb2f36d093f2dba97bb4214e8fd852ad7c",
		"chain_id": 1,
		"chain_name": "eth-mainnet",
		"items": [
				{
						"total_count": 1781,
						"latest_transaction": {
								"block_signed_at": "2024-10-30T17:00:11Z",
								"tx_hash": "0x3ea9c546a3143d841b1f880031b7bebfc46bd0ab676a5446c359249c5aa9c6e7"
						},
						"earliest_transaction": {
								"block_signed_at": "2024-10-30T17:00:11Z",
								"tx_hash": "0x3ea9c546a3143d841b1f880031b7bebfc46bd0ab676a5446c359249c5aa9c6e7"
						}
				}
		]
}`

const mockEmptySummary = `{
		"updated_at": "2024-12-08T09:40:01.777472701Z",
		"address": "0x2abf0cfb2f36d093f2dba97bb4214e8fd852ad7c",
		"chain_id": 1,
		"chain_name": "eth-mainnet",
		"items": []
}`
