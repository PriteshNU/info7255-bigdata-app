package services

import (
	"encoding/json"
	"info7255-bigdata-app/models"
	"info7255-bigdata-app/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

type PlanService interface {
	GetPlan(c *gin.Context, key string) (models.Plan, error)
	CreatePlan(c *gin.Context, plan models.Plan) error
	DeletePlan(c *gin.Context, key string) error
	GetAllPlans(ctx *gin.Context) ([]models.Plan, error)
}

type planService struct {
	repo repositories.RedisRepo
}

func NewPlanService(repo repositories.RedisRepo) PlanService {
	return &planService{
		repo: repo,
	}
}

func (ps *planService) GetPlan(ctx *gin.Context, key string) (models.Plan, error) {
	var plan models.Plan

	value, err := ps.repo.Get(ctx, key)
	if err != nil {
		log.Printf("Error getting the plan from redis : %v", err)
		return plan, err
	}

	if err := json.Unmarshal([]byte(value), &plan); err != nil {
		log.Printf("Error unmarshalling the plan from redis : %v", err)
		return plan, err
	}

	return plan, nil
}

func (ps *planService) CreatePlan(ctx *gin.Context, plan models.Plan) error {
	key := plan.ObjectId

	value, err := json.Marshal(plan)
	if err != nil {
		log.Printf("Error marshalling plan: %v", err)
		return err
	}

	err = ps.repo.Set(ctx, key, string(value))
	if err != nil {
		log.Printf("Error saving plan to redis: %v", err)
		return err
	}

	return nil
}

func (ps *planService) DeletePlan(ctx *gin.Context, key string) error {
	_, err := ps.GetPlan(ctx, key)
	if err != nil {
		log.Printf("Error getting the plan from the redis : %v", err)
		return err
	}

	err = ps.repo.Delete(ctx, key)
	if err != nil {
		log.Printf("Error deleting the plan from redis : %v", err)
		return err
	}

	return nil
}

func (ps *planService) GetAllPlans(ctx *gin.Context) ([]models.Plan, error) {
	var plans []models.Plan

	keys, err := ps.repo.Keys(ctx, "*")
	if err != nil {
		log.Printf("Error fetching all keys from redis : %v", err)
		return nil, err
	}

	for _, key := range keys {
		plan, err := ps.GetPlan(ctx, key)
		if err != nil {
			log.Printf("Error fetching the plan from redis : %v", err)
			return nil, err
		}
		plans = append(plans, plan)
	}

	return plans, nil
}
