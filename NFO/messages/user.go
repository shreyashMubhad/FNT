package messages

import (
	"DC/FnO/constant"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils"
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
)


func newLogonIN() *models.LogonIN {

	logger.Log().Debug("Initializing new LogonIN structure")

    passwordBytes := [12]byte{}
    copy(passwordBytes[:], constants.Password)
    
    brokeridBytes := [5]byte{}
    copy(brokeridBytes[:], constants.Brokerid)
    

    logonIN :=  &models.LogonIN{
        MessageHeader: models.InternalHeader{
            TransactionCode: constants.SIGNON_IN,
            LogTime:         0,
            AlphaChar:       [2]byte{}, // this field contains the number of streams from which the drop copy data feed is sent. User needs to send the DC_DOWNLOAD_REQUEST (8000) for each stream to download the trade data
            UserId:          32106,     //47409
            ErrorCode:       0,
            TimeStamp:       [8]byte{},
            TimeStamp1:      [8]byte{},
            TimeStamp2:      [8]byte{},
            MessageLength:   constants.SingOnInMessageLen,
        },
        UserId:    32106,
        Reserved:  [8]byte{},
        Password:  passwordBytes, // FfeA4145
        Reserved1: [4]byte{},
        Reserved2: [38]byte{},
        BrokerId:  brokeridBytes,
        Reserved3: [119]byte{},
        Reserved4: [16]byte{},
        Reserved5: [16]byte{},
        Reserved6: [16]byte{},
    }
    // if some isshue in size uncomment the below line to check the size of the structure
	// logger.Log().Info("LogonIN structure created", zap.Any("LogonIN", logonIN))
	return logonIN
}

func getUserLoginRequest() models.ExternalLogonInHeader {
	logger.Log().Debug("Creating user login request")
	
	logonHeader := newLogonIN()
    checksum := utils.GetCheckSum(logonHeader)
    logger.Log().Debug("Checksum calculated", zap.ByteString("Checksum", checksum[:]))

    request := models.ExternalLogonInHeader{
        BodyLen:  constants.ExternalLogonInHeaderLen,
        SeqNum:   utils.GetSequenceNumber(),
        CheckSum: [16]byte(checksum),
        MsgData:  *logonHeader,
    }
	return request
}

func UserLogin(ctx context.Context, conn net.Conn) (int, error) {
    if conn == nil {
        logger.Log().Error("Connection object is nil")
        return 0, fmt.Errorf("connection is nil")
    }

    logger.Log().Debug("User-Login Started")
    defer logger.Log().Debug("User-Login Ended")

    // Build the request
    request := getUserLoginRequest()

    logger.Log().Debug("sending user logonIn request")
    // Send request
    if err := utils.Send(conn, request, constants.SIGNON_IN_REQUEST); err != nil {
        logger.Log().Error("Failed to send user-login request", zap.Error(err))
        return 0, fmt.Errorf("send user-login request: %w", err)
    }
    logger.Log().Info("Request sent", zap.Any("request", request))

    // Receive response
    dst, _, streamCount, _, err := utils.Recv(conn)
    if err != nil {
        logger.Log().Error("Failed to receive user-login response", zap.Error(err))
        return 0, fmt.Errorf("receive user-login response: %w", err)
    }
    logger.Log().Info("user logonOut response received", zap.Any("parsedResponse", dst))

    return streamCount, nil
}
