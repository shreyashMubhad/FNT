package constants

const (
	ErrorResponseHeaderSize			= 102
	ExternalLogonInHeaderLen		= 300
	ExternalLogonOutHeaderLen		= 300
	SingOnInMessageLen 				= 278
	DownloadRequestLen				= 70
	TradeConfirmationLen			= 318
	
	
	//ERROR CODE ID
	ERR_INVALID_USER_TYPE			= uint16(16001) 		//Invalid User Type
	ERR_USER_ALREADY_SIGNED_ON 		= uint16(16004) 		//User already signed on.
	ERR_INVALID_SIGNON 				= uint16(16006) 		//Invalid sign-on, Please try again.
	ERR_SIGNON_NOT_POSSIBLE 		= uint16(16007) 		//Signing on to the trading system is restricted. Please try later on.
	ERR_INVALID_BROKER_OR_BRANCH 	= uint16(16041) 		//Trading Member does not exist in the system.
	ERR_USER_NOT_FOUND 				= uint16(16042) 		//Dealer does not exist in the system.
	ERR_PROGRAM_ERROR 				= uint16(16056) 		//Program error.
	ERR_SYSTEM_ERROR 				= uint16(16104) 		//System could not complete your transaction - ADMIN notified.
	ERR_CANT_COMPLETE_YOUR_REQUEST 	= uint16(16123) 		//System not able to complete your request. Please try again.
	ERR_USER_IS_DISABLED 			= uint16(16134) 		//This Dealer is disabled. Please call the Exchange.
	ERR_INVALID_USER_ID 			= uint16(16148) 		//Invalid Dealer ID entered.
	ERR_INVALID_TRADER_ID 			= uint16(16154) 		//Invalid Trader ID entered.
	ERR_BROKER_NOT_ACTIVE 			= uint16(16285) 		//The broker is not active.
	
	// Transaction Codes
	
	SIGNON_IN 						= uint16(2300) 	//SIGN_ON_REQUEST_IN 
	HEARTBEAT 						= uint16(23506) //HEARTBEAT
	SIGNON_OUT	 					= uint16(2301) 	//SIGN_ON_REQUEST_OUT
	DC_DOWNLOAD 					= uint16(8000) 	//DROP_COPY_DOWNLOAD_REQUEST
	TRADE_CONFIRMATION 				= uint16(2222) 	//TRADE_CONFIRMATION
	
	// resuest type
	SIGNON_IN_REQUEST				= "signon-in"
	HEARTBEAT_REQUESRT				= "heartbeat"
	SIGN_ON_OUT_RESPONCE			= "signon-out"
	DC_DOWNLOAD_REQUEST				= "drop-copy-downland-request"
	ERROR_RESPONCE					= "error-responce"
	TRADE_CONFIRMATION_RESPONCE		= "trade-confirmation"
	
	//user info
	Password     					= "pass@1234567"
	Brokerid						= "07730"
	
	
	ERR_INVALID_USER_MSG				= "Invalid User Type"
	ERR_USER_ALREADY_SIGNED_ON_MSG		= "User already signed on"
	ERR_INVALID_SIGNON_MSG				= "Invalid sign-on"
	ERR_SIGNON_NOT_POSSIBLE_MSG			= "Signing on to trading system is restricted"
	ERR_INVALID_BROKER_OR_BRANCH_MSG	= "Trading Member does not exist"
	ERR_USER_NOT_FOUND_MSG				= "Dealer does not exist"
	ERR_PROGRAM_ERROR_MSG				= "Program error"
	ERR_SYSTEM_ERROR_MSG				= "System could not complete your transaction"
	ERR_CANT_COMPLETE_YOUR_REQUEST_MSG	= "System not able to complete your request"
	ERR_USER_IS_DISABLED_MSG			= "This Dealer is disabled"
	ERR_INVALID_USER_ID_MSG				= "Invalid Dealer ID entered"
	ERR_INVALID_TRADER_ID_MSG			= "Invalid Trader ID entered"
	ERR_BROKER_NOT_ACTIVE_MSG			= "The broker is not active"
	)