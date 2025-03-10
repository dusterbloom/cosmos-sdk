package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestTokensToConsensusPower() {
	suite.Require().Equal(int64(0), suite.app.StakingKeeper.TokensToConsensusPower(suite.ctx, suite.app.StakingKeeper.Config.PowerReduction.Sub(sdk.NewInt(1))))
	suite.Require().Equal(int64(1), suite.app.StakingKeeper.TokensToConsensusPower(suite.ctx, suite.app.StakingKeeper.Config.PowerReduction))
}

func (suite *KeeperTestSuite) TestTokensFromConsensusPower() {
	suite.Require().Equal(sdk.NewInt(0), suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 0))
	suite.Require().Equal(suite.app.StakingKeeper.Config.PowerReduction, suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 1))
}
