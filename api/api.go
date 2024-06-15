package api

import (
	"net/http"
	"slagboomapplicatie_code/database"

	"github.com/gin-gonic/gin"
)

// @Summary Check of een kenteken toegang heeft tot de parkeerplaats
// @Description Check of een kenteken toegang heeft tot de parkeerplaats
// @Tags Kenteken
// @Accept json
// @Produce json
// @Param kenteken query string true "Kenteken"
// @Success 200 {object} any
// @Failure 400 {object} any
// @Failure 500 {object} any
// @Router /check [get]
func CheckKenteken(kenteken string, c *gin.Context) {
	if kenteken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "geen kenteken opgegeven",
		})
		return
	}
	found, voornaam, achternaam, telefoonnummer, err := database.CheckKentekenInDatabase(kenteken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "fout bij het ophalen van het kenteken",
		})
		return
	}

	if found {
		c.JSON(http.StatusOK, gin.H{
			"message":        "Kenteken heeft toegang tot de parkeerplaats.",
			"status":         "ok",
			"kenteken":       kenteken,
			"voornaam":       voornaam,
			"achternaam":     achternaam,
			"telefoonnummer": telefoonnummer,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Kenteken heeft geen toegang tot de parkeerplaats.",
			"status":   "nok",
			"kenteken": kenteken,
		})
	}
}
