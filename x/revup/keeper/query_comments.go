package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"revup/x/revup/types"
)

func (k Keeper) Comments(c context.Context, req *types.QueryCommentsRequest) (*types.QueryCommentsResponse, error) {
	// Return an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable to store a list of reviews
	var comments []*types.Review

	// Unwrap the context to obtain information about the environment
	ctx := sdk.UnwrapSDKContext(c)

	// Get the key-value module store using the store key ("chain" in this case)
	store := ctx.KVStore(k.storeKey)

	// Access the part of the store that stores reviews (using the ReviewKey, which is "Review-value-")
	reviewStore := prefix.NewStore(store, []byte(types.ReviewKey))

	// Paginate the review store based on the PageRequest
	pageRes, err := query.Paginate(reviewStore, req.Pagination, func(key []byte, value []byte) error {
		var review types.Review
		if err := k.cdc.Unmarshal(value, &review); err != nil {
			return err
		}

		comments = append(comments, &review)

		return nil
	})

	// Return an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of reviews and pagination information
	return &types.QueryCommentsResponse{Review: comments, Pagination: pageRes}, nil
}
