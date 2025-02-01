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
