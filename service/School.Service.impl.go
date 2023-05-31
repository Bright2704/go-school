package service

import (
	"context"
	"errors"

	"example/BatteryTracking/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchoolServiceImpl struct {
	schoolcollection *mongo.Collection
	ctx				  context.Context
}

func NewSchoolService(schoolcollection *mongo.Collection, ctx context.Context) SchoolService {
	return &SchoolServiceImpl{
		schoolcollection: schoolcollection,
		ctx:			  ctx,
	}
}

func (u *SchoolServiceImpl) CreateSchool(school *entity.UserSchool) error {
	_, err := u.schoolcollection.InsertOne(u.ctx, school)
	return err
}

func (u *SchoolServiceImpl) GetSchool(name *string) (*entity.UserSchool, error) {
	var user *entity.UserSchool
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.schoolcollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *SchoolServiceImpl) GetAll() ([]*entity.UserSchool, error) {
	var users []*entity.UserSchool
	cursor, err := u.schoolcollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
		}
	for cursor.Next(u.ctx) {
		var user entity.UserSchool
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}	

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)
	
	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *SchoolServiceImpl) UpdateSchool(user *entity.UserSchool) error {
	filter := bson.D{primitive.E{Key: "name", Value: user.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name},primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "address", Value: user.Address}}}}
	result, _ := u.schoolcollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *SchoolServiceImpl) DeleteSchool(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := u.schoolcollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
