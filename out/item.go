package out

import (
	"fmt"

	"github.com/erexo/gobiaitem/in"
)

type Item struct {
	ServerId uint16
	ClientId uint16
	Name     string
	Role     ItemRole
}

func NewItem(client uint16, item *in.Item) *Item {
	role := ItemRoleAll
	if r, ok := item.Attributes["role"]; ok {
		role = Role(r)
	}
	return &Item{
		ServerId: uint16(item.Id),
		ClientId: client,
		Name:     Title(item.Name),
		Role:     role,
	}
}

func Type() string {
	return `type Item struct {
	ServerId uint16
	ClientId uint16
	Name     string
	Role     ItemRole
}`
}

func (i *Item) String() string {
	return fmt.Sprintf(`Item{%d, %d, "%s", %d}`, i.ServerId, i.ClientId, i.Name, i.Role)
}
