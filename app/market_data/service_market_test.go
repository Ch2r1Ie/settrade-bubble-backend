package market_data

import (
	"context"
	"errors"
	"testing"

	"github.com/Ch2r1Ie/Stock-Bubble/app/market_data/mock"
	"github.com/Ch2r1Ie/Stock-Bubble/yahoo_finance.go"
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

func (s *TestServiceSuite) Test_Service_happyCase() {

	symbols := []string{"PTTGC", "GPSC"}
	daterange := yahoo_finance.RangeOneYear
	interval := yahoo_finance.IntervalOneDay

	mockPTTGC := yahoo_finance.Stock{
		Meta: yahoo_finance.Meta{
			Currency:             "THB",
			Symbol:               "PTTGC.BK",
			ExchangeName:         "SET",
			ExchangeTimezoneName: "Asia/Bangkok",
			InstrumentType:       "EQUITY",
		},
		Timestamp: []int64{1233455},
		Indicators: yahoo_finance.Indicators{
			Quote: []yahoo_finance.Quote{
				{
					Low:    []float64{24.123},
					High:   []float64{24.222},
					Volume: []float64{123423},
					Open:   []float64{24.155},
					Close:  []float64{24.200},
				},
			},
		},
	}

	mockGPSC := yahoo_finance.Stock{
		Meta: yahoo_finance.Meta{
			Currency:             "THB",
			Symbol:               "GPSC.BK",
			ExchangeName:         "SET",
			ExchangeTimezoneName: "Asia/Bangkok",
			InstrumentType:       "EQUITY",
		},
		Timestamp: []int64{1233455},
		Indicators: yahoo_finance.Indicators{
			Quote: []yahoo_finance.Quote{
				{
					Low:    []float64{25.123},
					High:   []float64{25.222},
					Volume: []float64{234123},
					Open:   []float64{25.155},
					Close:  []float64{25.200},
				},
			},
		},
	}

	s.mYahooFinance.EXPECT().Info("PTTGC.BK", daterange, interval).Times(1).Return(mockPTTGC, nil)
	s.mYahooFinance.EXPECT().Info("GPSC.BK", daterange, interval).Times(1).Return(mockGPSC, nil)

	infos, err := s.srv.stocks(symbols, daterange, interval)

	s.NoError(err)
	s.NotNil(infos)
	s.Equal(2, len(*infos))

	expectedInfos := []yahoo_finance.Info{
		format_response(mockPTTGC),
		format_response(mockGPSC),
	}

	for i, expectedInfo := range expectedInfos {
		s.Equal(expectedInfo, (*infos)[i])
	}
}

func (s *TestServiceSuite) Test_Service_failCase() {

	symbols := []string{"PTTGC", "GPSC"}
	daterange := yahoo_finance.RangeOneYear
	interval := yahoo_finance.IntervalOneDay

	mockPTTGC := yahoo_finance.Stock{
		Meta: yahoo_finance.Meta{
			Currency:             "THB",
			Symbol:               "PTTGC.BK",
			ExchangeName:         "SET",
			ExchangeTimezoneName: "Asia/Bangkok",
			InstrumentType:       "EQUITY",
		},
		Timestamp: []int64{1233455},
		Indicators: yahoo_finance.Indicators{
			Quote: []yahoo_finance.Quote{
				{
					Low:    []float64{24.123},
					High:   []float64{24.222},
					Volume: []float64{123423},
					Open:   []float64{24.155},
					Close:  []float64{24.200},
				},
			},
		},
	}

	s.mYahooFinance.EXPECT().Info("PTTGC.BK", daterange, interval).Times(1).Return(mockPTTGC, errors.New("unknown errors"))

	infos, err := s.srv.stocks(symbols, daterange, interval)

	s.Equal(err, errors.New("unknown errors"))
	s.Nil(infos)
}
