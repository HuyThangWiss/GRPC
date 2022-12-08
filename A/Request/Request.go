package Request

type UpdatePost struct {
	Email     string    `json:"email,omitempty" bson:"email,omitempty"  binding:"required"`
	Code     string    `json:"code,omitempty" bson:"code,omitempty"  binding:"required"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"  binding:"required"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"  binding:"required"`
	User      string    `json:"user,omitempty" bson:"user,omitempty" binding:"required"`
}

