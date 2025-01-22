package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ozlemugur/go-clean-arch-tt/internal/entity"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

type messageRoutes struct {
	t usecase.Messager
	l logger.Interface
}

func newMessageRoutes(handler *gin.RouterGroup, t usecase.Messager, l logger.Interface) {
	r := &messageRoutes{t, l}

	h := handler.Group("/messages")
	{
		h.GET("/sent", r.GetSentMessages)
		h.POST("", r.InsertMessage)
		//	h.POST("/automatic-message-sender", r.ControlMessageScheduler)
	}
}

type messageResponse struct {
	Messages []entity.Message `json:"Messages"`
	Count    int              `json:"count"`           // Number of messages in the response
	Status   string           `json:"status"`          // Status of the response (e.g., "success", "empty")
	Error    string           `json:"error,omitempty"` // Error message, optional field
}

// @Summary     Retrieve sent messages
// @Description Retrieve all messages with status "sent" from the database
// @ID          get-sent-messages
// @Tags        messages
// @Accept      json
// @Produce     json
// @Success     200 {array} messageResponse
// @Failure     500 {object} response
// @Router      /messages/sent [get]
func (r *messageRoutes) GetSentMessages(c *gin.Context) {
	messages, err := r.t.GetSentMessages(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - GetSentMessages")
		errorResponse(c, http.StatusInternalServerError, "Failed to retrieve messages")
		return
	}
	response := messageResponse{
		Messages: messages,
		Count:    len(messages),
		Status:   "success",
	}
	if len(messages) == 0 {
		response.Status = "empty"
		response.Messages = []entity.Message{}
	}
	c.JSON(http.StatusOK, response)
}

// @Summary     Insert a message
// @Description Add a new message to the database
// @ID          insert-message
// @Tags        messages
// @Accept      json
// @Produce     json
// @Param       message body entity.Message true "Message to insert"
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /messages [post]
func (r *messageRoutes) InsertMessage(c *gin.Context) {
	var msg entity.Message

	// Parse the incoming JSON payload
	if err := c.ShouldBindJSON(&msg); err != nil {
		// Respond with a bad request error if the JSON is invalid
		r.l.Error(err, "http - v1 - InsertMessage - invalid input")
		errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Insert the message using the use case layer
	if err := r.t.InsertMessage(c.Request.Context(), msg); err != nil {
		// Log the error and respond with an internal server error
		r.l.Error(err, "http - v1 - InsertMessage - use case error")
		errorResponse(c, http.StatusInternalServerError, "Failed to insert message")
		return
	}
	c.JSON(http.StatusOK, nil)
}
