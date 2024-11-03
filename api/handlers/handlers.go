package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vincentandreas/dealls/api/middleware"
	"github.com/vincentandreas/dealls/models"
	"github.com/vincentandreas/dealls/repo"
	"github.com/vincentandreas/dealls/utilities"
)

type BaseHandler struct {
	Repo      *repo.Implementation
}

func NewBaseHandler(repo *repo.Implementation) *BaseHandler {
	return &BaseHandler{
		Repo:      repo,
	}
}


func HandleRequests(h *BaseHandler, router *gin.Engine) {
    // router.GET("/users", getUsers)

    router.POST("/register", h.createUser)
    router.POST("/login", h.login)

    authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/feeling", h.insertFeeling)
        // authorized.GET("/recommendation", h.getRecommendation)
    }
}


func (h *BaseHandler) createUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.Repo.SaveUser(newUser,c)
    c.Status(201)
}

func (h *BaseHandler) login(c *gin.Context) {
    var tryLoginUser models.UserLogin
    if err := c.ShouldBindJSON(&tryLoginUser); err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    user, err := h.Repo.Login(tryLoginUser, c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    session := sessions.Default(c)
    session.Set("user_id", user.ID)
    session.Save()

    c.JSON(http.StatusOK, gin.H{"message": "Success"})
}




// func (h *BaseHandler) getRecommendation(c *gin.Context){
//     userId, _ := c.Get("user_id")
//     userIdInt := utilities.ExtractId(userId)
    


//     c.JSON(http.StatusOK, gin.H{"message": "Success"})
// }

func (h  *BaseHandler) insertFeeling(c *gin.Context) {
    userId, _ := c.Get("user_id")
    userIdInt := utilities.ExtractId(userId)

    var feel models.Feeling
    if err := c.ShouldBindJSON(&feel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.Repo.SaveHistory(userIdInt, feel, c)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Success"})
}