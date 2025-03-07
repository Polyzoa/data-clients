package bitquery

import (
	"fmt"
)

var TransferResponse = fmt.Sprintf("{\"data\": %s}", MockTransferData)

const MockTransferData = `{
    "senders": [
      {
        "average": 105.09108861604719,
        "fails": "332204",
        "success": "38953254",
        "token": {
          "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
        },
        "total": "39285458"
      },
      {
        "average": 92.86240902780614,
        "fails": "1649443",
        "success": "202351216",
        "token": {
          "address": "0xdac17f958d2ee523a2206206994597c13d831ec7"
        },
        "total": "204000659"
      }
    ],
    "transactions": [
      {
        "fails": "2250272",
        "success": "215309779",
        "token": {
          "address": "0xdac17f958d2ee523a2206206994597c13d831ec7"
        },
        "total": "217560051"
      },
      {
        "fails": "377602",
        "success": "47005146",
        "token": {
          "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
        },
        "total": "47382748"
      }
    ],
    "transfer_stats": [
      {
        "average": 113571.8337436754,
        "median": 1027.2622780000002,
        "receivers": "15318936",
        "standard_deviation": 2354586.7529895436,
        "token": {
          "Currency": {
            "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
          }
        }
      },
      {
        "average": 31562.866076022685,
        "median": 689.1262725,
        "receivers": "41937016",
        "standard_deviation": 1155362.0342683846,
        "token": {
          "Currency": {
            "address": "0xdac17f958d2ee523a2206206994597c13d831ec7"
          }
        }
      }
    ],
    "transfers": [
      {
        "average": 105.09108861604719,
        "fails": "332204",
        "success": "38953254",
        "token": {
          "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
        },
        "total": "39285458"
      },
      {
        "average": 92.86240902780614,
        "fails": "1649443",
        "success": "202351216",
        "token": {
          "address": "0xdac17f958d2ee523a2206206994597c13d831ec7"
        },
        "total": "204000659"
      }
    ]
}`

var ContractResponse = fmt.Sprintf("{\"data\": %s}", ContractDataResponse)

const ContractDataResponse = `{
  "transactions": [
    {
      "hash": "0x2f1c5c2b44f771e942a8506148e256f94f1a464babc938ae0690c6e34cd79190",
      "sender": {
        "address": "0x36928500bc1dcd7af6a2b4008875cc336b927d57",
        "smartContract": {
          "contractType": "Generic"
        }
      },
      "creates": {
        "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
        "smartContract": {
          "currency": {
            "decimals": 6,
            "name": "Tether USD",
            "symbol": "USDT",
            "tokenType": "ERC20"
          },
          "contractType": "Token"
        }
      },
      "block": {
        "timestamp": {
          "unixtime": 1511829681
        }
      }
    },
{
      "hash": "0x2f1c5c2b44f771e942a8506148e256f94f1a464babc938ae0690c6e34cd79190",
      "sender": {
        "address": "0x36928500bc1dcd7af6a2b4008875cc336b927d57",
        "smartContract": {
          "contractType": "Generic"
        }
      },
      "creates": {
        "address": "0xdac17f958d2ee523a2206206994597c13d831ec8",
        "smartContract": {
          "currency": {
            "decimals": 6,
            "name": "Tether USD",
            "symbol": "USDT",
            "tokenType": "ERC20"
          },
          "contractType": "Generic"
        }
      },
      "block": {
        "timestamp": {
          "unixtime": 1511829681
        }
      }
    },
		{
      "hash": "0xe7e0fe390354509cd08c9a0168536938600ddc552b3f7cb96030ebef62e75895",
      "sender": {
        "address": "0x95ba4cf87d6723ad9c0db21737d862be80e93911",
        "smartContract": {
          "contractType": null
        }
      },
      "creates": {
        "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
        "smartContract": {
          "currency": {
            "decimals": 6,
            "name": "USD//C",
            "symbol": "USDC",
            "tokenType": "ERC20"
          },
          "contractType": "Token"
        }
      },
      "block": {
        "timestamp": {
          "unixtime": 1533324504
        }
      }
    }
  ]
}`

// var ContractResponseWithGeneric = fmt.Sprintf("{\"data\": %s}", ContractDataResponseWithGeneric)

// const ContractDataResponseWithGeneric = `{
//   "transactions": [
//     {
//       "hash": "0x2f1c5c2b44f771e942a8506148e256f94f1a464babc938ae0690c6e34cd79190",
//       "sender": {
//         "address": "0x36928500bc1dcd7af6a2b4008875cc336b927d57",
//         "smartContract": {
//           "contractType": null
//         }
//       },
//       "creates": {
//         "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
//         "smartContract": {
//           "currency": {
//             "decimals": 6,
//             "name": "Tether USD",
//             "symbol": "USDT",
//             "tokenType": "ERC20"
//           },
//           "contractType": "Generic"
//         }
//       },
//       "block": {
//         "timestamp": {
//           "unixtime": 1511829681
//         }
//       }
//     },
//     {
//       "hash": "0x2f1c5c2b44f771e942a8506148e256f94f1a464babc938ae0690c6e34cd79190",
//       "sender": {
//         "address": "0x36928500bc1dcd7af6a2b4008875cc336b927d57",
//         "smartContract": {
//           "contractType": null
//         }
//       },
//       "creates": {
//         "address": "0x0755c158ddfad8a50cb9d98005954ee2ca9289d6",
//         "smartContract": {
//           "currency": {
//             "decimals": 6,
//             "name": "Tether USD",
//             "symbol": "USDT",
//             "tokenType": "ERC711"
//           },
//           "contractType": "Token"
//         }
//       },
//       "block": {
//         "timestamp": {
//           "unixtime": 1511829681
//         }
//       }
//     },
//     {
//       "hash": "0xe7e0fe390354509cd08c9a0168536938600ddc552b3f7cb96030ebef62e75895",
//       "sender": {
//         "address": "0x95ba4cf87d6723ad9c0db21737d862be80e93911",
//         "smartContract": {
//           "contractType": null
//         }
//       },
//       "creates": {
//         "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
//         "smartContract": {
//           "currency": {
//             "decimals": 6,
//             "name": "USD//C",
//             "symbol": "USDC",
//             "tokenType": "ERC20"
//           },
//           "contractType": "Token"
//         }
//       },
//       "block": {
//         "timestamp": {
//           "unixtime": 1533324504
//         }
//       }
//     }
//   ]
// }`

var MockHoldersResponse = fmt.Sprintf("{\"data\": %s}", HoldersDataResponse)

const HoldersDataResponse = `{
	"burnt_balance": [
		{
			"balance": 9591.008416
		}
	],
	"high_holders": [
		{
			"average": 181573.10887515705,
			"holders": "344779",
			"median": 2710.7223329999997,
			"standard_deviation": 16042080.190114936
		}
	],
	"holders": [
		{
			"total": "142762.27806229",
			"average": 9957.924793468854,
			"holders": "6362776",
			"median": 14.991373,
			"standard_deviation": 3744829.82386695
		}
	],
	"low_holders": [
		{
			"average": 70.22428187172984,
			"holders": "6029365",
			"median": 11.240175,
			"standard_deviation": 152.36404479261492
		}
	],
	"top_holders": [
		{
				"Holder": {
						"Address": "0x5ee5bf7ae06d1be5997a1a72006fe6c607ec6de8"
				},
				"balance": {
						"Amount": "33824.90611966"
				}
		},
		{
				"Holder": {
						"Address": "0xa3a7b6f88361f48403514059f1f16c8e78d60eec"
				},
				"balance": {
						"Amount": "9472.96070051"
				}
		},
		{
				"Holder": {
						"Address": "0xc3d688b66703497daa19211eedff47f25384cdc3"
				},
				"balance": {
						"Amount": "6856.45687538"
				}
		},
		{
				"Holder": {
						"Address": "0xcbd12525cdd4cd76455859ea3da141c412e54f96"
				},
				"balance": {
						"Amount": "4998.00000000"
				}
		},
		{
				"Holder": {
						"Address": "0xccf4429db6322d5c611ee964527d42e5d685dd6a"
				},
				"balance": {
						"Amount": "3684.56499000"
				}
		},
		{
				"Holder": {
						"Address": "0x40ec5b33f54e0e8a33a975908c5ba1c14e5bbbdf"
				},
				"balance": {
						"Amount": "3574.54581406"
				}
		},
		{
				"Holder": {
						"Address": "0x3ee18b2214aff97000d974cf647e7c347e8fa585"
				},
				"balance": {
						"Amount": "3241.69222955"
				}
		},
		{
				"Holder": {
						"Address": "0x9ff58f4ffb29fa2266ab25e75e2a8b3503311656"
				},
				"balance": {
						"Amount": "2961.43369585"
				}
		},
		{
				"Holder": {
						"Address": "0xbf72da2bd84c5170618fbe5914b0eca9638d5eb5"
				},
				"balance": {
						"Amount": "2616.95610282"
				}
		},
		{
				"Holder": {
						"Address": "0x4197ba364ae6698015ae5c1468f54087602715b2"
				},
				"balance": {
						"Amount": "1908.70816522"
				}
		}
	]
}`

const DataResponse = `{
  "data": {
    "results": [
      {
        "total": 22223,
        "success": 22167,
        "fails": 56
      }
    ]
  }
}`

const TransfersDataResponse = `{
        "data": {
            "transactions": [
                {
                    "fails": "15088",
                    "success": "1483754",
                    "total": "1498842"
                }
            ],
            "transfers": [
                {
                    "fails": "13410",
                    "success": "821810",
                    "total": "835220"
                }
            ]
        }
}`

const BalanceResponse = `{
			"data": {
					"holders": [
							{
									"holders": "104594",
									"supply": "142512.74104314"
							}
					]
			}
}`
const BalanceResponseEmpty = `{
			"data": {
					"holders": []
			}
}`
