package users

import (
	"github.com/dompaksi/bookstore_users-api/domain/users"
	"github.com/dompaksi/bookstore_users-api/service"
	"github.com/dompaksi/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	  if err != nil{
	    return
	  }
	  if err:= json.Unmarshal(bytes, &user); err != nil {
	    fmt.Println(err.Error())
	    return
	  }*/

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invaled json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*func FindUser(c *gin.Context)  {
  c.String(http.StatusNotImplemented, "no implemented")
}
*/
