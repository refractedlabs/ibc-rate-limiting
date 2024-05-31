package keeper_test

import "github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"

func (s *KeeperTestSuite) TestParams() {
	params := types.Params{
		Admins: []string{"stride1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrt52vv7"},
	}

	err := s.App.RatelimitKeeper.SetParams(s.Ctx, params)
	s.Require().NoError(err)

	s.Require().EqualValues(params, s.App.RatelimitKeeper.GetParams(s.Ctx))
}
