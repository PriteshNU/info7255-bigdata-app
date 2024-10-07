package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"info7255-bigdata-app/models"
	"info7255-bigdata-app/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PlanHandler struct {
	service services.PlanService
}

func NewPlanHandler(service services.PlanService) *PlanHandler {
	return &PlanHandler{
		service: service,
	}
}

func (ph *PlanHandler) GetPlan(c *gin.Context) {
	objectId, found := c.Params.Get("objectId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "objectId is required"})
		return
	}

	plan, err := ph.service.GetPlan(c, objectId)
	if err != nil {
		log.Printf("Failed to fetch plan with err : %v", err.Error())
		if err.Error() == "key not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	currentETag := generateETag(plan)
	clientETag := strings.TrimSpace(c.GetHeader("If-None-Match"))
	if clientETag == currentETag {
		c.Status(http.StatusNotModified)
		return
	}

	clientETag = strings.TrimSpace(c.GetHeader("If-Match"))
	if clientETag != "" && clientETag != currentETag {
		c.Status(http.StatusPreconditionFailed)
		return
	}

	c.Header("ETag", currentETag)

	c.JSON(http.StatusOK, plan)
}

func (ph *PlanHandler) CreatePlan(c *gin.Context) {
	var planRequest models.Plan

	if err := c.ShouldBindBodyWith(&planRequest, binding.JSON); err != nil {
		log.Printf("Bad request with error : %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid fields in the request"})
		return
	}

	existingPlan, err := ph.service.GetPlan(c, planRequest.ObjectId)
	if err == nil && existingPlan.ObjectId != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Plan already exists"})
		return
	}

	if err := ph.service.CreatePlan(c, planRequest); err != nil {
		log.Printf("Failed to create plan with error : %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plan"})
		return
	}

	eTag := generateETag(planRequest)
	c.Header("ETag", eTag)

	c.JSON(http.StatusCreated, gin.H{"message": "Plan created successfully"})
}

func (ph *PlanHandler) DeletePlan(c *gin.Context) {
	objectId, found := c.Params.Get("objectId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "objectId is required"})
		return
	}

	if err := ph.service.DeletePlan(c, objectId); err != nil {
		log.Printf("Failed to delete plan with err : %v", err.Error())
		if err.Error() == "key not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (ph *PlanHandler) GetAllPlans(c *gin.Context) {
	plans, err := ph.service.GetAllPlans(c)
	if err != nil {
		log.Printf("Failed to fetch all plans with err : %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, plans)
}

func generateETag(plan interface{}) string {
	dataBytes, err := json.Marshal(plan)
	if err != nil {
		log.Printf("Error marshalling plan: %v", err)
		return ""
	}
	return generateSHA1Hash(dataBytes)
}

func generateSHA1Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
