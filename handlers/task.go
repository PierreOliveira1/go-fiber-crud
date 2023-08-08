package handlers

import (
	"context"
	"log"
	"time"

	"go-fiber-crud/database"
	"go-fiber-crud/models"
	"go-fiber-crud/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")

	var tasks []models.Task
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func GetTaskByID(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")

	id := c.Params("id")
	var task models.Task

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(task)
}

func CompleteTask(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")

	id := c.Params("id")
	var updateData struct {
		Completed bool `json:"completed"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		var error struct {
			message string
		}

		error.message = "Falha"
		return c.Status(400).JSON(error)
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"completed": updateData.Completed}})

	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateTask(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")
	validate := validator.New()
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(task); err != nil {
		return utils.HandleValidationErrors(c, err)
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")

	id := c.Params("id")
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return err
	}

	task.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": task})
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteTask(c *fiber.Ctx) error {
	collection := database.Database.Collection("tasks")

	id := c.Params("id")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
