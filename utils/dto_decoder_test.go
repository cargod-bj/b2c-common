package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

func TestDecodeDto(t *testing.T) {
	sec := 1593685382170 / 1e3
	tm := time.Unix(int64(sec), 0)
	ts := timestamp.Timestamp{Seconds: int64(1593685382)}
	d, _ := decimal.NewFromString("123.123")
	m := InModel{
		1593685382170, 1593685382170, 1593685382170, 1593685382170,
		//tm,
		&tm,
		//tm,
		&tm,
		//tm,
		&tm,
		//tm,
		&tm,
		//ts,
		&ts,
		ts,
		&ts,
		&d,
		"123.123",
	}
	o := OutModel{}
	err := DecodeDto(m, &o)
	print(err)

	t.Fatal("测试失败")
}

type InModel struct {
	PUint642Time    uint64
	PUint642TimePtr uint64
	PInt642Time     int64
	PInt642TimePtr  int64
	//PTime2U64             time.Time
	PTimePtr2U64 *time.Time
	//PTime2I64             time.Time
	PTimePtr2I64 *time.Time
	//PTime2Timestamp       time.Time
	PTimePtr2Timestamp *time.Time
	//PTime2TimestampPtr    time.Time
	PTimePtr2TimestampPtr *time.Time
	//PTimestamp2Time       timestamp.Timestamp
	PTimestampPtr2Time    *timestamp.Timestamp
	PTimestamp2TimePtr    timestamp.Timestamp
	PTimestampPtr2TimePtr *timestamp.Timestamp
	PDecimalPtr2String    *decimal.Decimal
	PString2DecimalPtr    string
}

type OutModel struct {
	PUint642Time    time.Time
	PUint642TimePtr *time.Time
	PInt642Time     time.Time
	PInt642TimePtr  *time.Time
	//PTime2U64             uint64
	PTimePtr2U64 uint64
	//PTime2I64             int64
	PTimePtr2I64 int64
	//PTime2Timestamp       timestamp.Timestamp
	PTimePtr2Timestamp timestamp.Timestamp
	//PTime2TimestampPtr    *timestamp.Timestamp
	PTimePtr2TimestampPtr *timestamp.Timestamp
	//PTimestamp2Time       time.Time
	PTimestampPtr2Time time.Time
	//PTimestamp2TimePtr    *time.Time
	PTimestampPtr2TimePtr *time.Time
	PDecimalPtr2String    *string
	PString2DecimalPtr    *decimal.Decimal
}
