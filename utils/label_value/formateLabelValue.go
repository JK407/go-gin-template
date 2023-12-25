package label_value

type CommonLabelForIntValue struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type CommonLabelFoStrValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func FormatLabelForIntValue(common []CommonLabelForIntValue, value int) interface{} {
	for _, m := range common {
		if m.Value == value {
			return m
		}
	}
	return nil
}

func FormatLabelForStringValue(common []CommonLabelFoStrValue, value string) interface{} {
	for _, m := range common {
		if m.Value == value {
			return m
		}
	}
	return nil
}
