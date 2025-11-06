package out

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/erexo/gobiassets/in"
)

func SageTreeHeader() string {
	return `type SageNode struct {
	Id       uint16
	Type     NodeType
	Icon     NodeIcon
	Stats    []NodeStat
	Position image.Point
}

type NodeType int

const (
	SageNodeTypeRegular NodeType = 0
	SageNodeTypeEpic    NodeType = 1
	SageNodeTypeRoot    NodeType = 2
)

type NodeStat struct {
	Value int16
	Icon  NodeIcon
}

type NodeIcon int

const (
	NodeIconNone             NodeIcon = 0
	NodeIconNinjutsu         NodeIcon = 1
	NodeIconFuinjutsu        NodeIcon = 2
	NodeIconBukijutsu        NodeIcon = 3
	NodeIconTaijutsu         NodeIcon = 4
	NodeIconSenjutsu         NodeIcon = 5
	NodeIconGenjutsu         NodeIcon = 6
	NodeIconVitality         NodeIcon = 7
	NodeIconControl          NodeIcon = 8
	NodeIconSpeed            NodeIcon = 9
	NodeIconHealth           NodeIcon = 10
	NodeIconHealthPercent    NodeIcon = 11
	NodeIconChakra           NodeIcon = 12
	NodeIconChakraPercent    NodeIcon = 13
	NodeIconArmor            NodeIcon = 14
	NodeIconJutsuDamage      NodeIcon = 15
	NodeIconPvEJutsuDamage   NodeIcon = 16
	NodeIconJutsuCritChance  NodeIcon = 17
	NodeIconJutsuCritDamage  NodeIcon = 18
	NodeIconWeaponDamage     NodeIcon = 19
	NodeIconWeaponCritChance NodeIcon = 20
	NodeIconWeaponCritDamage NodeIcon = 21
	NodeIconProtection       NodeIcon = 22
	NodeIconSoul             NodeIcon = 23
)

func (n NodeIcon) Icon() string {
	switch n {
	case NodeIconNinjutsu:
		return "ninjutsu"
	case NodeIconFuinjutsu:
		return "fuinjutsu"
	case NodeIconBukijutsu:
		return "bukijutsu"
	case NodeIconTaijutsu:
		return "taijutsu"
	case NodeIconSenjutsu:
		return "senjutsu"
	case NodeIconGenjutsu:
		return "genjutsu"
	case NodeIconVitality:
		return "vitality"
	case NodeIconControl:
		return "control"
	case NodeIconSpeed:
		return "speed"
	case NodeIconHealth:
		return "health"
	case NodeIconHealthPercent:
		return "hpercent"
	case NodeIconChakra:
		return "chakra"
	case NodeIconChakraPercent:
		return "mpercent"
	case NodeIconArmor:
		return "armor"
	case NodeIconJutsuDamage:
		return "jutsudmg"
	case NodeIconPvEJutsuDamage:
		return "jutsupve"
	case NodeIconJutsuCritChance:
		return "jutsucrit"
	case NodeIconJutsuCritDamage:
		return "jutsucritdmg"
	case NodeIconWeaponDamage:
		return "weapondmg"
	case NodeIconWeaponCritChance:
		return "weaponcrit"
	case NodeIconWeaponCritDamage:
		return "weaponcritdmg"
	case NodeIconProtection:
		return "protection"
	case NodeIconSoul:
		return "soul"
	default:
		return ""
	}
}

func (n NodeIcon) Description(value int32) string {
	switch n {
	case NodeIconNinjutsu:
		return fmt.Sprintf("{一%+d} ninjutsu", value)
	case NodeIconFuinjutsu:
		return fmt.Sprintf("{一%+d} fuinjutsu", value)
	case NodeIconBukijutsu:
		return fmt.Sprintf("{一%+d} bukijutsu", value)
	case NodeIconTaijutsu:
		return fmt.Sprintf("{一%+d} taijutsu", value)
	case NodeIconSenjutsu:
		return fmt.Sprintf("{一%+d} senjutsu", value)
	case NodeIconGenjutsu:
		return fmt.Sprintf("{一%+d} genjutsu", value)
	case NodeIconVitality:
		return fmt.Sprintf("{一%+d} vitality", value)
	case NodeIconControl:
		return fmt.Sprintf("{一%+d} control", value)
	case NodeIconSpeed:
		return fmt.Sprintf("{一%+d} speed", value)
	case NodeIconHealth:
		return fmt.Sprintf("{一%+d} health", value)
	case NodeIconHealthPercent:
		return fmt.Sprintf("{一%+d%%} health", value)
	case NodeIconChakra:
		return fmt.Sprintf("{一%+d} chakra", value)
	case NodeIconChakraPercent:
		return fmt.Sprintf("{一%+d%%} chakra", value)
	case NodeIconArmor:
		return fmt.Sprintf("{一%+d} armor", value)
	case NodeIconJutsuDamage:
		return fmt.Sprintf("{一%+d%%} jutsu damage", value)
	case NodeIconPvEJutsuDamage:
		return fmt.Sprintf("{一%+d%%} PvE jutsu damage", value)
	case NodeIconJutsuCritChance:
		return fmt.Sprintf("{一%+d%%} jutsu crit chance", value)
	case NodeIconJutsuCritDamage:
		return fmt.Sprintf("{一%+d%%} jutsu crit damage", value)
	case NodeIconWeaponDamage:
		return fmt.Sprintf("{一%+d%%} weapon damage", value)
	case NodeIconWeaponCritChance:
		return fmt.Sprintf("{一%+d%%} weapon crit chance", value)
	case NodeIconWeaponCritDamage:
		return fmt.Sprintf("{一%+d%%} weapon crit damage", value)
	case NodeIconProtection:
		v := float64(value) / 10.
		if v == float64(int(v)) {
			return fmt.Sprintf("{一%+.0f%%} protection", v)
		} else {
			return fmt.Sprintf("{一%+.1f%%} protection", v)
		}
	case NodeIconSoul:
		return fmt.Sprintf("{一%+d} soul", value)
	default:
		return ""
	}
}

func (n NodeIcon) DescribeEpic() (title string, description string, bonusDescription string) {
	switch n {
	case 1:
		return "Rikudo Sennin", "", ""
	case 2:
		return "Mistward", "Gain short invisibility when on low health", "   {四Triggers when below} {三20%} {四health}\n   {四Invisibility lasts} {三3 seconds}\n   {四Has} {三2 minutes} {四cooldown}\n   {四Works only in PvP}"
	case 3:
		return "Stormshield", "Absorb portion of the damage as chakra", "   {四Gain} {三1%} {四damage absorption for each} {三2500} {四max chakra}\n   {四ie. having 25000 max chakra will absorb 10% damage}\n   {敗Stormshield will do nothing if Chakra Shield is active}"
	case 4:
		return "Chakra Reflux", "Gain {一25%} chance to reset cooldown on casted special jutsu", "   {四When special is reset soul and chakra is refunded}"
	case 5:
		return "Lotus", "Weapon attacks grant {一+1%} protection for {一2} seconds, stacking up to {一+7%}", ""
	case 6:
		return "Grievous Wounds", "", ""
	case 7:
		return "Precision", "Increase {一weapon damage} the further you are from your target", "   {三+5%} {四weapon damage at} {三3} {四range}\n  {三+10%} {四weapon damage at} {三4} {四range}\n  {三+15%} {四weapon damage at} {三5} {四range}\n  {三+20%} {四weapon damage at} {三6} {四range}"
	case 8:
		return "Close Quarters", "{敗Cannot use distance weapon}", ""
	case 9:
		return "Overgrowth", "", ""
	case 10:
		return "Thorns", "{一25%} of your armor is reflected on weapon damage you take", ""
	default:
		return "", "", ""
	}
}`
}

func SaveSageTree() {
	defer LogTime("SageTree")()

	tree := in.ReadSageTree()

	var nodesStr, linksStr []string
	for _, node := range tree.Nodes {
		var statsStr []string
		for _, stat := range node.Stats {
			statsStr = append(statsStr, fmt.Sprintf(`
			{%d, NodeIcon(%d)},`, stat.Value, stat.Icon))
		}
		tabStr := "    "
		statTabStr := "   "
		if len(statsStr) != 0 {
			statsStr = append(statsStr, "\n		")
			tabStr = ""
			statTabStr = ""
		}
		nodesStr = append(nodesStr, fmt.Sprintf(`
	{
		Id:   %s%d,
		Type: %sNodeType(%d),
		Icon: %sNodeIcon(%d),
		Stats: %s[]NodeStat{%s},
		Position: image.Pt(%d, %d),
	},`, tabStr, node.Id, tabStr, node.Type, tabStr, node.Icon, statTabStr, strings.Join(statsStr, ""), node.Position.X, node.Position.Y))
	}
	for _, link := range tree.Links {
		linksStr = append(linksStr, fmt.Sprintf(`
	{%d, %d},`, link.Aid, link.Bid))
	}

	_ = os.MkdirAll(outputDir, 0755)
	f, err := os.OpenFile(path.Join(outputDir, sageTreeFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, `// Code generated by "gobiassets" using 'go generate'. DO NOT EDIT.

package %s

import (
	"fmt"
	"image"
)

%s

var SageNodes = []*SageNode{%s
}

var SageLinks = []struct{ IdA, IdB uint16 }{%s
}
`, packageName, SageTreeHeader(), strings.Join(nodesStr, ""), strings.Join(linksStr, "")); err != nil {
		panic(err)
	}
}
