package market_data

import (
	"context"
	"testing"

	"github.com/Ch2r1Ie/Stock-Bubble/app/market_data/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type TestServiceSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	ctx    context.Context
	logger *zap.Logger

	mYahooFinance *mock.MockAPI
	srv           marketSrv
}

func (s *TestServiceSuite) SetupTest() {

	s.ctrl = gomock.NewController(s.T())
	s.ctx = context.Background()
	s.logger = zap.NewNop()

	s.mYahooFinance = mock.NewMockAPI(s.ctrl)
	s.srv = NewMarketService(s.mYahooFinance)

}

func TestService(t *testing.T) {
	suite.Run(t, new(TestServiceSuite))
}

func (t *TestServiceSuite) TearDownTest() {
	t.ctrl.Finish()
}
