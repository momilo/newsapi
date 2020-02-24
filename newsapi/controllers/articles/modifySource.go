package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type modifySourceMsg struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Add  *bool  `json:"add"`
}

// ModifySource handles a request which seeks to add or delete an RSS source in the API, actioning it appropriately
func ModifySource(c *gin.Context) {
	if repo == nil {
		sendInternalError(c)
		return
	}

	msg := &modifySourceMsg{}
	if err := c.BindJSON(msg); err != nil || msg.Add == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody})
		return
	}

	if *msg.Add {
		ok := repo.AddNewSource(msg.Name, msg.Url)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": errAddingSource})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully added a new source. " +
			"Articles from it should become available shortly."})
	} else {
		ok := repo.DelSource(msg.Name)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletingSource})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully deleted a source. " +
			"Please note that - in this api version - this does not affect already cached articles"})
	}
}

// TODO: implement deletion of articles for the source which was removed from the API
