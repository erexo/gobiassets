package in

const (
	nodeStart  = 0xfe
	nodeEnd    = 0xff
	escapeChar = 0xfd
)

type BinaryTreeReader struct {
	BinaryNode
}

func NewBinaryTreeReader(data []byte) *BinaryTreeReader {
	return &BinaryTreeReader{
		BinaryNode{
			buffer: data,
		},
	}
}

func (b *BinaryTreeReader) GetNode() *BinaryNode {
	if !b.advance() {
		panic("advance")
	}
	return b.getNodeData()
}

func (b *BinaryTreeReader) GetNextNode() *BinaryNode {
	value := b.ReadU8()
	if value != nodeStart {
		return nil
	}
	value = b.ReadU8()

	level := 1
	for {
		value = b.ReadU8()
		switch value {
		case nodeEnd:
			level--
			if level == 0 {
				value = b.ReadU8()
				if value != nodeStart {
					return nil
				}
				b.index--
				return b.getNodeData()
			}
		case nodeStart:
			level++
		case escapeChar:
			b.ReadU8()
		}
	}
}

func (b *BinaryTreeReader) getNodeData() *BinaryNode {
	start := b.index
	value := b.ReadU8()
	if value != nodeStart {
		return nil
	}
	var node []byte
	for {
		value = b.ReadU8()
		if value == nodeEnd || value == nodeStart {
			break
		}
		if value == escapeChar {
			value = b.ReadU8()
		}
		node = append(node, value)
	}
	b.index = start
	return &BinaryNode{node, 0}
}

func (b *BinaryTreeReader) advance() bool {
	start := false
	if b.index == 0 {
		start = true
		b.index = 4
	}
	if b.ReadU8() != nodeStart {
		return false
	}
	if start {
		b.index--
		return true
	}

	value := b.ReadU8()
	for {
		value = b.ReadU8()
		switch value {
		case nodeEnd:
			return false
		case nodeStart:
			b.index--
			return true
		case escapeChar:
			b.ReadU8()
		}
	}
}
