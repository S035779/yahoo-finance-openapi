package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/pasdam/yahoo-finance-openapi/gen/go/yq1"
	"github.com/stretchr/testify/assert"
)

func Test_ParseQuote(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "quote_valid.json"))
	assert.NoError(t, err)

	result := &yq1.QuoteResponse{}
	err = json.Unmarshal(data, result)
	assert.NoError(t, err)

	assert.Equal(t, &yq1.QuoteResponse{
		QuoteResponse: &yq1.QuoteResponseQuoteResponse{
			Result: &[]yq1.QuoteResult{
				{
					Symbol: strPtr("TSLA"),
					FiftyTwoWeekLowChangePercent: &yq1.QuoteResultFiftyTwoWeekLowChangePercent{
						Raw: fPtr(2.3255618),
						Fmt: strPtr("232.56%"),
					},
					Language: strPtr("en-US"),
					RegularMarketDayRange: &yq1.QuoteResultRegularMarketDayRange{
						Raw: strPtr("611.8 - 628.3472"),
						Fmt: strPtr("611.80 - 628.35"),
					},
					RegularMarketDayHigh: &yq1.QuoteResultRegularMarketDayHigh{
						Raw: fPtr(628.3472),
						Fmt: strPtr("628.35"),
					},
					FiftyTwoWeekHighChange: &yq1.QuoteResultFiftyTwoWeekHighChange{
						Raw: fPtr(-277.09003),
						Fmt: strPtr("-277.09"),
					},
					FiftyTwoWeekRange: &yq1.QuoteResultFiftyTwoWeekRange{
						Raw: strPtr("187.43 - 900.4"),
						Fmt: strPtr("187.43 - 900.40"),
					},
					FirstTradeDateMilliseconds: i64Ptr(1277818200000),
					ExchangeDataDelayedBy:      i32Ptr(0),
					FiftyTwoWeekLow: &yq1.QuoteResultFiftyTwoWeekLow{
						Raw: fPtr(187.43),
						Fmt: strPtr("187.43"),
					},
					Market: strPtr("us_market"),
					RegularMarketVolume: &yq1.QuoteResultRegularMarketVolume{
						Raw:     i64Ptr(23853471),
						Fmt:     strPtr("23.853M"),
						LongFmt: strPtr("23,853,471"),
					},
					QuoteSourceName: strPtr("Nasdaq Real Time Price"),
					MessageBoardId:  strPtr("finmb_27444752"),
					PriceHint:       i32Ptr(2),
					SourceInterval:  i32Ptr(15),
					RegularMarketDayLow: &yq1.QuoteResultRegularMarketDayLow{
						Raw: fPtr(611.8),
						Fmt: strPtr("611.80"),
					},
					Exchange:              strPtr("NMS"),
					ShortName:             strPtr("Tesla, Inc."),
					Region:                strPtr("US"),
					FullExchangeName:      strPtr("NasdaqGS"),
					GmtOffSetMilliseconds: i32Ptr(-14400000),
					RegularMarketOpen: &yq1.QuoteResultRegularMarketOpen{
						Raw: fPtr(613.37),
						Fmt: strPtr("613.37"),
					},
					RegularMarketTime: &yq1.QuoteResultRegularMarketTime{
						Raw: i32Ptr(1624046404),
						Fmt: strPtr("4:00PM EDT"),
					},
					RegularMarketChangePercent: &yq1.QuoteResultRegularMarketChangePercent{
						Raw: fPtr(1.0882294),
						Fmt: strPtr("1.09%"),
					},
					QuoteType: strPtr("EQUITY"),
					FiftyTwoWeekLowChange: &yq1.QuoteResultFiftyTwoWeekLowChange{
						Raw: fPtr(435.88),
						Fmt: strPtr("435.88"),
					},
					FiftyTwoWeekHighChangePercent: &yq1.QuoteResultFiftyTwoWeekHighChangePercent{
						Raw: fPtr(-0.30774102),
						Fmt: strPtr("-30.77%"),
					},
					Tradeable: bPtr(false),
					Currency:  strPtr("USD"),
					RegularMarketPreviousClose: &yq1.QuoteResultRegularMarketPreviousClose{
						Raw: fPtr(616.6),
						Fmt: strPtr("616.60"),
					},
					FiftyTwoWeekHigh: &yq1.QuoteResultRegularMarketDayHigh{
						Raw: fPtr(900.4),
						Fmt: strPtr("900.40"),
					},
					ExchangeTimezoneName: strPtr("America/New_York"),
					RegularMarketChange: &yq1.QuoteResultRegularMarketChange{
						Raw: fPtr(6.710022),
						Fmt: strPtr("6.71"),
					},
					ExchangeTimezoneShortName: strPtr("EDT"),
					RegularMarketPrice: &yq1.QuoteResultRegularMarketPrice{
						Raw: fPtr(623.31),
						Fmt: strPtr("623.31"),
					},
					MarketState: strPtr("CLOSED"),
					Triggerable: bPtr(true),
				},
			},
		},
	}, result)
}

func bPtr(val bool) *bool {
	return &val
}

func i32Ptr(val int32) *int32 {
	return &val
}

func i64Ptr(val int64) *int64 {
	return &val
}

func fPtr(val float32) *float32 {
	return &val
}

func strPtr(val string) *string {
	return &val
}
