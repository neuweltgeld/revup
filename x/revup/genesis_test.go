package revup_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "revup/testutil/keeper"
	"revup/testutil/nullify"
	"revup/x/revup"
	"revup/x/revup/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RevupKeeper(t)
	revup.InitGenesis(ctx, *k, genesisState)
	got := revup.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
