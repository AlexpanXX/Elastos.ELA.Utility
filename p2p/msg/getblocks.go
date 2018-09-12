package msg

import (
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

// MaxBlockLocatorsPerMsg is the maximum number of block locator hashes allowed
// per message.
const MaxBlockLocatorsPerMsg = 500

// Ensure GetBlocks implement p2p.Message interface.
var _ p2p.Message = (*GetBlocks)(nil)

type GetBlocks struct {
	Locator  []*common.Uint256
	HashStop common.Uint256
}

func NewGetBlocks(locator []*common.Uint256, hashStop common.Uint256) *GetBlocks {
	msg := new(GetBlocks)
	msg.Locator = locator
	msg.HashStop = hashStop
	return msg
}

func (msg *GetBlocks) CMD() string {
	return p2p.CmdGetBlocks
}

func (msg *GetBlocks) MaxLength() uint32 {
	return 4 + (MaxBlockLocatorsPerMsg * common.UINT256SIZE) + common.UINT256SIZE
}

func (msg *GetBlocks) Serialize(writer io.Writer) error {
	return common.WriteElements(writer, uint32(len(msg.Locator)), msg.Locator, msg.HashStop)
}

func (msg *GetBlocks) Deserialize(reader io.Reader) error {
	count, err := common.ReadUint32(reader)
	if err != nil {
		return err
	}

	for i:= uint32(0); i< count ; i++{
		var hash common.Uint256
		if err := hash.Deserialize(reader); err != nil {
			return err
		}
		msg.Locator = append(msg.Locator, &hash)
	}

	return msg.HashStop.Deserialize(reader)
}
