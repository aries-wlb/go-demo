package match

var auCategories = []AUCategory{
	{
		Name: "八大除了墨大",
		Conditions: AUCondition{
			BaseCondition: BaseCondition{
				TOEFL: 80,
				IELTS: 6.5,
			},
		},
	},
}

func matchAUCondition(c AUCondition, c_ AUCondition) bool {
	return (c.TOEFL >= c_.TOEFL || c.IELTS >= c_.IELTS) && (c.YBorSeven)
}

func AUMatch(condition AUCondition) string {
	for _, category := range auCategories {
		matched := matchAUCondition(condition, category.Conditions)
		if matched {
			return category.Name
		}
	}
	return ""
}
