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

type autoMessageResponse struct {
	Description string `json:"Description"`
	Status      string `json:"status"`
}

// @Summary     Control the automatic message sender scheduler
// @Description Starts or stops the automatic message sender scheduler based on the provided action ("start" or "stop").
// @ID          control-message-scheduler
// @Tags        Scheduler
// @Accept      json
// @Produce     json
// @Param       action body entity.AutoMessageSender true "Action for the scheduler (start or stop)"
// @Success     200 {object} autoMessageResponse "Operation completed successfully with a success message"
// @Failure     400 {object} autoMessageResponse "Invalid input or unsupported action"
// @Failure     500 {object} autoMessageResponse "Internal server error while starting or stopping the scheduler"
// @Router      /automatic-message-sender [post]
func (r *autoMessageSenderRoutes) ControlMessageScheduler(c *gin.Context) {
	var req entity.AutoMessageSender

	// Parse the incoming JSON payload
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - ControlMessageScheduler - invalid input")
		c.JSON(http.StatusBadRequest, autoMessageResponse{
			Status:      "error",
			Description: "Invalid input",
		})
		return
	}

	switch req.Action {
	case "start":
		// Start the scheduler
		err := r.a.StartAutoMessageSender(c.Request.Context())
		if err != nil {
			r.l.Error(err, "http - v1 - ControlMessageScheduler - start error")
			c.JSON(http.StatusInternalServerError, autoMessageResponse{
				Status:      "error",
				Description: "Failed to start scheduler",
			})
			return
		}
		c.JSON(http.StatusOK, autoMessageResponse{
			Status:      "success",
			Description: "Message scheduler started successfully",
		})

	case "stop":
		// Stop the scheduler
		err := r.a.StopAutoMessageSender(c.Request.Context())
		if err != nil {
			r.l.Error(err, "http - v1 - ControlMessageScheduler - stop error")
			c.JSON(http.StatusInternalServerError, autoMessageResponse{
				Status:      "error",
				Description: "Failed to stop scheduler",
			})
			return
		}
		c.JSON(http.StatusOK, autoMessageResponse{
			Status:      "success",
			Description: "Message scheduler stopped successfully",
		})

	default:
		// Handle invalid actions
		c.JSON(http.StatusBadRequest, autoMessageResponse{
			Status:      "error",
			Description: "Invalid action. Use 'start' or 'stop'",
		})
	}
}
