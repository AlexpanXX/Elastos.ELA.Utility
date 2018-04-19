package msg

import (
	"bytes"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

type DataReq struct {
	Type uint8
	Hash Uint256
}

func NewDataReq(invType uint8, hash Uint256) *DataReq {
	dataReq := new(DataReq)
	dataReq.Type = invType
	dataReq.Hash = hash
	return dataReq
}

func (msg *DataReq) CMD() string {
	return "getdata"
}

func (msg *DataReq) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := WriteElements(buf, msg.Type, msg.Hash)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (msg *DataReq) Deserialize(body []byte) error {
	buf := bytes.NewReader(body)
	err := ReadElements(buf, &msg.Type, &msg.Hash)
	if err != nil {
		return err
	}

	return nil
}
