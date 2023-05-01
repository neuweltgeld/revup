package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"revup/x/revup/types"
)

func (k msgServer) CreateReview(goCtx context.Context, msg *types.MsgCreateReview) (*types.MsgCreateReviewResponse, error) {
	// Unwrap the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a review instance with values from the message
	var review = types.Review{
		Creator: msg.Creator, // Assign the creator from the message
		Product: msg.Product, // Assign the product from the message
		Comment: msg.Comment, // Assign the comment from the message
		Rate:    msg.Rate,    // Assign the rate from the message
	}

	// Add a review to the store and get back the ID
	id := k.AppendReview(ctx, review)

	// Return the ID of the newly created review
	return &types.MsgCreateReviewResponse{Id: id}, nil
}
