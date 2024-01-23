package models

type Datalink struct{
	
	Shortlink string `json:"shortlink" gorm:"primary_key"`
	Link string `json:"link" gorm:"unique"`
}

type Postlink struct{
	Link string `json:"link" binding:"required"`
}
