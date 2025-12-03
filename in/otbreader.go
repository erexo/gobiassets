package in

import (
	"os"
	"path/filepath"
)

type (
	ServerClientMap = map[uint16]uint16
	ClientServerMap = map[uint16][]uint16
)

func ReadOtb() (ServerClientMap, ClientServerMap) {
	data, err := os.ReadFile(filepath.Join(dataPath, itemsOtbFile))
	if err != nil {
		panic(err)
	}

	tree := NewBinaryTreeReader(data)
	node := tree.GetNode()

	node.ReadU8()  // 0
	node.ReadU32() // flags

	attr := node.ReadU8()
	if attr == 0x01 { // version
		if node.ReadU16() != 140 { // datalen
			panic("datalen")
		}
		node.ReadU32() // ver
		node.ReadU32()
		node.ReadU32()
	}

	node = tree.GetNode()
	if node == nil {
		panic("node")
	}

	serverClient := make(ServerClientMap)
	clientServer := make(ClientServerMap)
	for {
		node.ReadU8()
		_ = node.ReadU32() // flags
		var serverId, clientId uint16
		for !node.Empty() {
			attribute := node.ReadU8()
			datalen := node.ReadU16()
			switch attribute {
			case attributeServerID:
				serverId = node.ReadU16()
			case attributeClientID:
				clientId = node.ReadU16()
			case attributeGroundSpeed:
				node.ReadU16()
			case attributeName:
				node.Read(datalen)
			case attributeSpriteHash:
				node.Read(datalen)
			case attributeMinimaColor:
				node.ReadU16()
			case attributeMaxReadWriteChars:
				node.ReadU16()
			case attributeMaxReadChars:
				node.ReadU16()
			case attributeLight:
				node.ReadU16()
				node.ReadU16()
			case attributeStackOrder:
				node.ReadU8()
			case attributeTradeAs:
				node.ReadU16()
			default:
				node.Skip(datalen)
			}
		}
		//if flags&(1<<5) != 0 { // pickupable
		serverClient[serverId] = clientId
		//}
		clientServer[clientId] = append(clientServer[clientId], serverId)

		node = tree.GetNextNode()
		if node == nil {
			break
		}
	}
	return serverClient, clientServer
}

const (
	attributeServerID          = 0x10
	attributeClientID          = 0x11
	attributeName              = 0x12
	attributeGroundSpeed       = 0x14
	attributeSpriteHash        = 0x20
	attributeMinimaColor       = 0x21
	attributeMaxReadWriteChars = 0x22
	attributeMaxReadChars      = 0x23
	attributeLight             = 0x2A
	attributeStackOrder        = 0x2B
	attributeTradeAs           = 0x2D
)
