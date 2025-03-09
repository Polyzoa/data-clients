package bitquery

import (
	"github.com/massigerardi/graphql"
)

var StatsQuery = Query{
	Query: `query TransactionsQuery($network: evm_network, $addresses: [String!]) {
  data: EVM(dataset: archive, network: $network) {
    transfer_stats: Transfers(where: {Transfer: {Currency: {SmartContract: {in: $addresses}}}}) {
      receivers: count(distinct: Transfer_Receiver)
      average(of: Transfer_Amount)
      standard_deviation(of: Transfer_Amount)
      median(of: Transfer_Amount)
      token: Transfer {
        Currency {
          address: SmartContract
        }
      }
    }
		transfers: Calls(
      where: {Call: {To: {in: $addresses}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}}
    ) {
      token: Transaction {
        address: To
      }
      total: count
      fails: count(if: {Call: {Success: false}})
      success: count(if: {Call: {Success: true}})
    }
    senders: Calls(
      where: {Call: {To: {in: $addresses}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}}
      orderBy: {ascendingByField: "success"}
    ) {
      token: Transaction {
        address: To
      }
      total: count
      fails: count(if: {Call: {Success: false}})
      success: count(if: {Call: {Success: true}})
    }
    transactions: Transactions(where: {Transaction: {To: {in: $addresses}}}) {
      token: Transaction {
        address: To
      }
      total: count
      success: count(if: {Transaction: {}, TransactionStatus: {Success: true}})
      fails: count(if: {Transaction: {}, TransactionStatus: {Success: false}})
    }
  }
}
`,
	Params: nil,
	Client: graphql.NewClient(string(EndpointV2)),
}

var StatsQueryInTime = Query{
	Query: `query TransactionsQuery($network: evm_network, $addresses: [String!],
			$after: DateTime, $before: DateTime
	) {
  data: EVM(dataset: archive, network: $network) {
    transfer_stats: Transfers(where: {
			Transfer: {Currency: {SmartContract: {in: $addresses}}},
			Block: {Time: {before: $before, after: $after}}}
		) {
      receivers: count(distinct: Transfer_Receiver)
      average(of: Transfer_Amount)
      standard_deviation(of: Transfer_Amount)
      median(of: Transfer_Amount)
			total: sum(of: Transfer_Amount)
      token: Transfer {
        Currency {
          address: SmartContract
        }
      }
    }
		transfers: Calls(
      where: {Call: {To: {in: $addresses}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}}
    ) {
      token: Transaction {
        address: To
      }
      total: count
      fails: count(if: {Call: {Success: false}})
      success: count(if: {Call: {Success: true}})
    }
    senders: Calls(
      where: {Call: {To: {in: $addresses}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}}
      orderBy: {ascendingByField: "success"}
    ) {
      token: Transaction {
        address: To
      }
      total: count
      fails: count(if: {Call: {Success: false}})
      success: count(if: {Call: {Success: true}})
    }
    transactions: Transactions(where: {Transaction: {To: {in: $addresses}}}) {
      token: Transaction {
        address: To
      }
      total: count
      success: count(if: {Transaction: {}, TransactionStatus: {Success: true}})
      fails: count(if: {Transaction: {}, TransactionStatus: {Success: false}})
    }
  }
`,
	Params: nil,
	Client: graphql.NewClient(string(EndpointV2)),
}
var ContractQuery = Query{
	Query: `
query ContractQuery($addresses: [String!] = "", $chain: EthereumNetwork!) {
  data: ethereum(network: $chain) {
    transactions(txCreates: {in: $addresses}) {
			hash
      sender {
        address
				smartContract {
          contractType
        }
      }
      creates {
        address
        smartContract {
          currency {
            decimals
            name
            symbol
            tokenType
          }
          contractType
        }
      }
			block {
        timestamp {
          unixtime
        }
      }
    }
  }
`,
	Params: nil,
	Client: graphql.NewClient(string(EndpointV2)),
}

var BalanceQuery = Query{
	Query: `query Holders($address: String!, $network: evm_network, $date: String!) {
    data: EVM(dataset: archive, network: $network) {
        holders: TokenHolders(
            date: $date
            where: {Balance: {Amount: {gt: "0"}}, Holder: {Address: {notIn: ["0x000000000000000000000000000000000000dead", "0x0000000000000000000000000000000000000000"]}}}
            tokenSmartContract: $address
        ) {
						supply: sum(of: Balance_Amount)
            holders: uniq(of: Holder_Address)
        }
	 }
}`,
	Params: nil,
	Client: graphql.NewClient(string(EndpointV2)),
}

var HoldersQuery = Query{
	Query: `query Holders($address: String!, $network: evm_network, $date: String!, $amount: String!) {
    data: EVM(dataset: archive, network: $network) {
        holders: TokenHolders(
            date: $date
            where: {Balance: {Amount: {gt: "0"}}}
            tokenSmartContract: $address
        ) {
						total: sum(of: Balance_Amount)
            average(of: Balance_Amount)
            holders: uniq(of: Holder_Address)
            standard_deviation(of: Balance_Amount)
            median(of: Balance_Amount)
        }
        high_holders: TokenHolders(
            date: $date
            where: {Balance: {Amount: {ge: $amount}}, Holder: {Address: {notIn: ["0x000000000000000000000000000000000000dead", "0x0000000000000000000000000000000000000000"]}}}
            tokenSmartContract: $address
        ) {
            total: sum(of: Balance_Amount)
						average(of: Balance_Amount)
            holders: uniq(of: Holder_Address)
            standard_deviation(of: Balance_Amount)
            median(of: Balance_Amount)
        }
        burnt_balance: TokenHolders(
            date: $date
            where: {Holder: {Address: {in: ["0x000000000000000000000000000000000000dead", "0x0000000000000000000000000000000000000000"]}}}
            tokenSmartContract: $address
        ) {
            balance: average(of: Balance_Amount)
        }
        low_holders: TokenHolders(
            date: $date
            where: {Balance: {Amount: {le: $amount, gt: "0"}}, , Holder: {Address: {notIn: ["0x000000000000000000000000000000000000dead", "0x0000000000000000000000000000000000000000"]}}}
            tokenSmartContract: $address
        ) {
            total: sum(of: Balance_Amount)
            average(of: Balance_Amount)
            holders: uniq(of: Holder_Address)
            standard_deviation(of: Balance_Amount)
            median(of: Balance_Amount)
        }
        top_holders: TokenHolders(
                date: $date
                limit: {count: 10}
                orderBy: {descending: Balance_Amount}
                tokenSmartContract: $address
            ) {
                Holder {
                    Address
                }
                balance: Balance {
                    Amount
            }
        }
    }
}`,
	Params: nil,
	Client: graphql.NewClient(string(EndpointV2)),
}
