package utils

import (
	constants "DC/FnO/constant"
	"DC/FnO/pkg/handlers"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils/conversions"
	"encoding/binary"

	// "encoding/hex"
	"errors"
	"fmt"
	"net"

	"go.uber.org/zap"
)

func Send(conn net.Conn, data interface{}, requestType string) error {
	if conn == nil {
		logger.Log().Error("Connection object is nil")
		return fmt.Errorf("connection is nil")
	}

	logger.Log().Debug("Preparing to send data", zap.String("requestType", requestType))

	dataToSend, err := conversions.ToLitteEndian(data, requestType)
	// m, err := conn.Read(dataToSend)

	// logger.Log().Info("Bytes send", zap.Int("size", m))

	if err != nil {
		logger.Log().Error("Failed to convert the request to little-endian format", zap.Error(err))
		return err
	}

	if len(dataToSend) == 0 {
		logger.Log().Error("Data to send is empty", zap.String("requestType", requestType))
		return fmt.Errorf("no data to send for request type: %s", requestType)
	}

	logger.Log().Info("  ", zap.Int("bytes send", len(dataToSend)))

	// logger.Log().Debug("Converted data to little-endian format", zap.String("hex-format", hex.EncodeToString(dataToSend)))

	if _, err := conn.Write(dataToSend); err != nil {
		logger.Log().Error("Failed to send the data over the connection", zap.Error(err))
		return err
	}

	// logger.Log().Info("Data sent successfully",zap.String("requestType", requestType),zap.Int("bytesWritten", size),zap.Any("dataInByteFormat", dataToSend),)
	return nil
}

func Recv(conn net.Conn) (interface{}, []byte, int, uint32, error) {
	// Define a fixed buffer size for reading
	bytesReceived := make([]byte, 1024)
	n, err := conn.Read(bytesReceived)
	if err != nil {
		logger.Log().Error("Failed to read")
		return nil, nil, 0, 0, err
	}

	// fmt.Println("Bytes received before decoding:", bytesReceived)

	type HeaderWithTransCode struct {
		BodyLen         uint16
		SeqNum          uint32
		CheckSum        [16]byte
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

	headerWithTrans := &HeaderWithTransCode{}
	if err := conversions.FromLittleEndian(bytesReceived[:62], headerWithTrans, "sign-on"); err != nil {
		logger.Log().Error("Unable to convert from to little endian")
	} else {
		// logger.Log().Debug("Decoded HeaderWithTransCode", zap.Any("header", headerWithTrans))
	}

	var dst interface{}
	var streamCount int
	var lastSeqSendToExcng uint32

	// Check the TransactionCode of the response and call the appropriate function
	switch headerWithTrans.TransactionCode {
	case constants.SIGNON_OUT:
		switch headerWithTrans.ErrorCode {
		case 0:
			logger.Log().Info("", zap.Int("bytes received", n))
			dst = &models.ExternalLogonOutHeader{}
			logger.Log().Info("LogonOut response received")
			var err error
			streamCount, err = handlers.HandleLogonOut(bytesReceived[:n], dst)
			if err != nil {
				logger.Log().Error("Failed to handle LogonOut response", zap.Error(err))
				return nil, nil, 0, 0, err
			}
			if streamCount > 0 {
				return dst, bytesReceived[:n], streamCount, 0, nil
			}
		default:
			logger.Log().Info("", zap.Int("bytes received", n))
			dst = &models.ExternalLogonOutError{}
			errorMessage := handlers.HandleErrorResponse(bytesReceived[:n], dst)
			if errorMessage != "" && errorMessage != "Unknown error code received" {
				logger.Log().Error("user logonIn failed, ", zap.String("error message", errorMessage))
				return nil, nil, 0, 0, errors.New(errorMessage)
			}
		}
	case constants.TRADE_CONFIRMATION:

		logger.Log().Debug("We are inside the case of Trade Confermation 1 ***************************")

		lastSend := headerWithTrans.TimeStamp1
		lastSeqSendToExcng = uint32(binary.LittleEndian.Uint64(lastSend[:]))
		logger.Log().Debug("lastSeqSendToExcng : ", zap.Uint32("", lastSeqSendToExcng))
		logger.Log().Info("", zap.Int("bytes received", n))
		logger.Log().Info("TradeConfirmation response received")
		trade, err := handlers.HandleTradeConfirmation(bytesReceived[:n])
		logger.Log().Debug("We are inside the case of Trade Confermation 2 ***************************")
		if err != nil {
			logger.Log().Error("Failed to handle TradeConfirmation response", zap.Error(err))
			return nil, nil, 0, lastSeqSendToExcng, err
		}
		dst = *trade
		logger.Log().Debug("We are inside the case of Trade Confermation 3 ***************************")
	default:
		logger.Log().Error("Unexpected response size", zap.Int("bytes received", n))
		return nil, nil, 0, 0, fmt.Errorf("unexpected response size: %d", n)
	}
	logger.Log().Debug("We are inside the case of Trade Confermation 4 ***************************")

	logger.Log().Debug("Here we are printing the values of sequence NUmber ", zap.Uint32("", lastSeqSendToExcng))

	return dst, bytesReceived[:n], streamCount, lastSeqSendToExcng, nil
}
