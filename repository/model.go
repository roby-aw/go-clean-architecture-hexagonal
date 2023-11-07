package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rolename    string             `bson:"rolename,omitempty" binding:"required" json:"rolename"`
	Rolelabel   string             `bson:"rolelabel,omitempty" binding:"required" json:"rolelabel"`
	Description string             `bson:"description,omitempty" binding:"required" json:"description"`
}

type RegisterUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
}

type FilterQuery bson.M

func NewFilterQuery() FilterQuery {
	return FilterQuery(bson.M{})
}

func (q FilterQuery) SetDeleteAtExists(exist bool) FilterQuery {
	q["deleted_at"] = bson.M{"$exists": exist}
	return q
}

func (q FilterQuery) SetRole() FilterQuery {
	q["$lookup"] = bson.M{
		"from":         "role",
		"localField":   "role_id",
		"foreignField": "id",
		"as":           "role",
	}
	return q
}

func (q FilterQuery) SetID(id primitive.ObjectID) FilterQuery {
	q["_id"] = id
	return q
}

func (q FilterQuery) SetEmail(email string) FilterQuery {
	q["email"] = email
	return q
}

func (q FilterQuery) SetDeleteAt() FilterQuery {
	q["deleted_at"] = time.Now()
	return q
}

func (q FilterQuery) SetUpdateAt() FilterQuery {
	q["updated_at"] = time.Now()
	return q
}

func (q FilterQuery) SetUpdate(data interface{}) FilterQuery {
	q["$set"] = data
	return q
}

func (q FilterQuery) SetSortDescending(field string) FilterQuery {
	q["$sort"] = bson.M{field: -1}
	return q
}

func (q FilterQuery) SetSortAscending(field string) FilterQuery {
	q["$sort"] = bson.M{field: 1}
	return q
}

func (q FilterQuery) SetRoleID(id int) FilterQuery {
	q["role_id"] = id
	return q
}

func (q FilterQuery) SetPassword(password string) FilterQuery {
	q["password"] = password
	return q
}

func (q FilterQuery) SetFullname(fullname string) FilterQuery {
	q["fullname"] = fullname
	return q
}
