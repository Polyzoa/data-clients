//go:build !mock

package bitquery

import "github.com/massigerardi/graphql"

func getGraphqlClient(url string) GraphqlClient {
	return graphql.NewClient(url)
}
