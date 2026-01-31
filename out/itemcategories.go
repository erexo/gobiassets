package out

type ItemCategory int

const (
	ItemCategoryContainers ItemCategory = iota
	ItemCategoryHelmets
	ItemCategoryArmors
	ItemCategoryLegs
	ItemCategoryShoes
	ItemCategoryAmulets
	ItemCategoryShields
	ItemCategoryGloves
	ItemCategoryRings
	ItemCategoryScrolls
	ItemCategorySwords
	ItemCategoryDistance
	ItemCategoryShooters
	ItemCategoryAmmunition
	ItemCategoryExercise
	ItemCategoryTrainers
	ItemCategoryPills
	ItemCategoryFood
	ItemCategoryCurrency
	ItemCategoryNindoCoins
	ItemCategoryEnchanting
	ItemCategoryValuables
	ItemCategoryDolls
	ItemCategoryMissions

	ItemCategoryFirst = ItemCategoryContainers
	ItemCategoryLast  = ItemCategoryMissions
)

type ItemRole uint8

const (
	ItemRoleAll ItemRole = iota
	ItemRoleNinjutsu
	ItemRoleWeapons
	ItemRoleDefense
)

type ItemType uint8

const (
	ItemTypeNone ItemType = iota
	ItemTypeBoss
	ItemTypeMission
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

func ItemCategoryPrefix() string {
	return `//go:generate enumer -type=ItemCategory -trimprefix ItemCategory -output itemcategory_string.go
//go:generate enumer -type=ItemRole -trimprefix ItemRole -output itemrole_string.go

type ItemCategory int

const (
	ItemCategoryContainers ItemCategory = iota
	ItemCategoryHelmets
	ItemCategoryArmors
	ItemCategoryLegs
	ItemCategoryShoes
	ItemCategoryAmulets
	ItemCategoryShields
	ItemCategoryGloves
	ItemCategoryRings
	ItemCategoryScrolls
	ItemCategorySwords
	ItemCategoryDistance
	ItemCategoryShooters
	ItemCategoryAmmunition
	ItemCategoryExercise
	ItemCategoryTrainers
	ItemCategoryPills
	ItemCategoryFood
	ItemCategoryCurrency
	ItemCategoryNindoCoins
	ItemCategoryEnchanting
	ItemCategoryValuables
	ItemCategoryDolls
	ItemCategoryMissions

	ItemCategoryFirst = ItemCategoryContainers
	ItemCategoryLast  = ItemCategoryMissions
)

type ItemRole uint8

const (
	ItemRoleAll ItemRole = iota
	ItemRoleNinjutsu
	ItemRoleWeapons
	ItemRoleDefense
)

type ItemType uint8

const (
	ItemTypeNone ItemType = iota
	ItemTypeBoss
	ItemTypeMission
)`
}

var Items = map[ItemCategory][]uint16{
	ItemCategoryContainers: {
		1987,  // Bag
		1991,  // Konoha Bag
		1992,  // Suna Bag
		1995,  // Kumo Bag
		8019,  // Shopping Bag
		12861, // Sand Jug
		1988,  // Backpack
		1998,  // Konoha Backpack
		1999,  // Suna Backpack
		2002,  // Kumo Backpack
		1994,  // Yuki Backpack
		1993,  // Hoshi Backpack
		2003,  // Iwa Backpack
		2001,  // Oto Backpack
		2004,  // Kiri Backpack
		1996,  // Marble Backpack
		8922,  // Toad Pouch
		2365,  // Akatsuki Backpack
		2000,  // Sealed Backpack
		12936, // Uchiha Backpack
		12046, // Emperor Backpack
		12079, // Mecha Backpack
		12254, // Akuma Backpack
		12807, // Kara Backpack
		14201, // Shinju Backpack
		13965, // Togedama Backpack
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
		7462,  // Samurai Helmet
		12854, // Enforcer Mask
		11972, // Elite Samurai Helmet
		2356,  // Red Headband
		2471,  // Ninja Shawl
		2409,  // Akatsuki Hat
		11591, // Yuki Cap
		8929,  // Crimson Mask
		7448,  // Golden Helmet
		11408, // Kagero Shawl
		11990, // Rikudou Bandana
		2515,  // Tobi Mask
		11419, // ANBU Mask
		12918, // Shinigami Mask
		7432,  // Black Samurai Helmet
		14278, // Iron Helmet
		11917, // Outcast Mask
		6536,  // Raiton Helmet
		2218,  // Katon Mask
		2659,  // Hanzo Mask
		11415, // Sentinel's Mask
		11965, // Kara Hood
		2091,  // Vile Protector
		12836, // Dark Helmet
		13581, // Puppet Mask
		14292, // Yokai Helmet
		13938, // Tsumeato Headpiece
		12045, // Emperor Helmet
		12937, // Sharingan Implant
		3973,  // Yami Headpiece
		8010,  // Kagero Kage Hat
		12784, // Boro Chakra Vents
		12793, // Code Headpiece
		12986, // Toneri Headband
		13945, // Ryu Headband
		13941, // Matsuri Hat
		12084, // Kinshiki Headpiece
		13953, // Hidari Hood
		13962, // Jura Hood
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
		2459,  // Kyokushin Armor
		3975,  // Monk Robe
		3969,  // Sound Armor
		3971,  // Cursed Armor
		2510,  // Samurai Armor
		12855, // Enforcer Cloak
		11973, // Elite Samurai Armor
		2539,  // Kagero Shirt
		2520,  // Legendary Shirt
		12809, // Nukenin Cape
		2472,  // Akatsuki Suit
		2458,  // Sachiko Armor
		11590, // Light Chakra Armor
		11950, // Gravedigger Robe
		7382,  // Shadow Robe
		11977, // Samurai Kimono
		11442, // Uchiha Coat
		12923, // Training Tracksuit
		11440, // Senju Armor
		2113,  // Kagero Armor
		2441,  // Rikudou Armor
		2645,  // Mystic Cape
		11911, // Fuguki Armor
		3963,  // Black Samurai Armor
		14279, // Iron Armor
		5885,  // Raiton Armor
		11892, // Outcast Cloak
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
		7900,  // Black Armor
		11966, // Kara Cloak
		11904, // Otsutsuki Armor
		12244, // Ninshu Armor
		12787, // Inner Cape
		2138,  // Vile Robe
		2540,  // Yami Cloak
		12837, // Dark Armor
		13582, // Puppet Cloak
		14293, // Yokai Armor
		13939, // Tsumeato Armor
		12090, // Mecha Armor
		2508,  // Sage Robe
		2436,  // Hyuuga Cloak
		12129, // Infected Armor
		12945, // Jigen Robe
		12794, // Code Jacket
		12987, // Toneri Robe
		13942, // Matsuri Cloak
		13949, // Mamushi Armor
		12082, // Momoshiki Robe
		12083, // Kinshiki Armor
		13954, // Hidari Cloak
		13963, // Jura Cloak
	},
	ItemCategoryLegs: {
		2647,  // Shinobi Legs
		2519,  // Bandit Legs
		7451,  // Fat Ninja Legs
		2468,  // Cyborg Legs
		11993, // Skeleton Legs
		2452,  // Hoshi Legs
		12316, // Iwa Legs
		12852, // Suna Legs
		2478,  // White Legs
		2530,  // Vampire Legs
		2654,  // Monk Legs
		2648,  // Kyokushin Legs
		2470,  // Cursed Guard Legs
		2504,  // Samurai Legs
		12856, // Enforcer Legs
		11974, // Elite Samurai Legs
		11409, // Kagero Shorts
		5940,  // Legendary Legs
		11946, // Bandit King Pants
		2649,  // Gray Pants
		2426,  // Akatsuki Legs
		7894,  // Sachiko Pants
		11948, // Gravedigger Legs
		11994, // Nadare Legs
		11443, // Uchiha Legs
		12924, // Training Pants
		11441, // Senju Legs
		11989, // Pierced Legs
		2542,  // Friar Legs
		2447,  // Rikudou Legs
		11410, // Kagero Pants
		12860, // Sand Legs
		11912, // Fuguki Legs
		2460,  // Black Samurai Legs
		14280, // Iron Pants
		11893, // Outcast Legs
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
		7899,  // Black Legs
		11967, // Kara Legs
		11905, // Otsutsuki Legs
		12247, // Ninshu Legs
		12788, // Inner Kilt
		11986, // Yami Legs
		12838, // Dark Legs
		13583, // Puppet Leg Piece
		14294, // Yokai Legs
		11889, // Kagero Battle Legs
		3967,  // Bounded Leg Piece
		12091, // Mecha Legs
		12131, // Infected Legs
		11884, // Glacier Legs
		12798, // Delta Pants
		12988, // Toneri Legs
		12992, // Energy Legs
		14188, // Mamushi Pants
		13946, // Ryu Leg Piece
		13955, // Hidari Legs
		13964, // Jura Legs
	},
	ItemCategoryShoes: {
		2396,  // Shinobi Sandals
		2195,  // Getta Shoes
		7428,  // Speed Boots
		3982,  // Leather Boots
		2642,  // Ninja Shoes
		2641,  // Cyborg Boots
		2643,  // Skeleton Shoes
		11992, // Hoshi Shoes
		2644,  // Heavy Boots
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
		12922, // Uchiha Shoes
		12925, // Training Sandals
		12920, // Senju Boots
		11975, // Black Samurai Boots
		14281, // Iron Boots
		2531,  // Fuguki Boots
		11891, // Outcast Shoes
		5907,  // Raiton Boots
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
		13584, // Puppet Shoes
		12130, // Infected Shoes
		14295, // Yokai Boots
		8009,  // Kagero Kage shoes
		12092, // Mecha Boots
		12795, // Code Boots
		12799, // Delta Boots
		12803, // Boro Boots
		12989, // Toneri Shoes
		14189, // Ryu Shoes
		13950, // Mamushi Boots
		12085, // Momoshiki Shoes
		13956, // Hidari Shoes
	},
	ItemCategoryAmulets: {
		2496,  // Konoha Protector
		2665,  // Suna Protector
		2537,  // Oto Protector
		2135,  // Bandit Necklace
		11951, // Undead Amulet
		11978, // Legendary Cloak
		11480, // Orochimaru Earrings
		2528,  // Golden Laurel
		7763,  // Sand Chakra
		7618,  // Chain of Betrayers Teeth
		12985, // Crystal Amulet
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
		12797, // Delta Earrings
		2200,  // Jashin Amulet
		12946, // Jigen Necklace
		12991, // Toneri Pendant
		2197,  // Tsunade Amulet
		14200, // Shinju Amulet
		13943, // Matsuri Scarf
		13951, // Mamushi Chain
		13957, // Hidari Chain
	},
	ItemCategoryShields: {
		12921, // Chakra Codex
		2353,  // Yagai Glove
		12919, // Retractable Shield
		12984, // Crystal Shield
		2457,  // Chakra Amplifier
		11479, // Gunbai
		11478, // Sussano Shield
		7459,  // Sealed Glove
		12301, // Oinin Shield
		12939, // Jagged Shield
		11953, // Frozen Chakra
		12802, // Boro Armguard
		12944, // Jigen Armguard
		12990, // Toneri Gaze
		11955, // Tal Shield
		13947, // Matsuri Armguard
		14187, // Ryu Shield
		13958, // Hidari Shield
	},
	ItemCategoryGloves: {
		2387,  // Bandages
		2394,  // Leather Gloves
		2402,  // Metal Gloves
		2443,  // Dark Gloves
		2412,  // Ninja Gloves
		2172,  // Assassin Gloves
		2390,  // Cyborg Gloves
		2446,  // Hoshi Gloves
		12314, // Iwa Gloves
		14259, // Suna Gloves
		2425,  // Red Gloves
		2421,  // Hyouton Gloves
		2391,  // Vampire Gloves
		2442,  // Bloody Gloves
		2121,  // Kyokushin Gloves
		2400,  // Gummy Gloves
		2407,  // Sound Gloves
		2424,  // Yellow Gloves
		2411,  // Samurai Gloves
		12858, // Enforcer Gloves
		2444,  // Legendary Gloves
		2422,  // Nukenin Gloves
		11945, // Bandit King Gloves
		14271, // Akatsuki Gloves
		2497,  // Reinforced Gloves
		8923,  // Fur Claws
		2494,  // Cursed Gloves
		14272, // Golden Gloves
		2440,  // Chakra Spikes
		14223, // Uchiha Cuffs
		14224, // Training Weights
		14225, // Senju Gloves
		11949, // Undead Gloves
		11957, // Sand Gloves
		14273, // Black Gloves
		14277, // Iron Gloves
		7385,  // Kagero Gloves
		2509,  // Heaven Glove
		// 14265, // Raiton Gloves
		// 14266, // Katon Gloves
		// 14284, // Doton Gloves
		12982, // Crystal Gloves
		12246, // Ninshu Gloves
		11958, // Demonic Gloves
		12786, // Inner Gloves
		12834, // Ravage Gloves
		13936, // Tsumeato Gloves
		14270, // Emperor Gloves
		2629,  // Kokinjo Gloves
		14276, // Shin Gloves
		12078, // Shallow Gloves
		12257, // Akuma Gloves
		12297, // Oinin Gloves
		11416, // Uprising Gloves
		14285, // Celestial Gloves
		14262, // Celestial Gloves [nin]
		14263, // Celestial Gloves [wep]
		14264, // Celestial Gloves [def]
		12796, // Code Kama
		2423,  // Shiver Gloves
		14269, // Energy Gloves
		14268, // Mamushi Gloves
		13940, // Matsuri Gloves
		14267, // Ryu Gloves
		13960, // Togedama Gloves

		///

		// 14260, // idk2 gloves

		///

		// 2534,  // Knuckle Duster
		// 2499,  // Black Glove
		// 2505,  // Raiton Master Gloves
		// 7379,  // Katon Glove
		// 11392, // Doton Hakai Glove
		// 11393, // Fuuton Hakai Glove
		// 11394, // Suiton Hakai Glove
		// 11395, // Raiton Hakai Glove
		// 11396, // Katon Hakai Glove
		// 2633,  // Sealed Basher
		// 2492,  // Heavy Spiked Glove
		// 11883, // Emperor Staff
		// 12941, // Celestial Glove
		// 12086, // Vital Knuckle Dusters
	},
	ItemCategoryRings: {
		2404,  // Sound Belt
		12791, // Inner Ring
		11914, // Pure Chakra
		2173,  // Akatsuki Ring
		2174,  // Akatsuki Ring
		11452, // Sannin Ring
		11453, // Sannin Ring
		12805, // Kara Ring
		12806, // Kara Ring
		12293, // Nukenin Ring
		12294, // Nukenin Ring
		12136, // Blood Signed
		2179,  // Gold Ring
		7697,  // Ruby Signet
		12938, // Uchiha Ring
		11952, // Doto Belt
	},
	ItemCategoryScrolls: {
		12874, // Onbu
		11741, // Gamakichi
		11742, // Katsuyuu
		11745, // Urushi
		11744, // Tarantula
		11743, // Pakkun
		12041, // Sharingan Spy
		13059, // Crystal
		11746, // Support Shinobi
		11985, // Ginkaku Soul
		11984, // Kinkaku Soul
		12932, // Shin
		12044, // Denka
		12043, // Kamatari
		12042, // Ningame
		12133, // Blood Blob
		13060, // Soul
		11954, // Chakred Support Shinobi
		12253, // Akuma
		12815, // Hate
		12816, // Terror
		12817, // Destruction
		14190, // Ichibi
		14191, // Nibi
		14192, // Sanbi
		14193, // Yonbi
		14194, // Gobi
		14195, // Rokubi
		14196, // Nanabi
		14197, // Hachibi
		14198, // Kyuubi
		14199, // Juubi
		13966, // Tsumeaka
	},
	ItemCategorySwords: {
		2428,  // Short Sword
		2419,  // Bandit Sword
		2403,  // Sword
		2386,  // Machete
		14275, // Explosive Sword
		2376,  // Katana
		2438,  // Assassin Katanas
		2434,  // Cyborg Katana
		14254, // Grand Basher
		14255, // Shibuki
		2431,  // Heavy Big Sword
		12313, // Shinobi Katana
		2383,  // Skeleton Sword
		7386,  // Vampire Sword
		2435,  // Kyokushin Sword
		7419,  // Bloody Cleaver
		2417,  // Mystic Katana
		2388,  // Madman Katana
		2406,  // Serpent Sword
		2451,  // Samurai Katana
		12859, // Enforcer Katana
		7391,  // Legendary Sword
		2379,  // Kusanagi
		7420,  // Mystic Sword
		12073, // Lightning Chakra Sword
		11943, // Bandit King Dirk
		2432,  // Triple-Blade Scythe
		2430,  // Shiny Sword
		2181,  // Doton Tweak Katana
		2185,  // Fuuton Tweak Katana
		2182,  // Suiton Tweak Katana
		2183,  // Raiton Tweak Katana
		2187,  // Katon Tweak Katana
		2488,  // Cursed Katana
		8924,  // White Katana
		2433,  // Raiton Swords
		11947, // Undead Sword
		2489,  // Crystal Katana
		14282, // Iron Katana
		11909, // Magma Sword
		6540,  // Raiton Katana
		2627,  // Katon Sword
		2630,  // Kagero Sword
		2461,  // Heaven Sword
		12981, // Crystal Sword
		7438,  // Samehada
		14261, // Kuroi Broadsword
		2632,  // Sealed Hatchet
		12088, // Vital Halberd
		7434,  // Royal Katana
		12245, // Ninshu Katana
		12785, // Inner Katana
		12835, // Ravage Sword
		13952, // Mamushi Cleaver
		14289, // Yokai Katana
		13937, // Tsumeato Sword
		2382,  // Raiga Katana
		12255, // Akuma Katana
		11913, // Emperor Sword
		11885, // Blossom Katana
		2385,  // Unreal Sword
		12296, // Oinin Katana
		8927,  // Ruby Sword
		12801, // Boro Dagger
		12942, // Celestial Sword
		2190,  // Hiramekarei
		11910, // Shadow Dagger
		11956, // Chakred Sword
		14274, // Matsuri Katar
		12087, // Vital Katana
		13961, // Togedama Sword
	},
	ItemCategoryDistance: {
		11417, // Training Chain
		7368,  // Kunai with Note
		2399,  // Shuriken
		7366,  // Sand Shuriken
		2410,  // Kunai
		2389,  // Demonwing Shuriken
		7367,  // Clone Kunais
		11469, // Elite Chain
		2143,  // Explosive Throwing Ball
		11398, // Doton Twisted Kunai
		11399, // Fuuton Twisted Kunai
		11400, // Suiton Twisted Kunai
		11401, // Raiton Twisted Kunai
		11402, // Katon Twisted Kunai
		12983, // Crystal Kunai
		1294,  // Heavy Throwing Ball
		7460,  // Yondaime Kunai
		12295, // Oinin Shuriken
		2157,  // Reinforced Kunai
		2501,  // Bashosen
		11936, // Unreal Blade
		12800, // Delta Blaster
		11894, // Outcast Blade
		13944, // Ryu Blaster
		2144,  // Raiton Shuriken
	},
	ItemCategoryShooters: {
		12948, // Red Thrower
		12949, // Blue Thrower
		12950, // Green Thrower
		2455,  // Senbon Shooter
		2456,  // Shooting Umbrella
		12077, // Aoi Umbrella
		12256, // Akuma Curse
		11960, // Shinobi Gauntlet
		12943, // Celestial Gauntlet
		12993, // Energy Gauntlet
		12994, // Vital Gauntlet
		13959, // Togedama Gauntlet
	},
	ItemCategoryAmmunition: {
		2546,  // Burst Senbon
		12792, // Inner Senbon
		11963, // Sealed Raiton Scroll
		14256, // Sealed Explosive Notes
		2543,  // Senbons
		2545,  // Poison Senbon
		11962, // Sealed Suiton Scroll
		11964, // Sealed Doton Scroll
		8614,  // Silver Senbon
		12249, // Ninshu Senbon
		11961, // Sealed Katon Scroll
	},
	ItemCategoryExercise: {
		12910, // Exercise Scroll
		14238, // Durable Exercise Scroll
		14246, // Lasting Exercise Scroll
		14235, // Exercise Seal
		14239, // Durable Exercise Seal
		14247, // Lasting Exercise Seal
		12912, // Exercise Sword
		14240, // Durable Exercise Sword
		14248, // Lasting Exercise Sword
		12911, // Exercise Glove
		14241, // Durable Exercise Glove
		14249, // Lasting Exercise Glove
		14236, // Exercise Armguard
		14242, // Durable Exercise Armguard
		14250, // Lasting Exercise Armguard
		12913, // Exercise Kunai
		14243, // Durable Exercise Kunai
		14251, // Lasting Exercise Kunai
		14237, // Exercise Weights
		14244, // Durable Exercise Weights
		14252, // Lasting Exercise Weights
		12914, // Exercise Note
		14245, // Durable Exercise Note
		14253, // Lasting Exercise Note
	},
	ItemCategoryTrainers: {
		12108, // Chunnin Training Kit
		12109, // Jounin Training Kit
		12110, // Swordman Training Kit
		12111, // ANBU Training Kit
		12112, // Akatsuki Training Kit
		12113, // Kage Training Kit
		12114, // Kendo Training Kit
		12115, // Shaolin Training Kit
		12901, // Personal Locker Kit
		12980, // Hireling Kit
	},
	ItemCategoryPills: {
		2673,  // Medic Pill
		2159,  // Chakra Pill
		2111,  // Soldier Pill
		8931,  // Katon Soldier Pill
		8944,  // Suiton Soldier Pill
		8936,  // Raiton Soldier Pill
		8940,  // Fuuton Soldier Pill
		8932,  // Doton Soldier Pill
		13072, // Soul Pill
		13562, // Moon Cupcake
	},
	ItemCategoryFood: {
		2666,  // Meat
		2671,  // Ham
		2689,  // Onigiri
		2695,  // Egg
		2687,  // Cookie
		8111,  // Exploding Cookie
		2362,  // Chips
		13571, // Carrot
		13586, // Grapes
		2672,  // Ramen
		6278,  // Cake
		7966,  // Birthday Cake
		6394,  // Cream Cake
		14286, // Golden Ramen
	},
	ItemCategoryCurrency: {
		2148, // Yen
		2152, // Green Yen Note
		2160, // Red Yen Note
		2685, // Crystal Yen Note
	},
	ItemCategoryNindoCoins: {
		12899, // Nindo Coins
	},
	ItemCategoryEnchanting: {
		12906, // Piece of Chakra
		11747, // Katon Stone
		11749, // Suiton Stone
		11750, // Fuuton Stone
		11751, // Raiton Stone
		11748, // Doton Stone
		12258, // Cursed Stone
		12908, // Chakra Dust
		12134, // Disenchanting Device
		12909, // Disenchanting Mallet
	},
	ItemCategoryValuables: {
		7759,  // Sharingan Eye
		13572, // Byakugan Eye
		11525, // Frozen Gem
		12140, // Unstable Device
		12299, // Toad Statue
		6533,  // Ceremonial Book
		12080, // Mecha Ticket
		7765,  // Weak Sand Chakra
		13564, // Cupcake Recipe
		12300, // Oinin Money
		7437,  // Strange Key
		12132, // Blood Cell
		13069, // Mutable Token
		13070, // Unstable Token
		13071, // Volatile Token
		12842, // Erexo's Eye
		12843, // Erexo's Horn
		12844, // Erexo's Brain
		14183, // Transmutation Cone
		13923, // Flower Petal
		8003,  // Katon Chakra
		8002,  // Suiton Chakra
		8006,  // Raiton Chakra
		8004,  // Fuuton Chakra
		8005,  // Doton Chakra
	},
	ItemCategoryDolls: {
		12875, // Bandit King Doll
		12876, // Gravedigger Doll
		12877, // Fuguki Doll
		12878, // Raiga Doll
		12879, // Ginkaku Doll
		12880, // Kinkaku Doll
		12881, // Aoi Doll
		12882, // Outcast Nemesis Doll
		12883, // Kagero Kage Doll
		12884, // Doto Doll
		12885, // Delphic Kernel Doll
		12886, // Kinshiki Doll
		12887, // Momoshiki Doll
		12888, // Samurai Emperor Doll
		12889, // Mecha Naruto Doll
		12890, // Kageboshi Doll
		12891, // Boro Doll
		12892, // Delta Doll
		12893, // Code Doll
		12978, // Guren Doll
		12957, // Shin Doll
		12958, // Jigen Doll
		12979, // Toneri Doll
		12894, // Chino Doll
		13061, // Dark Lord Doll
		13971, // Matsuri Doll
		13972, // Ryu Doll
		13973, // Mamushi Doll
		13974, // Hidari Doll
		13975, // Jura Doll
	},
	ItemCategoryMissions: {
		12104, // Teddy Bear
		12099, // Deadly Poison
		2547,  // Explosive Note
		11896, // Glass Tube
		12016, // Medic Bandages
		12103, // Kyokushin Costume
		12105, // Sun Medal
		12106, // Chimes of Terror
		2136,  // Undying God Emblem
		11880, // Neurotoxin
		12059, // Ice Shard
		12862, // Golden Dust

		11527, // Green Leaf
		11533, // Shadow Leaf
		11532, // Imperial Leaf
		11505, // Toxic Venom
		11520, // Ancient Stone
		11881, // Thorn Leaf
		11897, // Shadow Fabric

		12267, // Hell Bells
		12268, // Prism Glasses
		12266, // Undead Coin
		12270, // Cursed Scale
		12272, // Fiery Brooch
		12273, // Supply Package
		12277, // Crystal Gem
		12219, // Broken Light Emblem
		12220, // Light Emblem

		11901, // Heart Ring
		12064, // Frozen Stone
		12052, // Kagero Jewel
		12100, // Legendary Icon
		11900, // Jungle Net
		12873, // Akatsuki Piece of Cloth

		11511, // Fertile Soil
		11506, // Muddy Hay
		8930,  // Anti-bug powder
		12291, // Vigor Spore
		12292, // Heaven Matter
		12833, // Concentrated Heaven Matter
		12066, // Iron Key

		1958, // Uchiha Book

		11503, // Bamboo Stick
		1959,  // Darui Report
		11519, // Armor Piece
		11523, // Raiton Heart
		11524, // Katon Heart
		11522, // Aegis of Light

		11509, // Muddy Twig
		11512, // Frost Giant Fur
		11518, // Giant Snake Skull
		11539, // Empty Antidote
		8925,  // Full Antidote

		12057, // Godflower Shard
		12065, // Samurai Jewel
		12017, // Confidential Scrolls

		11935, // Power Shard
		11528, // Prism Book
		8017,  // Delphic Device
		12804, // Kara Emblem
		12907, // Inner Flower

		11521, // Sentinel Gooey
		11895, // Golden Coin
		11916, // Otsutsuki Soul
		8266,  // Otsutsuki Parchment

		12279, // Golden Sand
		12281, // Crimson Blossom
		12282, // Cursed Claw
		12283, // Vile Root
		12252, // Declaration
		12243, // Kageboshi Soul

		12758, // Torment Seed
		12682, // Vial of Gore Blood
		12480, // Heaven Plans
		12840, // Dark Essence

		5909, // White Fur
		5911, // Brown Fur
		5912, // Black Fur
		5910, // Green Fur

		11517, // Katon Feather
		11515, // Suiton Feather
		11514, // Raiton Feather
		11513, // Fuuton Feather
		11516, // Doton Feather
	},
}
