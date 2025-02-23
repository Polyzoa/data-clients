package bitquery

var SolanaTransferQuery = Query{
	Query: `query TransferQuery($address: String!) {
  data: Solana(network: solana) {
    Transfers(
      where: {Transfer: {Currency: {MintAddress: {is: $address}}}}
    ) {
      count
      success:count(if: {Transaction: {Result: {Success: true}}})
      fail:count(if: {Transaction: {Result: {Success: false}}})
    }
  }
}`,
	Params: nil,
	Url:    EndpointV3,
}

var SolanaTopHoldersQuery = Query{
	Query: `query TopHolders($address: String!, $date: String) {
  Solana {
    BalanceUpdates(
      orderBy: {descendingByField: "BalanceUpdate_Holding_maximum"}
      where: {BalanceUpdate: {Currency: {MintAddress: {is: $address}}}, Transaction: {Result: {Success: true}}, Block: {Date: {is: $date}}}
      limit: {count: 10}
    ) {
      BalanceUpdate {
        Account {
          Address
          Owner
        }
        Holding: PostBalance(maximum: Block_Slot)
      }
    }
  }
}`,
	Params: nil,
	Url:    EndpointV3,
}

const SolanaSupplyQuery = `query SupplyQuery($address: String!) {
  Solana(network: solana) {
    BalanceUpdates(
      where: {BalanceUpdate: {Account: {}, Currency: {MintAddress: {is: $address}}, Amount: {gt: "0"}}}
    ) {
      supply: sum(of: BalanceUpdate_PostBalance)
      average(of: BalanceUpdate_PostBalance)
      standard_deviation(of: BalanceUpdate_PostBalance)
    }
  }
}
`

var SolanaTradersQuery = Query{
	Query: `query SolanaTradersQuery($address: String!) {
  Solana {
    DEXTrades(
      where: {Trade: {Buy: {Currency: {MintAddress: {is: $address}}}}}
      limit: {count: 100}
      orderBy: {ascending: Block_Time}
    ) {
      Trade {
        Buy {
          Amount
          Account {
            Address
            Owner
            Token {
              Owner
            }
          }
        }
      }
    }
  }
}`,
	Params: nil,
	Url:    EndpointV3,
}
