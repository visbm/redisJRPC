package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"redisjrpc/internal/database/mongodb"
	"redisjrpc/internal/repository/models"
	"redisjrpc/pkg/logger"
	"time"
)

type mongoRepository struct {
	client *mongodb.MongoDB
	logger logger.Logger
}

func NewArticleMongoRepository(cl *mongodb.MongoDB, logger logger.Logger) *mongoRepository {
	return &mongoRepository{
		client: cl,
		logger: logger,
	}
}

func MapArticlesMongoDoc (a models.Article) bson.M {
	return bson.M{
		"_id": a.ID,
		"url": a.URL,
		"title": a.Title,
	}
}

func MapMongoDocToArticle (a bson.M) models.Article {
	return models.Article{
		ID: a["_id"].(string),
		URL: a["url"].(string),
		Title: a["title"].(string),
	}
}

func (r *mongoRepository) GetArticle(id string) (models.Article, error) {
	collection := r.client.Cl.Database("education").Collection("articles")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	var resp bson.M
	err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resp)
	if err != nil {
		r.logger.Errorf("Error getting article: %v", err)
		return models.Article{}, err
	}
	article := MapMongoDocToArticle(resp)
	return article, nil
}

func (r mongoRepository) GetArticles() ([]models.Article, error) {
	collection := r.client.Cl.Database("education").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var article []models.Article

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		r.logger.Errorf("Error getting mongo cursor: %v", err)
		return article, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {		
		var resp bson.M
		if err := cursor.Decode(&resp); err != nil {
			r.logger.Errorf("Error getting article: %v", err)	
			continue 		
		}
		a := MapMongoDocToArticle(resp)
		article = append(article, a)
	}
	err = cursor.Err()
	if err != nil {
		r.logger.Errorf("Error mongo cursor: %v", err)
		return article, err
	}

	return article, nil
}



func (r mongoRepository) CreateArticle(a models.Article) (models.Article, error) {
	collection := r.client.Cl.Database("education").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()	

	doc := MapArticlesMongoDoc(a)

	response, err := collection.InsertOne(ctx, doc)
	if err != nil {
		r.logger.Errorf("Error creating article:  %v", err)
		return models.Article{}, err
	}
	
	article := models.Article{
		ID: response.InsertedID.(string),
		URL: a.URL,
		Title: a.Title,
	}

	return article, nil
}
