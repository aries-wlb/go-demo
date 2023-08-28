package match

var hkCategories = []HKCategory{
	{
		Name: "Top3",
		Conditions: HKCondition{
			IELTS:  6.5,
			ALevel: aLevelScore.AScore * 3,
		},
	},
	{
		Name: "Other not Top3",
		Conditions: HKCondition{
			IELTS:  6,
			ALevel: aLevelScore.AScore*2 + aLevelScore.BScore*1,
		},
	},
}

func matchHKCondition(c HKCondition, c_ HKCondition) bool {
	return c.IELTS > c_.IELTS && c.ALevel >= c_.ALevel
}

func HKMatch(condition HKCondition) string {
	for _, category := range hkCategories {
		matched := matchHKCondition(condition, category.Conditions)
		if matched {
			return category.Name
		}
	}
	return ""
}
