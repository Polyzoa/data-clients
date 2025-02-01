package bitquery

const SuccessTransactionsQuery = `query SuccessTransactions($address: String!, $chain: EthereumNetwork!)  {
  data: ethereum(network: $chain) {
    results: transactions(
      txTo: {is: $address}
    ) {
      total: count
      success: count(success: true)
      fails: count(success: false)
    }
  }
}
`

const SuccessTransfersQueryV2 = `query SuccessTransactionsQuery($network: evm_network, $address: String!) {
  data: EVM(dataset: archive, network: $network) {
    results: Calls(
      where: {Call: {To: {is: $address}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}} #eq==0->external
    ) {
      total: count
      fails: count(if: {Call: {Success: false}})
      success: count(if: {Call: {Success: true}})
    }
  }
}`

const QueryTransactionStats = `query ($network: evm_network, $address: String!) {
  data: EVM(dataset: archive, network: $network) {
    transactions: Transactions(where: {Transaction: {To: {is: $address}}}) {
      success: count(if: {TransactionStatus: {Success: true}})
      fails: count(if: {TransactionStatus: {Success: false}})
      total: count
    }
    transfers: Calls(
      where: {Call: {To: {is: $address}, Signature: {Name: {is: "transfer"}}, Depth: {eq: 0}}} #eq==0->external
        ) {
        total: count
        fails: count(if: {Call: {Success: false}})
        success: count(if: {Call: {Success: true}})
        }
    }

}`
