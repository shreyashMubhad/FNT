package models

import "time"

type InternalHeader struct {
	TransactionCode uint16
	LogTime         uint32
	AlphaChar       [2]byte
	UserId          uint32
	ErrorCode       uint16
	TimeStamp       [8]byte
	TimeStamp1      [8]byte
	TimeStamp2      [8]byte
	MessageLength   uint16
}

// heartbeat-------------------------------------------------------------------------------
type ExternalHeartbeatHeader struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  Heartbeat
}

type Heartbeat struct {
	MessageHeader InternalHeader
}

// logon request---------------------------------------------------------------------------
type ExternalLogonInHeader struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  LogonIN
}

type LogonIN struct {
	MessageHeader InternalHeader
	UserId        uint32
	Reserved      [8]byte
	Password      [12]byte
	Reserved1     [4]byte
	Reserved2     [38]byte
	BrokerId      [5]byte
	Reserved3     [119]byte
	Reserved4     [16]byte
	Reserved5     [16]byte
	Reserved6     [16]byte
}

//logon	responce--------------------------------------------------------------------------

type ExternalLogonOutHeader struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  LogonOut
}

type LogonOut struct {
	MessageHeader InternalHeader
	UserId        uint32
	Reserved      [8]byte
	Password      [12]byte
	Reserved1     [4]byte
	Reserved2     [38]byte
	BrokerId      [5]byte
	Reserved3     [119]byte
	Reserved4     [16]byte
	Reserved5     [16]byte
	Reserved6     [16]byte
}

// logon error------------------------------------------------------------------------------
type ExternalLogonOutError struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  ErrorResponce
}

type ErrorResponce struct {
	MessageHeader InternalHeader
	Reserves      [14]byte
	ErrorMessage  [128]byte
}

// dc download request----------------------------------------------------------------------
type ExternalDcDownloadHeader struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  DcDownloadRequest
}

type DcDownloadRequest struct {
	MessageHeader InternalHeader
	SeqNum        float64
}

// trade confirmation-----------------------------------------------------------------------
type ExternalTradeConfirmation struct {
	BodyLen  uint16
	SeqNum   uint32
	CheckSum [16]byte
	MsgData  TradeConfirmation
}

type ContrectInfo struct {
	InstrumentName [6]byte
	Symbol         [10]byte
	ExpiryDate     int32
	StrikePrice    int32
	OptionType     [2]byte
	CALevel        [2]byte
}

type TradeConfirmation struct {
	MessageHeader         InternalHeader
	ResponseOrderNumber   int64
	BrokerId              [5]byte
	Reserved              byte
	TradeNum              int32
	AccountNum            [10]byte
	BuySell               [2]byte
	OriginalVol           int32
	DisclosedVol          int32
	RemainingVol          int32
	DisclosedVolRemaining int32
	Price                 int64
	StOrderFlags          int16 // this is an struct, for more refer the doc
	Gtd                   int32
	FillNumber            int32
	FillQty               int32
	FillPrice             int32
	VolFilledToday        int32
	ActivityType          [2]byte
	ActivityTime          int32
	OpOrderNum            float64
	OpBrokerId            [5]byte
	Token                 int32
	ContrectDesc          ContrectInfo
	OpenClose             byte
	OldOpenClose          byte
	BookType              byte
	NewVolume             int32
	OldAccountNumber      [10]byte
	Participant           [12]byte
	OldParticipant        [12]byte
	AdditionalOrderFlags  int16 // this is an struct, for more refer the doc
	ReservedFiller        byte
	GiveUpTrade           byte
	ReservedFiller2       byte
	PAN                   [10]byte
	OldPan                [10]byte
	AlgoId                int32
	AlgoCategory          [2]byte
	LastActivityReference int64
	NnfField              float64
}

// ------------------------------
func (T *TradeConfirmation) FnsetResOrderNum(resOrdNum int64) {
	T.ResponseOrderNumber = resOrdNum
}

func (T *TradeConfirmation) FngetResOrderNum() int64 {
	return T.ResponseOrderNumber
}

func (ON *LogonIN) GetType() int                        { return 0 }
func (ERR *ExternalLogonOutError) GetType() int         { return 0 }
func (TRADE *ExternalTradeConfirmation) GetType() int   { return 0 }
func (OUT *ExternalLogonOutHeader) GetType() int        { return 0 }
func (DCEX *ExternalDcDownloadHeader) GetType() int     { return 0 }
func (HEARTBEAT *ExternalHeartbeatHeader) GetType() int { return 0 }
func (DC *DcDownloadRequest) GetType() int              { return 0 }

//---------------------------------------------------------

// type ExtractTrade struct {
// 	ResponseOrderNumber   int64
// 	BrokerId              string
// 	TradeNum              int32
// 	AccountNum            string
// 	BuySell               string
// 	OriginalVol           int32
// 	DisclosedVol          int32
// 	RemainingVol          int32
// 	DisclosedVolRemaining int32
// 	Price                 int64
// 	StOrderFlags          int16
// 	Gtd                   int32
// 	FillNumber            int32
// 	FillQty               int32
// 	FillPrice             int64
// 	VolFilledToday        int32
// 	ActivityType          int16
// 	ActivityTime          int32
// 	OpOrderNum            float64
// 	OpBrokerId            string
// 	Token                 int32
// 	InstrumentName        string
// 	Symbol                string
// 	ExpiryDate            time.Time //need changes
// 	StrikePrice           int32
// 	OptionType            string
// 	CALevel               string
// 	OpenClose             string
// 	OldOpenClose          string
// 	AdditionalOrderFlags  string
// 	GiveUpTrade           string
// 	PAN                   string
// 	OldPan                string
// 	AlgoId                int32
// 	AlgoCategory          string
// 	LastActivityReference string
// 	NnfField              string
// }

type ExtractTrade struct {
	ResponseOrderNumber   int64     `gorm:"column:response_order_number"` 
	BrokerId              string    `gorm:"column:broker_id"`
	TradeNum              int32     `gorm:"column:trade_num"`
	AccountNum            string    `gorm:"column:account_num"`
	BuySell               string    `gorm:"column:buy_sell"`
	OriginalVol           int32     `gorm:"column:original_vol"`
	DisclosedVol          int32     `gorm:"column:disclosed_vol"`
	RemainingVol          int32     `gorm:"column:remaining_vol"`
	DisclosedVolRemaining int32     `gorm:"column:disclosed_vol_remaining"`
	Price                 int64     `gorm:"column:price"`
	StOrderFlags          int16     `gorm:"column:st_order_flags"`
	Gtd                   int32     `gorm:"column:gtd"`
	FillNumber            int32     `gorm:"column:fill_number"`
	FillQty               int32     `gorm:"column:fill_qty"`
	FillPrice             int64     `gorm:"column:fill_price"`
	VolFilledToday        int32     `gorm:"column:vol_filled_today"`
	ActivityType          string     `gorm:"column:activity_type"`
	ActivityTime          int32     `gorm:"column:activity_time"`
	OpOrderNum            float64   `gorm:"column:op_order_num"`
	OpBrokerId            string    `gorm:"column:op_broker_id"`
	Token                 int32     `gorm:"column:token"`
	InstrumentName        string    `gorm:"column:instrument_name"`
	Symbol                string    `gorm:"column:symbol"`
	ExpiryDate            time.Time `gorm:"column:expiry_date"` 
	StrikePrice           int32     `gorm:"column:strike_price"`
	OptionType            string    `gorm:"column:option_type"`
	CALevel               string    `gorm:"column:ca_level"`
	OpenClose             string    `gorm:"column:open_close"`
	OldOpenClose          string    `gorm:"column:old_open_close"`
	AdditionalOrderFlags  string    `gorm:"column:additional_order_flags"`
	GiveUpTrade           string    `gorm:"column:give_up_trade"`
	PAN                   string    `gorm:"column:pan"`
	OldPan                string    `gorm:"column:old_pan"`
	AlgoId                int32     `gorm:"column:algo_id"`
	AlgoCategory          string    `gorm:"column:algo_category"`
	LastActivityReference string    `gorm:"column:last_activity_reference"`
	NnfField              string    `gorm:"column:nnf_field"`
}

func (ExtractTrade) TableName() string {
    return "extract_trades"
}
