package out

type BossClass uint8

const (
	BossClassNone BossClass = iota
	BossClassRegular
	BossClassDaily
	BossClassMini
)

func BossClassPrefix() string {
	return `type BossClass uint8

const (
	BossClassNone BossClass = iota
	BossClassRegular
	BossClassDaily
	BossClassMini
)`
}
