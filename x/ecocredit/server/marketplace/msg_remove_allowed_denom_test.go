package marketplace

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type removeAllowedDenomSuite struct {
	*baseSuite
	err error
}

func TestRemoveAllowedDenom(t *testing.T) {
	gocuke.NewRunner(t, &removeAllowedDenomSuite{}).Path("./features/msg_remove_allowed_denom.feature").Run()
}

func (s *removeAllowedDenomSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 1)
}

func (s *removeAllowedDenomSuite) AnAllowedDenomWithProperties(a gocuke.DocString) {
	var msg *marketplace.AllowedDenom

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	err = s.k.stateStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    msg.BankDenom,
		DisplayDenom: msg.DisplayDenom,
		Exponent:     msg.Exponent,
	})
	require.NoError(s.t, err)
}

func (s *removeAllowedDenomSuite) AliceAttemptsToRemoveABankDenomWithProperties(a gocuke.DocString) {
	var msg *marketplace.MsgRemoveAllowedDenom

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.RemoveAllowedDenom(s.ctx, msg)
}

func (s *removeAllowedDenomSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *removeAllowedDenomSuite) ExpectBankDenomIsRemoved(denom string) {
	_, err := s.marketStore.AllowedDenomTable().Get(s.ctx, denom)
	require.Error(s.t, err)
	require.ErrorIs(s.t, err, ormerrors.NotFound)
}

func (s *removeAllowedDenomSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *removeAllowedDenomSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
