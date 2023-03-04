package model

type Package struct {
	Name  string      `json:"name" firestore:"name"`
	Price float64     `json:"price" firestore:"price"`
	Foods FoodPackage `json:"foods" firestore:"foods"`
}

type FoodPackage struct {
	Monday    string `json:"monday" firestore:"monday"`
	Tuesday   string `json:"tuesday" firestore:"tuesday"`
	Wednesday string `json:"wednesday" firestore:"wednesday"`
	Thursday  string `json:"thursday" firestore:"thursday"`
	Friday    string `json:"friday" firestore:"friday"`
	Saturday  string `json:"saturday" firestore:"saturday"`
	Sunday    string `json:"sunday" firestore:"sunday"`
}

type PackageResponse struct {
	Id    string              `json:"id"`
	Name  string              `json:"name"`
	Price float64             `json:"price"`
	Foods FoodPackageResponse `json:"foods"`
}

type FoodPackageResponse struct {
	Monday    FoodResponse `json:"monday"`
	Tuesday   FoodResponse `json:"tuesday"`
	Wednesday FoodResponse `json:"wednesday"`
	Thursday  FoodResponse `json:"thursday"`
	Friday    FoodResponse `json:"friday"`
	Saturday  FoodResponse `json:"saturday"`
	Sunday    FoodResponse `json:"sunday"`
}

type PackageRequest struct {
	Name  string             `json:"name" binding:"required"`
	Price float64            `json:"price" binding:"required"`
	Foods FoodPackageRequest `json:"foods" binding:"required"`
}

type FoodPackageRequest struct {
	Monday    string `json:"monday" binding:"required"`
	Tuesday   string `json:"tuesday" binding:"required"`
	Wednesday string `json:"wednesday" binding:"required"`
	Thursday  string `json:"thursday" binding:"required"`
	Friday    string `json:"friday" binding:"required"`
	Saturday  string `json:"saturday" binding:"required"`
	Sunday    string `json:"sunday" binding:"required"`
}
