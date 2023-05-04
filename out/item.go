package out

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/erexo/gobiassets/in"
)

type Item struct {
	ServerId    uint16
	ClientId    uint16
	Name        string
	Role        ItemRole
	Description string
	Weight      float32
	Worth       int64
	Attributes  Attributes
}

type ItemAttribute uint8

const (
	ItemAttributeAttack ItemAttribute = iota
	ItemAttributeDefense
	ItemAttributeRange
	ItemAttributeBreakChance
	ItemAttributeArmor
	ItemAttributeContainerSize
	ItemAttributeDuration
	ItemAttributeCharges
	ItemAttributeLevel

	ItemAttributeSpeed
	ItemAttributePreventDrop
	ItemAttributeReduceDeath
	ItemAttributeManaShield
	ItemAttributeMaxHealth
	ItemAttributeMaxHealthPercent
	ItemAttributeMaxMana
	ItemAttributeMaxManaPercent
	ItemAttributeHealthTicks
	ItemAttributeHealthGain
	ItemAttributeManaTicks
	ItemAttributeManaGain
	ItemAttributeSoul
	ItemAttributeSkillAll
	ItemAttributeSkillWeapons
	ItemAttributeSkillMagic
	ItemAttributeSkillFist
	ItemAttributeSkillClub
	ItemAttributeSkillSword
	ItemAttributeSkillDist
	ItemAttributeSkillShield
	ItemAttributeSkillAxe
	ItemAttributeSkillFish
	ItemAttributeMagicPercent
	ItemAttributeMagicPvePercent
	ItemAttributeMeleePercent
	ItemAttributeHealingPercent
	ItemAttributeProtection
)

type Attributes []byte

type AttributeValue struct {
	Attribute ItemAttribute
	Value     int32
}

const attSize = 5

func (w Attributes) ReadAttribute(attribute ItemAttribute) (int32, bool) {
	for i := 0; i+attSize <= len(w); i += attSize {
		if ItemAttribute(w[i]) == attribute {
			return int32(w[i+1]) |
				int32(w[i+2])<<8 |
				int32(w[i+3])<<16 |
				int32(w[i+4])<<24, true
		}
	}
	return 0, false
}

func (w Attributes) ReadAttributes() []AttributeValue {
	ret := []AttributeValue{}
	for i := 0; i+attSize <= len(w); i += attSize {
		ret = append(ret, AttributeValue{
			Attribute: ItemAttribute(w[i]),
			Value: int32(w[i+1]) |
				int32(w[i+2])<<8 |
				int32(w[i+3])<<16 |
				int32(w[i+4])<<24,
		})
	}
	return ret
}

func (w *Attributes) writeAttr(attr ItemAttribute, value int64) {
	if value == 0 {
		return
	}
	if value > math.MaxInt32 {
		panic("attribute bigger than int32")
	}

	*w = append(*w, uint8(attr),
		uint8(value),
		uint8(value>>8),
		uint8(value>>16),
		uint8(value>>24))
}

func NewItem(client uint16, item *in.Item, attrItem *in.Item) *Item {
	return &Item{
		ServerId:    uint16(item.Id),
		ClientId:    client,
		Name:        Title(item.Name),
		Role:        Role(item.Attributes.ReadString("role")),
		Description: item.Attributes.ReadString("description"),
		Weight:      float32(item.Attributes.Read("weight")) / 100,
		Worth:       item.Attributes.Read("worth"),
		Attributes:  getAttrs(attrItem),
	}
}

func getAttrs(item *in.Item) Attributes {
	attr := Attributes{}
	attr.writeAttr(ItemAttributeAttack, item.Attributes.Read("attack", "extraatk", "elementphysical", "elementfire", "elementenergy", "elementearth", "elementice", "elementholy", "elementdeath"))
	attr.writeAttr(ItemAttributeDefense, item.Attributes.Read("defense", "extradef"))
	attr.writeAttr(ItemAttributeRange, item.Attributes.Read("range"))
	attr.writeAttr(ItemAttributeBreakChance, item.Attributes.Read("breakchance"))
	attr.writeAttr(ItemAttributeArmor, item.Attributes.Read("armor"))
	attr.writeAttr(ItemAttributeContainerSize, item.Attributes.Read("containersize"))
	attr.writeAttr(ItemAttributeDuration, item.Attributes.Read("duration"))
	attr.writeAttr(ItemAttributeCharges, item.Attributes.Read("charges"))
	attr.writeAttr(ItemAttributeLevel, item.Attributes.Read("level"))
	attr.writeAttr(ItemAttributeSpeed, item.Attributes.Read("speed"))
	attr.writeAttr(ItemAttributePreventDrop, item.Attributes.Read("preventdrop"))
	attr.writeAttr(ItemAttributeReduceDeath, item.Attributes.Read("reducedeathpercent"))
	attr.writeAttr(ItemAttributeManaShield, item.Attributes.Read("manashield"))
	attr.writeAttr(ItemAttributeMaxHealth, item.Attributes.Read("maxhealthpoints"))
	attr.writeAttr(ItemAttributeMaxHealthPercent, item.Attributes.ReadPercent("maxhealthpercent"))
	attr.writeAttr(ItemAttributeMaxMana, item.Attributes.Read("maxmanapoints"))
	attr.writeAttr(ItemAttributeMaxManaPercent, item.Attributes.ReadPercent("maxmanapercent"))
	attr.writeAttr(ItemAttributeHealthTicks, item.Attributes.Read("healthticks"))
	attr.writeAttr(ItemAttributeHealthGain, item.Attributes.Read("healthgain"))
	attr.writeAttr(ItemAttributeManaTicks, item.Attributes.Read("manaticks"))
	attr.writeAttr(ItemAttributeManaGain, item.Attributes.Read("managain"))
	attr.writeAttr(ItemAttributeSoul, item.Attributes.Read("soulpoints"))
	attr.writeAttr(ItemAttributeSkillAll, item.Attributes.Read("allskills"))
	attr.writeAttr(ItemAttributeSkillWeapons, item.Attributes.Read("skillweapons"))
	attr.writeAttr(ItemAttributeSkillMagic, item.Attributes.Read("magiclevelpoints"))
	attr.writeAttr(ItemAttributeSkillFist, item.Attributes.Read("skillfist"))
	attr.writeAttr(ItemAttributeSkillClub, item.Attributes.Read("skillclub"))
	attr.writeAttr(ItemAttributeSkillSword, item.Attributes.Read("skillsword"))
	attr.writeAttr(ItemAttributeSkillDist, item.Attributes.Read("skilldist"))
	attr.writeAttr(ItemAttributeSkillShield, item.Attributes.Read("skillshield"))
	attr.writeAttr(ItemAttributeSkillAxe, item.Attributes.Read("skillaxe"))
	attr.writeAttr(ItemAttributeSkillFish, item.Attributes.Read("skillfish"))
	attr.writeAttr(ItemAttributeMagicPercent, item.Attributes.ReadPercent("increasemagicpercent"))
	attr.writeAttr(ItemAttributeMagicPvePercent, item.Attributes.ReadPercent("increasemagicpvepercent"))
	attr.writeAttr(ItemAttributeMeleePercent, item.Attributes.ReadPercent("increasemeleepercent"))
	attr.writeAttr(ItemAttributeHealingPercent, item.Attributes.ReadPercent("increasehealingpercent"))
	attr.writeAttr(ItemAttributeProtection, item.Attributes.Read("absorbpercentall"))
	return attr
}

func ItemType() string {
	return `type Item struct {
	ServerId    uint16
	ClientId    uint16
	Name        string
	Role        ItemRole
	Description string
	Weight      float32
	Worth       int64
	Attributes  Attributes
}

type ItemAttribute uint8

const (
	ItemAttributeAttack ItemAttribute = iota
	ItemAttributeDefense
	ItemAttributeRange
	ItemAttributeBreakChance
	ItemAttributeArmor
	ItemAttributeContainerSize
	ItemAttributeDuration
	ItemAttributeCharges
	ItemAttributeLevel

	ItemAttributeSpeed
	ItemAttributePreventDrop
	ItemAttributeReduceDeath
	ItemAttributeManaShield
	ItemAttributeMaxHealth
	ItemAttributeMaxHealthPercent
	ItemAttributeMaxMana
	ItemAttributeMaxManaPercent
	ItemAttributeHealthTicks
	ItemAttributeHealthGain
	ItemAttributeManaTicks
	ItemAttributeManaGain
	ItemAttributeSoul
	ItemAttributeSkillAll
	ItemAttributeSkillWeapons
	ItemAttributeSkillMagic
	ItemAttributeSkillFist
	ItemAttributeSkillClub
	ItemAttributeSkillSword
	ItemAttributeSkillDist
	ItemAttributeSkillShield
	ItemAttributeSkillAxe
	ItemAttributeSkillFish
	ItemAttributeMagicPercent
	ItemAttributeMagicPvePercent
	ItemAttributeMeleePercent
	ItemAttributeHealingPercent
	ItemAttributeProtection
)

type Attributes []byte

type AttributeValue struct {
	Attribute ItemAttribute
	Value     int32
}

const attSize = 5

func (w Attributes) ReadAttribute(attribute ItemAttribute) (int32, bool) {
	for i := 0; i+attSize <= len(w); i += attSize {
		if ItemAttribute(w[i]) == attribute {
			return int32(w[i+1]) |
				int32(w[i+2])<<8 |
				int32(w[i+3])<<16 |
				int32(w[i+4])<<24, true
		}
	}
	return 0, false
}

func (w Attributes) ReadAttributes() []AttributeValue {
	ret := []AttributeValue{}
	for i := 0; i+attSize <= len(w); i += attSize {
		ret = append(ret, AttributeValue{
			Attribute: ItemAttribute(w[i]),
			Value: int32(w[i+1]) |
				int32(w[i+2])<<8 |
				int32(w[i+3])<<16 |
				int32(w[i+4])<<24,
		})
	}
	return ret
}`
}

func (i *Item) String() string {
	var attrsStr strings.Builder
	if len(i.Attributes) == 0 {
		attrsStr.WriteString("nil")
	} else {
		attrsStr.WriteString("[]byte{")
		for j, v := range i.Attributes {
			attrsStr.WriteString(strconv.Itoa(int(v)))
			if j < len(i.Attributes)-1 {
				attrsStr.WriteString(", ")
			}
		}
		attrsStr.WriteByte('}')
	}
	return fmt.Sprintf(`Item{%d, %d, "%s", %d, "%s", %.2f, %d, %s}`,
		i.ServerId, i.ClientId, i.Name, i.Role, i.Description, i.Weight, i.Worth, attrsStr.String())
}
