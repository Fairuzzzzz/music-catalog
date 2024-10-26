package tracks

import (
	"net/http"

	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req trackactivities.TrackActivitiesRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("userID")
	err := h.service.UpsertTrackActivites(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
