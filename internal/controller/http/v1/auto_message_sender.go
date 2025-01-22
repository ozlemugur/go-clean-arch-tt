package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ozlemugur/go-clean-arch-tt/internal/entity"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

type autoMessageSenderRoutes struct {
	a usecase.AutoMessager
	l logger.Interface
}

func newAutoMessageSenderRoutes(handler *gin.RouterGroup, a usecase.AutoMessager, l logger.Interface) {
	r := &autoMessageSenderRoutes{a, l}

	h := handler.Group("/automatic-message-sender")
	{
		h.POST("", r.ControlMessageScheduler)
	}
}

// @Summary     Control automatic message sender scheduler
// @Description Starts or stops the automatic message sending scheduler based on the provided action.
// @ID          control-message-scheduler
// @Tags        scheduler
// @Accept      json
// @Produce     json
// @Param       action body entity.AutoMessageSender true "Scheduler action (start or stop)"
// @Success     200 {object} response "Operation completed successfully"
// @Failure     400 {object} response "Invalid input"
// @Failure     500 {object} response "Internal server error"
// @Router      /automatic-message-sender [post]
func (r *autoMessageSenderRoutes) ControlMessageScheduler(c *gin.Context) {
	var req entity.AutoMessageSender

	// Parse the incoming JSON payload
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - ControlMessageScheduler - invalid input")
		errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	switch req.Action {
	case "start":
		// Start the scheduler
		err := r.a.StartAutoMessageSender(c.Request.Context())
		if err != nil {
			r.l.Error(err, "http - v1 - ControlMessageScheduler - start error")
			errorResponse(c, http.StatusInternalServerError, "Failed to start scheduler")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Message scheduler started successfully"})

	case "stop":
		// Stop the scheduler
		err := r.a.StopAutoMessageSender(c.Request.Context())
		if err != nil {
			r.l.Error(err, "http - v1 - ControlMessageScheduler - stop error")
			errorResponse(c, http.StatusInternalServerError, "Failed to stop scheduler")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Message scheduler stopped successfully"})

	default:
		// Handle invalid actions
		errorResponse(c, http.StatusBadRequest, "Invalid action. Use 'start' or 'stop'")
	}
}
