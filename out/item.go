package out

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/erexo/gobiaitem/in"
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

func (w Attributes) ReadAttributes() []AttributeValue {
	const attSize = 5
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

func NewItem(client uint16, item *in.Item) *Item {
	attr := Attributes{}
	attr.writeAttr(ItemAttributeAttack, readAttribute(item.Attributes, "attack", "extraatk", "elementphysical", "elementfire", "elementenergy", "elementearth", "elementice", "elementholy", "elementdeath"))
	attr.writeAttr(ItemAttributeDefense, readAttribute(item.Attributes, "defense", "extradef"))
	attr.writeAttr(ItemAttributeRange, readAttribute(item.Attributes, "range"))
	attr.writeAttr(ItemAttributeBreakChance, readAttribute(item.Attributes, "breakchance"))
	attr.writeAttr(ItemAttributeArmor, readAttribute(item.Attributes, "armor"))
	attr.writeAttr(ItemAttributeContainerSize, readAttribute(item.Attributes, "containersize"))
	attr.writeAttr(ItemAttributeDuration, readAttribute(item.Attributes, "duration"))
	attr.writeAttr(ItemAttributeCharges, readAttribute(item.Attributes, "charges"))
	attr.writeAttr(ItemAttributeLevel, readAttribute(item.Attributes, "level"))
	attr.writeAttr(ItemAttributeSpeed, readAttribute(item.Attributes, "speed"))
	attr.writeAttr(ItemAttributePreventDrop, readAttribute(item.Attributes, "preventdrop"))
	attr.writeAttr(ItemAttributeReduceDeath, readAttribute(item.Attributes, "reducedeathpercent"))
	attr.writeAttr(ItemAttributeManaShield, readAttribute(item.Attributes, "manashield"))
	attr.writeAttr(ItemAttributeMaxHealth, readAttribute(item.Attributes, "maxhealthpoints"))
	attr.writeAttr(ItemAttributeMaxHealthPercent, readAttributePercent(item.Attributes, "maxhealthpercent"))
	attr.writeAttr(ItemAttributeMaxMana, readAttribute(item.Attributes, "maxmanapoints"))
	attr.writeAttr(ItemAttributeMaxManaPercent, readAttributePercent(item.Attributes, "maxmanapercent"))
	attr.writeAttr(ItemAttributeHealthTicks, readAttribute(item.Attributes, "healthticks"))
	attr.writeAttr(ItemAttributeHealthGain, readAttribute(item.Attributes, "healthgain"))
	attr.writeAttr(ItemAttributeManaTicks, readAttribute(item.Attributes, "manaticks"))
	attr.writeAttr(ItemAttributeManaGain, readAttribute(item.Attributes, "managain"))
	attr.writeAttr(ItemAttributeSoul, readAttribute(item.Attributes, "soulpoints"))
	attr.writeAttr(ItemAttributeSkillAll, readAttribute(item.Attributes, "allskills"))
	attr.writeAttr(ItemAttributeSkillWeapons, readAttribute(item.Attributes, "skillweapons"))
	attr.writeAttr(ItemAttributeSkillMagic, readAttribute(item.Attributes, "magiclevelpoints"))
	attr.writeAttr(ItemAttributeSkillFist, readAttribute(item.Attributes, "skillfist"))
	attr.writeAttr(ItemAttributeSkillClub, readAttribute(item.Attributes, "skillclub"))
	attr.writeAttr(ItemAttributeSkillSword, readAttribute(item.Attributes, "skillsword"))
	attr.writeAttr(ItemAttributeSkillDist, readAttribute(item.Attributes, "skilldist"))
	attr.writeAttr(ItemAttributeSkillShield, readAttribute(item.Attributes, "skillshield"))
	attr.writeAttr(ItemAttributeSkillAxe, readAttribute(item.Attributes, "skillaxe"))
	attr.writeAttr(ItemAttributeSkillFish, readAttribute(item.Attributes, "skillfish"))
	attr.writeAttr(ItemAttributeMagicPercent, readAttributePercent(item.Attributes, "increasemagicpercent"))
	attr.writeAttr(ItemAttributeMagicPvePercent, readAttributePercent(item.Attributes, "increasemagicpvepercent"))
	attr.writeAttr(ItemAttributeMeleePercent, readAttributePercent(item.Attributes, "increasemeleepercent"))
	attr.writeAttr(ItemAttributeHealingPercent, readAttributePercent(item.Attributes, "increasehealingpercent"))
	attr.writeAttr(ItemAttributeProtection, readAttribute(item.Attributes, "absorbpercentall"))

	return &Item{
		ServerId:    uint16(item.Id),
		ClientId:    client,
		Name:        Title(item.Name),
		Role:        Role(readAttributeString(item.Attributes, "role")),
		Description: readAttributeString(item.Attributes, "description"),
		Weight:      float32(readAttribute(item.Attributes, "weight")) / 100,
		Worth:       readAttribute(item.Attributes, "worth"),
		Attributes:  attr,
	}
}

func readAttribute(attr in.Attributes, names ...string) int64 {
	var ret int64
	for _, name := range names {
		if v, ok := attr[strings.ToLower(name)]; ok {
			value, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Println(err)
			}
			ret += value
		}
	}
	return ret
}

func readAttributeString(attr in.Attributes, name string) string {
	if v, ok := attr[strings.ToLower(name)]; ok {
		return v
	}
	return ""
}

func readAttributePercent(attr in.Attributes, name string) int64 {
	v := readAttribute(attr, name)
	if v == 0 {
		return 0
	}
	return v - 100
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

func (w Attributes) ReadAttributes() []AttributeValue {
	const attSize = 5
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
