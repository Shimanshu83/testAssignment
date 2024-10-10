package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Post("find-pairs", findPairsHandler)

}

type RequestBody struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Result struct {
	Solutions [][]int `json:"solutions"`
}

func findPairsHandler(c *fiber.Ctx) error {

	var requestBody RequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"err":    "bad request from client",
		})
	}

	if len(requestBody.Numbers) < 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"err":    "not proper input",
		})
	}

	result := findPairs(requestBody.Numbers, requestBody.Target)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"body": Result{
			Solutions: result,
		},
	})

}

// used hashmap to optimized the performance
// now it wll take o(n) time compexity
func findPairs(numbers []int, target int) [][]int {

	var result [][]int

	for i := 0; i < len(numbers); i += 1 {

		for j := i + 1; j < len(numbers); j += 1 {
			if numbers[i]+numbers[j] == target {
				singleResult := []int{i, j}
				result = append(result, singleResult)
			}

		}

	}

	return result
}
