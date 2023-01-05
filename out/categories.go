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
	ItemCategoryDistance
	ItemCategoryShooters
	ItemCategoryAmmunition
	ItemCategoryBands
	ItemCategoryPills
	ItemCategoryFood
	ItemCategoryCurrency
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
	ItemCategoryDistance
	ItemCategoryShooters
	ItemCategoryAmmunition
	ItemCategoryBands
	ItemCategoryPills
	ItemCategoryFood
	ItemCategoryCurrency
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
)`
}

var Items = map[ItemCategory][]uint16{
	ItemCategoryContainers: {
		1990,  // Present
		6104,  // Jewel Case
		1987,  // Bag
		1991,  // Konoha Bag
		1992,  // Suna Bag
		1995,  // Kumo Bag
		3939,  // Forest Bag
		5927,  // Pirate Bag
		1997,  // Golden Bag
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
		2387,  // Bandages
		2394,  // Leather Gloves
		2402,  // Metal Gloves
		2443,  // Black Gloves
		2412,  // Ninja Gloves
		2172,  // Assassin Gloves
		2390,  // Cyborg Gloves
		2421,  // Hyouton Gloves
		2400,  // Gummy Gloves
		12314, // Shinobi Gloves
		2425,  // Red Gloves
		2391,  // Vampire Gloves
		2121,  // Kyokushin Gloves
		2442,  // Bloody Gloves
		2407,  // Sound Glove
		2411,  // Samurai Gloves
		12858, // Enforcer Glove
		2444,  // Legendary Gloves
		2534,  // Knuckle Duster
		2422,  // Golden Gloves
		11945, // Bandit King Glove
		2439,  // Raiton Gloves
		2497,  // Reinforced Glove
		2494,  // Cursed Glove
		8923,  // Fur Claws
		2440,  // Chakra Spikes
		2499,  // Black Glove
		11949, // Undead Glove
		11391, // Hakai Glove
		2120,  // Razor Claws
		2505,  // Raiton Master Gloves
		7379,  // Katon Glove
		7385,  // Kagero Glove
		11392, // Doton Hakai Glove
		2509,  // Heaven Glove
		11958, // Demonic Glove
		11393, // Fuuton Hakai Glove
		2633,  // Sealed Basher
		12246, // Ninshu Glove
		2492,  // Heavy Spiked Glove
		12786, // Inner Gloves
		12834, // Ravage Glove
		11394, // Suiton Hakai Glove
		2423,  // Shiver Gloves
		11416, // Uprising Gloves
		12078, // Shallow Glove
		12257, // Akuma Glove
		11418, // Shadow Gloves
		12297, // Oinin Gloves
		2629,  // Kokinjo Gloves
		11395, // Raiton Hakai Glove
		12796, // Code Kama
		11883, // Emperor Staff
		11957, // Prism Glove
		2658,  // Sealed Golden Glove
		11396, // Katon Hakai Glove
		12086, // Vital Knuckle Dusters
	},
	ItemCategorySwords: {
		2428,  // Short Sword
		2419,  // Bandit Sword
		7591,  // Jagged Sword
		11477, // Katana
		2386,  // Machete
		2384,  // Katanas
		2438,  // Assassin Katanas
		2434,  // Cyborg Katana
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
		11943, // Bandit King Dirk
		2432,  // Triple-Blade Scythe
		2430,  // Shiny Sword
		2488,  // Cursed Katana
		8924,  // White Katana
		2433,  // Raiton Swords
		11947, // Undead Sword
		2489,  // Crystal Katana
		11909, // Magma Sword
		2188,  // Tweak Katana
		6540,  // Raiton Katana
		2627,  // Katon Sword
		2630,  // Kagero Sword
		2181,  // Doton Tweak Katana
		2461,  // Heaven Sword
		7438,  // Samehada
		2185,  // Fuuton Tweak Katana
		2632,  // Sealed Hatchet
		12245, // Ninshu Katana
		12785, // Inner Katana
		12835, // Ravage Sword
		2182,  // Suiton Tweak Katana
		2382,  // Raiga Katana
		12255, // Akuma Katana
		7434,  // Royal Katana
		11885, // Blossom Katana
		12073, // Lightning Chakra Sword
		12296, // Oinin Katana
		8927,  // Ruby Sword
		11913, // Emperor Sword
		2183,  // Raiton Tweak Katana
		2385,  // Unreal Sword
		2190,  // Hiramekarei
		12801, // Boro Dagger
		11956, // Chakred Sword
		11910, // Shadow Dagger
		6528,  // Searing Katana
		2187,  // Katon Tweak Katana
		12087, // Vital Katana

	},
	ItemCategoryDistance: {
		11417, // Training Chain
		2399,  // Shuriken
		7366,  // Sand Shuriken
		2410,  // Kunai
		7368,  // Kunai with Note
		2389,  // Demonwing Shuriken
		7367,  // Clone Kunais
		11469, // Elite Chain
		11397, // Twisted Kunai
		2157,  // Reinforced Kunai
		11398, // Doton Twisted Kunai
		11399, // Fuuton Twisted Kunai
		7460,  // Yondaime Kunai
		11894, // Outcast Blade
		2143,  // Explosive Throwing Ball
		11400, // Suiton Twisted Kunai
		2144,  // Raiton Shuriken
		12295, // Oinin Shuriken
		1294,  // Heavy Throwing Ball
		2501,  // Bashosen
		11401, // Raiton Twisted Kunai
		11936, // Unreal Blade
		12800, // Delta Blaster
		11402, // Katon Twisted Kunai
	},
	ItemCategoryShooters: {
		2455,  // Senbon Shooter
		2456,  // Shooting Umbrella
		12077, // Aoi Umbrella
		12256, // Akuma Curse
		11960, // Shinobi Gauntlet
	},
	ItemCategoryAmmunition: {
		2546,  // Burst Senbon
		11963, // Sealed Raiton Scroll
		12792, // Inner Senbon
		2543,  // Senbons
		2545,  // Poison Senbon
		8614,  // Silver Senbon
		11962, // Sealed Suiton Scroll
		11964, // Sealed Doton Scroll
		12249, // Ninshu Senbon
		11961, // Sealed Katon Scroll
	},
	ItemCategoryBands: {
		2553,  // Brown Band
		2551,  // Red Band
		2550,  // Blue Band
		2548,  // Yellow Band
		2580,  // Green Band
		2533,  // Gray Band
		11414, // Black Band
		12088, // Vital Halberd
	},
	ItemCategoryPills: {
		2673, // Medic Pill
		2159, // Chakra Pill
		2111, // Soldier Pill
		8931, // Katon Soldier Pill
		8944, // Suiton Soldier Pill
		8936, // Raiton Soldier Pill
		8940, // Fuuton Soldier Pill
		8932, // Doton Soldier Pill
		5900, // Raiton Gem
		7850, // Blue Shard
		8309, // Red Shard
		6569, // Candy
	},
	ItemCategoryFood: {
		2666, // Meat
		2671, // Ham
		2689, // Onigiri
		2695, // Egg
		2687, // Cookie
		8111, // Exploding Cookie
		2362, // Chips
		2672, // Ramen
		6278, // Cake
		7966, // Birthday Cake
	},
	ItemCategoryCurrency: {
		2148, // Yen
		2152, // Green Yen Note
		2160, // Red Yen Note
		2685, // Crystal Yen Note
		1691, // Premium Ticket
	},
	ItemCategoryValuables: {
		5890,  // Adrenaline
		8306,  // Small Piece of Chakra
		8302,  // Piece of Chakra
		8300,  // Perfect Piece of Chakra
		7768,  // Mystery Chakra
		6535,  // Chakred Scroll
		11525, // Frozen Gem
		12140, // Unstable Device
		2500,  // Spring Box
		2445,  // Summer Box
		3970,  // Autumn Box
		2427,  // Winter Box
		2090,  // +20% experience box
		2142,  // +20% Experience Box [Event]
		1964,  // green hokage sign
		1963,  // blue hokage sign
		1965,  // prism hokage sign
		12134, // disenchanting device
		12299, // Toad Statue
		1982,  // Authority Tome
		1983,  // Blessed Tome
		1985,  // Tome of Destruction
		1986,  // Tome of Sins
		7620,  // Kami Skull
		2149,  // Gold Bar
		5670,  // Voodoo Skull
		6533,  // Ceremonial Book
		6534,  // Imperor's Trident
		2745,  // Blue Rose
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
		12894, // Chino Doll
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
