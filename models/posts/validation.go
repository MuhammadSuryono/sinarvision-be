package posts

import "github.com/thedevsaddam/govalidator"

func (param ParameterCreatePost) GetValidationRules() govalidator.MapData {
	return govalidator.MapData{
		"title":    []string{"required", "min:4"},
		"content":  []string{"required", "min:10"},
		"category": []string{"required"},
		"status":   []string{"required"},
	}
}

func (param ParameterCreatePost) GetValidationMessage() govalidator.MapData {
	return govalidator.MapData{
		"title":    []string{"required:Title can't empty", "min:Minimum length 4 characters"},
		"content":  []string{"required:Content can't empty"},
		"category": []string{"required:Category can't empty"},
		"status":   []string{"required:Status can't empty"},
	}
}

func (param ParameterUpdatePost) GetValidationRules() govalidator.MapData {
	return govalidator.MapData{
		"title":    []string{"required", "min:4"},
		"content":  []string{"required", "min:10"},
		"category": []string{"required"},
		"status":   []string{"required"},
	}
}

func (param ParameterUpdatePost) GetValidationMessage() govalidator.MapData {
	return govalidator.MapData{
		"title":    []string{"required:Title can't empty", "min:Minimum length 4 characters"},
		"content":  []string{"required:Content can't empty"},
		"category": []string{"required:Category can't empty"},
		"status":   []string{"required:Status can't empty"},
	}
}
