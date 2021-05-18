package rpc

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"
)

type (
	VoteAccountKey struct {
		NodePubkey       string `json:"nodePubkey"`
		VotePubkey       string `json:"votePubkey"`
		validatorBalance int    `json:"lastVote"`
	}

	GetVoteAccountsBalanceResponse struct {
		Result struct {
			Current []VoteAccountKey `json:"current"`
		} `json:"result"`
		Error rpcError `json:"error"`
	}
)

// https://docs.solana.com/developing/clients/jsonrpc-api#getvoteaccounts
func (c *RPCClient) GetVoteAccountsBalance(ctx context.Context, commitment Commitment) (*GetVoteAccountsBalanceResponse, error) {
	body, err := c.rpcRequest(ctx, formatRPCRequest("getVoteAccounts", []interface{}{commitment}))
	if err != nil {
		return nil, fmt.Errorf("RPC call failed: %w", err)
	}

	klog.V(3).Infof("getVoteAccounts response: %v", string(body))
	// for i, ch := range GetVoteAccountsBalanceResponse.Result {
	// 	fmt.println(ch)
	// 	GetVoteAccountsBalanceResponse[i].validatorBalance := '1'
	// }

	var resp GetVoteAccountsBalanceResponse
	// fmt.Print(resp.Result.Current)
	// for i, ch := range resp.Result.Current {
	// 	fmt.println(ch)
	// 	resp.Result.Current[i].validatorBalance := 1
	// }

	// if err = json.Unmarshal(body, &resp); err != nil {
	// 	return nil, fmt.Errorf("failed to decode response body: %w", err)
	// }

	// if resp.Error.Code != 0 {
	// 	return nil, fmt.Errorf("RPC error: %d %v", resp.Error.Code, resp.Error.Message)
	// }

	return &resp, nil
}
