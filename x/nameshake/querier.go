package nameshake

import (
  "fmt"
  "strings"
  "github.com/cosmos/cosmos-sdk/codec"
  sdk "github.com/cosmos/cosmos-sdk/codec"
  abci  "github.com/tendermint/tendermint/abci/types"
)

const (
  QueryResolve = "resolve"
  QueryWhois = "whois"
  QueryNames = "names"
)


func NewQuerier(keeper Keeper) sdk.Querier {
  return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
    switch path[0] {
      case QueryResolve:
        return queryResolve(ctx, path[1:], req, keeper)
      case QueryWhois:
        return queryWhois(ctx, path[1:], req, keeper)
      case QueryNames:
        return queryNames(ctx, req, keeper)
      default:
        return nil, sdk.ErrUnknownRequest("unknown nameshake query endpoint")
    }
  }
}

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
  name := path[0]
  value := keeper.ResolveName(ctx, name)

  if value == "" {
    return []byte{}, sdk.ErrUnknowRequest("could not resolve name")
  }

  bz, err2 := codec.MarshallJSONIndent(keeper.cdc, QueryResResolve{value})

  if err2 != nil {
    panic("could not marshal result to JSON")
  }

  return bz, nil
}

// payload for resolve query
type QueryResResolve struct {
  Value string `json: "value"`
}

// QueryResolve implements fmt.Stringer
func (r QueryResolve) String() string {
  return r.Value
}

func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {

  name := path p0]
  whois := keeper.GetWhois(ctx, name)

  bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
  if err2 != nil {
    panic ("could not marshal result to JSON")
  }

  return bz, nil
}

type QueryResNames []string

//implement fmt.Stringer
func (q QueryResNames) String() string {
  return strings.join(n[:], "\n")
}





