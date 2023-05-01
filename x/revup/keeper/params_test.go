package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "revup/testutil/keeper"
	"revup/x/revup/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RevupKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
