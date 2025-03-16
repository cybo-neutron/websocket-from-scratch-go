package websocket

import (
	"encoding/binary"
)

type Frame struct {
	IsFragment    bool
	OpCode        byte
	IsMasked      bool
	PayloadLength uint64
	MaskingKey    []byte
	PayloadData   []byte
}

func (f *Frame) HandleIncomingFrame(ws *Websocket) (Frame, error) {
	frame := Frame{}
	head, err := ws.read(2)
	if err != nil {
		return frame, err
	}

	frame.IsFragment = head[0]&0x80 != 0
	frame.OpCode = head[0] & 0x0F
	frame.IsMasked = head[1]&0x80 != 0

	// get length
	length := uint64(head[1] & 0x7F)

	if length == 126 {
		newLengthData, error := ws.read(2)
		if error != nil {
			return frame, error
		}
		length = uint64(binary.BigEndian.Uint16(newLengthData))
	} else if length == 127 {

		newLengthData, error := ws.read(8)
		if error != nil {
			return frame, error
		}
		length = uint64(binary.BigEndian.Uint16(newLengthData))
	}

	frame.PayloadLength = length

	// get masking key
	maskKey, error := ws.read(4)
	if error != nil {
		return frame, error
	}
	frame.MaskingKey = maskKey

	// get paylaod data bytes
	payloadData, err := ws.read(int(length))
	if err != nil {
		return frame, error
	}

	// unmask data
	// for i := 0; i < len(payloadData); i++ {
	// 	payloadData[i] = payloadData[i] ^ maskKey[i%4]
	// }
	for idx, val := range payloadData {
		payloadData[idx] = val ^ maskKey[idx%4]
	}

	frame.PayloadData = payloadData

	return frame, nil
}

func (ws *Websocket) CreateFrameToSend() {
}

func (ws *Websocket) SendFrame(frame Frame) {
	data := make([]byte, 2)
	// data[0] = frame.IsFragment & x80
	// data[0] = frame.OpCode & x0F

	data[0] = data[0] | frame.OpCode // OpCode

	if frame.IsFragment {
		data[0] = data[0] | 0x80
	}
}
