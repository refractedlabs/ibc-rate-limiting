package types_test

import (
	"github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParamsValidation(t *testing.T) {
	params := types.DefaultParams()
	require.NoError(t, params.Validate())

	params = types.Params{Admins: []string{}}
	require.NoError(t, params.Validate())

	params = types.Params{Admins: []string{"stride1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrt52vv7"}}
	require.NoError(t, params.Validate())

	params = types.Params{Admins: []string{"invalid_address"}}
	require.Error(t, params.Validate())
}
