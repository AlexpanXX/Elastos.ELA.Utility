package msg

import (
	"fmt"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

const defaultInvListSize = 100

// Ensure Inv implement p2p.Message interface.
var _ p2p.Message = (*Inv)(nil)

type Inv struct {
	InvList []*InvVect
}

func NewInv() *Inv {
	msg := &Inv{
		InvList: make([]*InvVect, 0, defaultInvListSize),
	}
	return msg
}

// AddInvVect adds an inventory vector to the message.
func (msg *Inv) AddInvVect(iv *InvVect) error {
	if len(msg.InvList)+1 > MaxInvPerMsg {
		return fmt.Errorf("AddInvVect too many invvect in message [max %v]", MaxInvPerMsg)
	}

	msg.InvList = append(msg.InvList, iv)
	return nil
}

func (msg *Inv) CMD() string {
	return p2p.CmdInv
}

func (msg *Inv) MaxLength() uint32 {
	return 4 + (MaxInvPerMsg * maxInvVectPayload)
}

func (msg *Inv) Serialize(writer io.Writer) error {
	err := common.WriteUint32(writer, uint32(len(msg.InvList)))
	if err != nil {
		return err
	}

	for _, vect := range msg.InvList {
		if err := vect.Serialize(writer); err != nil {
			return err
		}
	}

	return nil
}

func (msg *Inv) Deserialize(reader io.Reader) error {
	count, err := common.ReadUint32(reader)
	if err != nil {
		return err
	}
	// Limit to max inventory vectors per message.
	if count > MaxInvPerMsg {
		return fmt.Errorf("too many invvect in message [%v]", count)
	}

	msg.InvList = make([]*InvVect, 0, count)
	for i := uint32(0); i < count; i++ {
		var vect InvVect
		if err := vect.Deserialize(reader); err != nil {
			return err
		}
		msg.InvList = append(msg.InvList, &vect)
	}

	return nil
}
