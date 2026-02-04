package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	Name string `bson:"name"`
}

var collection *mongo.Collection

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	collection = client.Database("appdb").Collection("items")

	r := gin.Default()
	r.SetHTMLTemplate(loadTemplates())

	r.GET("/", showPage)
	r.POST("/add", addItem)

	r.Run(":8080")
}

func showPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.M{})
	var items []Item
	cursor.All(ctx, &items)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"items": items,
	})
}

func addItem(c *gin.Context) {
	name := c.PostForm("name")
	if name != "" {
		collection.InsertOne(context.TODO(), bson.M{"name": name})
	}
	c.Redirect(http.StatusFound, "/")
}

func loadTemplates() *template.Template {
	t := template.Must(template.ParseFiles("templates/index.html"))
	return t
}
