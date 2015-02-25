package handlers

import (
	"github.com/asm-products/landline-api/models"
	"github.com/gin-gonic/gin"
)

type TeamJSON struct {
  Email string `json:"email" binding:"required"`
  EncryptedPassword string `json:"password" binding:"required"`
  SSOSecret string `json:"secret" binding:"required"`
  SSOUrl string `json:"url" binding:"required"`
  Slug string `json:"name" binding:"required"`
}

func TeamsCreate(c *gin.Context) {
  var json TeamJSON

  c.Bind(&json)

	t := &models.Team{
    Email: json.Email,
    EncryptedPassword: json.EncryptedPassword,
    SSOSecret: json.SSOSecret,
    SSOUrl: json.SSOUrl,
    Slug: json.Slug,
  }

  team, err := models.FindOrCreateTeam(t)
  if err != nil {
    panic(err)
  }

  token := GenerateToken(team.Id)

  c.JSON(200, gin.H{"token": token})
}