package out

type MonsterCategory uint8

const (
	MonsterCategoryMonsters MonsterCategory = iota
	MonsterCategoryBosses
	MonsterCategorySaga

	MonsterCategoryFirst = MonsterCategoryMonsters
	MonsterCategoryLast  = MonsterCategorySaga
)

func MonsterCategoryPrefix() string {
	return `//go:generate enumer -type=MonsterCategory -trimprefix MonsterCategory -output monstercategory_string.go

type MonsterCategory uint8

const (
	MonsterCategoryMonsters MonsterCategory = iota
	MonsterCategoryBosses
	MonsterCategorySaga

	MonsterCategoryFirst = MonsterCategoryMonsters
	MonsterCategoryLast  = MonsterCategorySaga
)`
}
