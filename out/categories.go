package out

type ItemCategory int

const (
	ItemCategoryContainers ItemCategory = iota
	ItemCategoryHelmets
	ItemCategoryArmors
	ItemCategoryLegs
	ItemCategoryShoes
	ItemCategoryShields
	ItemCategoryAmulets
	ItemCategoryRings
	ItemCategoryScrolls
	ItemCategoryGloves
	ItemCategorySwords
	ItemCategoryDistances
	ItemCategoryBands
	ItemCategoryEdibles
	ItemCategoryValuables
	ItemCategoryCollectables

	ItemCategoryFirst = ItemCategoryContainers
	ItemCategoryLast  = ItemCategoryCollectables
)

type ItemRole uint8

const (
	ItemRoleAll ItemRole = iota
	ItemRoleNinjutsu
	ItemRoleWeapons
	ItemRoleDefense
)

func (i ItemRole) String() string {
	switch i {
	case ItemRoleAll:
		return "All"
	case ItemRoleNinjutsu:
		return "Ninjutsu"
	case ItemRoleWeapons:
		return "Weapons"
	case ItemRoleDefense:
		return "Defense"
	default:
		return "ItemRole"
	}
}

func Category() string {
	return `//go:generate enumer -type=ItemCategory -trimprefix ItemCategory -output itemcategory_string.go
//go:generate enumer -type=ItemRole -trimprefix ItemRole -output itemrole_string.go

type ItemCategory int

const (
	ItemCategoryContainers ItemCategory = iota
	ItemCategoryHelmets
	ItemCategoryArmors
	ItemCategoryLegs
	ItemCategoryShoes
	ItemCategoryShields
	ItemCategoryAmulets
	ItemCategoryRings
	ItemCategoryScrolls
	ItemCategoryGloves
	ItemCategorySwords
	ItemCategoryDistances
	ItemCategoryBands
	ItemCategoryEdibles
	ItemCategoryValuables
	ItemCategoryCollectables

	ItemCategoryFirst = ItemCategoryContainers
	ItemCategoryLast  = ItemCategoryCollectables
)

type ItemRole uint8

const (
	ItemRoleAll ItemRole = iota
	ItemRoleNinjutsu
	ItemRoleWeapons
	ItemRoleDefense
)`
}

var Items = map[ItemCategory][]uint16{
	ItemCategoryContainers: {
		1990,  // Present
		6104,  // Jewel Case
		7587,  // Bag
		1991,  // Konoha Bag
		1992,  // Suna Bag
		1995,  // Kumo Bag
		3939,  // Forest Bag
		5927,  // Pirate Bag
		7739,  // Golden Bag
		12861, // Sand Jug
		1998,  // Konoha Backpack
		1999,  // Suna Backpack
		2002,  // Kumo Backpack
		1994,  // Yuki Backpack
		1993,  // Hoshi Backpack
		2003,  // Iwa Backpack
		2001,  // Oto Backpack
		1996,  // Marble Backpack
		8922,  // Toad Pouch
		2365,  // Akatsuki Backpack
		2000,  // Sealed Backpack
		12046, // Emperor Backpack
		12079, // Mecha Backpack
		12254, // Akuma Backpack
		12807, // Kara Backpack
	},
	ItemCategoryHelmets: {
		2664,  // Goggles
		3972,  // Bandit Mask
		2490,  // Reinforced Protector
		2479,  // Straw Hat
		2498,  // Ninja Mask
		2170,  // Assassin Turban
		2450,  // Cyborg Goggles
		2395,  // Skeleton Helmet
		12315, // Iwa Mask
		12853, // Suna Mask
		2516,  // Vampire Mask
		11971, // Sound Mask
		2652,  // Cursed Protector
		6537,  // Obito Goggles
		2475,  // Blue ANBU Mask
		2523,  // Red ANBU Mask
		2515,  // Tobi Mask
		7462,  // Samurai Helmet
		11972, // Elite Samurai Helmet
		12854, // Enforcer Mask
		2356,  // Red Headband
		11419, // Ultimate ANBU Mask
		2471,  // Ninja Helmet
		2409,  // Akatsuki Hat
		11591, // Yuki Cap
		8929,  // Crimson Mask
		7448,  // Golden Helmet
		11408, // Kagero Shawl
		11990, // Rikudou Bandana
		7432,  // Black Samurai Helmet
		2659,  // Hanzo Mask
		11917, // Outcast Mask
		6536,  // Raiton Helmet
		2218,  // Katon Mask
		7454,  // Madara Mask
		11415, // Sentinel's Mask
		11965, // Kara Hood
		2091,  // Vile Protector
		12836, // Dark Helmet
		12045, // Emperor Helmet
		3973,  // Yami Headpiece
		8010,  // Kagero Kage Hat
		12793, // Code Headpiece
		12784, // Boro Chakra Vents
		12084, // Kinshiki Headpiece

	},
	ItemCategoryArmors: {
		2650,  // Gennin Armor
		2653,  // Leather Armor
		2525,  // Brown Coat
		2449,  // Black Coat
		2529,  // Ninja Suit
		2437,  // Assassin Shirt
		2526,  // Haku Coat
		2484,  // Cyborg Armor
		2631,  // Blue Shirt
		2420,  // Hoshi Coat
		12317, // Iwa Armor
		12851, // Suna Armor
		12318, // Hyouton Coat
		2463,  // Skeleton Armor
		2482,  // Vampire Coat
		2656,  // Stranger Coat
		2462,  // Chunnin Armor
		3975,  // Monk Robe
		2465,  // Jiraya Armor
		2486,  // Tsunade Armor
		2487,  // Orochimaru Armor
		2474,  // Temari Armor
		2459,  // Kyokushin Armor
		3969,  // Sound Armor
		3971,  // Cursed Armor
		2510,  // Samurai Armor
		2514,  // Sasuke Shirt
		2655,  // Madara Armor
		6539,  // Hashirama Armor
		11973, // Elite Samurai Armor
		12855, // Enforcer Cloak
		2520,  // Legendary Shirt
		2539,  // Kagero Shirt
		12809, // Nukenin Cape
		2472,  // Akatsuki Suit
		2458,  // Sachiko Armor
		11590, // Light Chakra Armor
		11950, // Gravedigger Robe
		7382,  // Shadow Robe
		11977, // Samurai Kimono
		11442, // Uchiha Coat
		2441,  // Rikudou Armor
		2113,  // Kagero Armor
		2645,  // Mystic Cape
		11911, // Fuguki Armor
		11440, // Senju Armor
		3963,  // Black Samurai Armor
		11892, // Outcast Cloak
		7900,  // Black Armor
		5885,  // Raiton Armor
		7380,  // Katon Armor
		7903,  // Sentinel's Robe
		2139,  // Kagero Robe
		2123,  // Raiton Robe
		2131,  // Kotodama Armor
		11592, // Chakra Armor
		2483,  // Heaven Armor
		12076, // Aoi Armor
		11937, // Unreal Cloak
		2557,  // Sannin Armor
		11966, // Kara Cloak
		11904, // Otsutsuki Armor
		12244, // Ninshu Armor
		12787, // Inner Cape
		2138,  // Vile Robe
		2540,  // Yami Cloak
		12837, // Dark Armor
		12090, // Mecha Armor
		2508,  // Sage Robe
		2436,  // Hyuuga Cloak
		12129, // Infected Armor
		12794, // Code Jacket
		12082, // Momoshiki Robe
		12083, // Kinshiki Armor
	},
	ItemCategoryLegs: {
		2647,  // Shinobi Legs
		2519,  // Bandit Legs
		7451,  // Fat Ninja Legs
		2468,  // Cyborg Legs
		2477,  // Training Legs
		11993, // Skeleton Legs
		2452,  // Hoshi Legs
		12316, // Iwa Legs
		12852, // Suna Legs
		2478,  // White Legs
		2530,  // Vampire Legs
		2469,  // Pharaoh Bandages
		2654,  // Monk Legs
		2648,  // Kyokushin Legs
		2470,  // Cursed Guard Legs
		2504,  // Samurai Legs
		11974, // Elite Samurai Legs
		12856, // Enforcer Legs
		5940,  // Legendary Legs
		11409, // Kagero Shorts
		11946, // Bandit King Pants
		2649,  // Gray Pants
		2426,  // Akatsuki Legs
		7894,  // Sachiko Pants
		11948, // Gravedigger Legs
		11994, // Nadare Legs
		11443, // Uchiha Legs
		11989, // Pierced Legs
		2542,  // Friar Legs
		2447,  // Rikudou Legs
		11410, // Kagero Pants
		12860, // Sand Legs
		11441, // Senju Legs
		11912, // Fuguki Legs
		2460,  // Black Samurai Legs
		11893, // Outcast Legs
		7899,  // Black Legs
		11890, // Raiga Legs
		5903,  // Raiton Legs
		7452,  // Katon Legs
		7902,  // Sentinel's Legs
		11411, // Kagero Legs
		11593, // Chakra Legs
		2485,  // Heaven Legs
		2198,  // Platinum Legs
		12074, // Aoi Legs
		11938, // Unreal Legs
		11412, // Sannin Legs
		11967, // Kara Legs
		11905, // Otsutsuki Legs
		12247, // Ninshu Legs
		12788, // Inner Kilt
		11986, // Yami Legs
		12838, // Dark Legs
		11889, // Kagero Battle Legs
		3967,  // Bounded Leg Piece
		12091, // Mecha Legs
		12131, // Infected Legs
		11884, // Glacier Legs
		12798, // Delta Pants
	},
	ItemCategoryShoes: {
		2396,  // Shinobi Sandals
		2195,  // Getta Shoes
		7428,  // Speed Boots
		3982,  // Leather Boots
		2642,  // Ninja Shoes
		2641,  // Cyborg Boots
		2643,  // Skeleton Shoes
		2644,  // Heavy Boots
		11992, // Hoshi Shoes
		7457,  // Vampire Boots
		2657,  // Kyokushin Shoes
		11991, // Sound Shoes
		2429,  // Samurai Shoes
		12857, // Enforcer Shoes
		2196,  // Kagero Shoes
		11976, // Legendary Boots
		2112,  // Akatsuki Shoes
		11995, // Sachiko Shoes
		11944, // Bandit King Boots
		11908, // Winged Shoes
		11988, // Golden Boots
		7429,  // Impulsed Plate Boots
		11975, // Black Samurai Boots
		2531,  // Fuguki Boots
		11891, // Outcast Shoes
		5901,  // Raiton Boots
		7383,  // Katon Shoes
		11481, // Raikage Boots
		11594, // Reinforced Boots
		11987, // Jounin Boots
		7381,  // Raiga shoes
		2473,  // Heaven Shoes
		12075, // Aoi shoes
		7730,  // Sannin Boots
		11968, // Kara Shoes
		11906, // Otsutsuki Boots
		12248, // Ninshu Boots
		12789, // Inner Boots
		12839, // Dark Boots
		12130, // Infected Shoes
		8009,  // Kagero Kage shoes
		12092, // Mecha Boots
		12795, // Code Boots
		12799, // Delta Boots
		12803, // Boro Boots
		12085, // Momoshiki Shoes
	},
	ItemCategoryShields: {
		2353,  // Yagai Glove
		2457,  // Chakra Amplifier
		11479, // Gunbai
		11478, // Sussano Shield
		7459,  // Sealed Glove
		12301, // Oinin Shield
		11953, // Frozen Chakra
		12802, // Boro Armguard
		11955, // Tal Shield
	},
	ItemCategoryAmulets: {
		2496,  // Konoha Protector
		2665,  // Suna Protector
		2537,  // Oto Protector
		2518,  // Honour Ball
		6579,  // Shinobi Mask
		2481,  // Konoha Defender
		2135,  // Bandit Necklace
		11951, // Undead Amulet
		11978, // Legendary Cloak
		11480, // Orochimaru Earrings
		2528,  // Golden Laurel
		7763,  // Sand Chakra
		7618,  // Chain of Betrayers Teeth
		7433,  // Genjutsu Necklace
		12047, // Samurai Amulet
		11588, // Rokkaku
		2201,  // Sannin Necklace
		8011,  // Red Chakra Amulet
		8013,  // Blue Chakra Amulet
		8012,  // Green Chakra Amulet
		11907, // Otsutsuki Necklace
		11579, // Juubi Amulet
		12081, // Mecha Amulet
		2197,  // Tsunade Amulet
		12797, // Delta Earrings
		2200,  // Jashin Amulet
	},
	ItemCategoryRings: {
		2446,  // Chakra Wings
		2404,  // Sound Belt
		12790, // Inner Ring
		2173,  // Akatsuki Ring
		11452, // Sannin Ring
		12805, // Kara Ring
		12293, // Nukenin Ring
		12136, // Blood Signed
		2179,  // Gold Ring
		7697,  // Ruby Signet
		11952, // Doto Belt
	},
	ItemCategoryScrolls: {
		2164,  // Might Scroll
		2165,  // Scroll of Nature
		2166,  // Power Scroll
		2167,  // Skill Scroll
		2168,  // Chakra Scroll
		2209,  // Healing Scroll
		2169,  // Protection Scroll
		2208,  // Speed Scroll
		2130,  // Scroll of Earth
		2134,  // Scroll of Heaven
		2361,  // Sage Scroll
		11914, // Pure Chakra
		11746, // Support Shinobi
		11741, // Gamakichi
		11742, // Katsuyuu
		11743, // Pakkun
		11744, // Tarantula
		11745, // Urushi
		12041, // Sharingan Spy
		11985, // Ginkaku Soul
		11984, // Kinkaku Soul
		12044, // Denka
		12043, // Kamatari
		12042, // Ningame
		12133, // Blood Blob
		11954, // Chakred Support Shinobi
		12253, // Akuma
		12815, // Hate
		12816, // Terror
		12817, // Destruction
	},
	ItemCategoryGloves: {
		2387, // Bandages

	},
	ItemCategorySwords: {
		2428, // Short Sword
	},
	ItemCategoryDistances: {
		11417, // Training Chain
	},
	ItemCategoryBands: {
		2553, // Brown Band
	},
	ItemCategoryEdibles: {
		2159, // Chakra Pill
	},
	ItemCategoryValuables: {
		2148, // Yen
	},
	ItemCategoryCollectables: {
		12104, // Teddy Bear
	},
}
