package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"revup/x/revup/types"
)

func (k Keeper) AppendReview(ctx sdk.Context, review types.Review) uint64 {
	// Retrieve the current review count
	count := k.GetReviewCount(ctx)

	// Set the review ID to the current count
	review.Id = count

	// Access the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ReviewKey))

	// Convert the review ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, review.Id)

	// Marshal the review into bytes
	appendedValue := k.cdc.MustMarshal(&review)

	// Insert the review bytes using the review ID as a key
	store.Set(byteKey, appendedValue)

	// Update the review count
	k.SetReviewCount(ctx, count+1)
	return count
}

func (k Keeper) GetReviewCount(ctx sdk.Context) uint64 {
	// Access the store using storeKey and ReviewCountKey
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ReviewCountKey))

	// Convert the ReviewCountKey to bytes
	byteKey := []byte(types.ReviewCountKey)

	// Retrieve the value of the count
	bz := store.Get(byteKey)

	// If the count value is not found (e.g., it's the first review), return 0
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetReviewCount(ctx sdk.Context, count uint64) {
	// Access the store using storeKey and ReviewCountKey
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ReviewCountKey))

	// Convert the ReviewCountKey to bytes
	byteKey := []byte(types.ReviewCountKey)

	// Convert the count from uint64 to bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of the count key to the updated count
	store.Set(byteKey, bz)
}
