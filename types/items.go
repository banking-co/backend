package types

type ItemRarity = string

const (
	ItemRarityDefault   ItemRarity = "common"
	ItemRarityUncommon  ItemRarity = "uncommon"
	ItemRarityRare      ItemRarity = "rare"
	ItemRarityEpic      ItemRarity = "epic"
	ItemRarityLegendary ItemRarity = "legendary"
)

type ItemType = string

const (
	ItemTypeBusinessStaff   ItemType = "business_staff"
	ItemTypeBusinessUpgrade ItemType = "business_upgrade"
	ItemTypeUserUpgrade     ItemType = "user_upgrade"
)
