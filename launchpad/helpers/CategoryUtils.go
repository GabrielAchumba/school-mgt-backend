package helpers

import categoryDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/dtos"

type CategoryUtils struct {
}

func NewCategoryUtils() CategoryUtils {
	return CategoryUtils{}
}

func (impl CategoryUtils) FindCategorys(categories []categoryDTOPackage.Category, parentId string) []categoryDTOPackage.Category {

	var _categories []categoryDTOPackage.Category

	for _, value := range categories {
		if value.ParentId == parentId {
			_categories = append(_categories, value)
		}
	}
	return _categories
}
