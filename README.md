# Data Providers

## Data Clients

### Bitquery

#### Define Query

```go
package bitquery

type Query struct {
    Url    Endpoint
    Query  string
    Params map[string]any
}
```

with:

* `Url`: the endpoint on which run the query. it can be:
  * EndpointV1 = "https://graphql.bitquery.io"
  * EndpointV2 = "https://streaming.bitquery.io/graphql"
  * EndpointV3 = "https://streaming.bitquery.io/eap"

* `Query`: the query to run
* `Params`: a map with the parameters accepted by the query

To run the Query, call `client.RunQuery(query, response)` with `response` the expected response data.

**Example**

```go
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
	Params: [string]any{
      "address": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
    },
	Url:    EndpointV3,
}

client = NewClient(apiKey, apiToken)
var response Response[any]
err := client.RunQuery(SolanaTransferQuery, &response)

```
