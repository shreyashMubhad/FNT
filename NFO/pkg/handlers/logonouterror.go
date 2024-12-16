package handlers

import (
	constants "DC/FnO/constant"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils/conversions"

	"go.uber.org/zap"
)

func HandleErrorResponse(data []byte, dts interface{}) string {
	// Parse the packet using `fromLittleEndian`
	if err := conversions.FromLittleEndian(data, dts, "error-response"); err != nil {
		logger.Log().Error("Failed to parse the packet", zap.Error(err))
		return "Failed to parse error-response packet"
	}

	// Type assertion to ensure `dts` is of the correct type
	response, ok := dts.(*models.ExternalLogonOutError)
	if !ok {
		logger.Log().Error("Unexpected type for error response", zap.Any("response", dts))
		return "Unexpected response type"
	}

	// Map error code to an appropriate message
	switch response.MsgData.MessageHeader.ErrorCode {
	case constants.ERR_INVALID_USER_TYPE:
		return constants.ERR_INVALID_USER_MSG
	case constants.ERR_USER_ALREADY_SIGNED_ON:
		return constants.ERR_USER_ALREADY_SIGNED_ON_MSG
	case constants.ERR_INVALID_SIGNON:
		return constants.ERR_INVALID_SIGNON_MSG
	case constants.ERR_SIGNON_NOT_POSSIBLE:
		return constants.ERR_SIGNON_NOT_POSSIBLE_MSG
	case constants.ERR_INVALID_BROKER_OR_BRANCH:
		return constants.ERR_INVALID_BROKER_OR_BRANCH_MSG
	case constants.ERR_USER_NOT_FOUND:
		return constants.ERR_USER_NOT_FOUND_MSG
	case constants.ERR_PROGRAM_ERROR:
		return constants.ERR_PROGRAM_ERROR_MSG
	case constants.ERR_SYSTEM_ERROR:
		return constants.ERR_SYSTEM_ERROR_MSG
	case constants.ERR_CANT_COMPLETE_YOUR_REQUEST:
		return constants.ERR_CANT_COMPLETE_YOUR_REQUEST_MSG
	case constants.ERR_USER_IS_DISABLED:
		return constants.ERR_USER_IS_DISABLED_MSG
	case constants.ERR_INVALID_USER_ID:
		return constants.ERR_INVALID_USER_ID_MSG
	case constants.ERR_INVALID_TRADER_ID:
		return constants.ERR_INVALID_TRADER_ID_MSG
	case constants.ERR_BROKER_NOT_ACTIVE:
		return constants.ERR_BROKER_NOT_ACTIVE_MSG
	default:
		return "Unknown error code received"
	}
}
