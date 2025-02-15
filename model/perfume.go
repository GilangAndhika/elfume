package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Perfume struct {
	PerfumeID   primitive.ObjectID `json:"perfume_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Brand       string             `json:"brand" bson:"brand"`
	Types       string             `json:"types" bson:"types"`           // Types of the perfume (e.g. Eau de Parfum, Pure Perfume, etc.)
	Categories  string             `json:"categories" bson:"categories"` // Fragrance categories (e.g. Floral, Fresh, Woody, etc.)
	Sizes       string             `json:"sizes" bson:"sizes"`           // Available sizes of the perfume (e.g. 50ml, 100ml, 200ml, etc.)
	Image       string             `json:"image" bson:"image"`
	Price       string             `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`
	Stock       string             `json:"stock" bson:"stock"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type FumeImgUpload struct {
	PerfumeID primitive.ObjectID `json:"perfume_id" bson:"_id"`
	Image     string             `json:"image" bson:"image"`
	FileName  string             `json:"file_name" bson:"file_name"`
}

type GithubUploadRequest struct {
	Message string `json:"message"`
	Content string `json:"content"`
}
