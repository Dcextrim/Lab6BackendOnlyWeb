package handlers

import (
	"net/http"

	"series-tracker/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// GetAllSeries obtiene todas las series
// @Summary Listar todas las series
// @Tags series
// @Produce json
// @Success 200 {array} models.Series
// @Router /api/series [get]
func (h *Handler) GetAllSeries(c *gin.Context) {
	var series []models.Series
	result := h.DB.Find(&series)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, series)
}

// CreateSeries crea una nueva serie
// @Summary Crear nueva serie
// @Tags series
// @Accept json
// @Produce json
// @Param series body models.Series true "Datos de la serie"
// @Success 201 {object} models.Series
// @Router /api/series [post]
func (h *Handler) CreateSeries(c *gin.Context) {
	var newSeries models.Series
	if err := c.ShouldBindJSON(&newSeries); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Create(&newSeries)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSeries)
}

// GetSeriesByID obtiene una serie por ID
// @Summary Obtener serie por ID
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id} [get]
func (h *Handler) GetSeriesByID(c *gin.Context) {
	id := c.Param("id")
	var series models.Series

	result := h.DB.First(&series, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	c.JSON(http.StatusOK, series)
}

// UpdateSeries actualiza una serie completa
// @Summary Actualizar serie
// @Tags series
// @Accept json
// @Produce json
// @Param id path int true "ID de la serie"
// @Param series body models.Series true "Datos actualizados"
// @Success 200 {object} models.Series
// @Router /api/series/{id} [put]
func (h *Handler) UpdateSeries(c *gin.Context) {
	id := c.Param("id")
	var series models.Series

	if err := h.DB.First(&series, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}

// DeleteSeries elimina una serie
// @Summary Eliminar serie
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 204
// @Router /api/series/{id} [delete]
func (h *Handler) DeleteSeries(c *gin.Context) {
	id := c.Param("id")
	result := h.DB.Delete(&models.Series{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateStatus actualiza el estado de una serie
// @Summary Actualizar estado
// @Tags series
// @Accept json
// @Produce json
// @Param id path int true "ID de la serie"
// @Param status body string true "Nuevo estado"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/status [patch]
func (h *Handler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct{ Status string }

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	series.Status = input.Status
	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}

// IncrementEpisode incrementa el contador de episodios
// @Summary Incrementar episodio
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/episode [patch]
func (h *Handler) IncrementEpisode(c *gin.Context) {
	id := c.Param("id")
	var series models.Series

	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	if series.LastEpisodeWatched < series.TotalEpisodes {
		series.LastEpisodeWatched++
		h.DB.Save(&series)
	}

	c.JSON(http.StatusOK, series)
}

// Upvote incrementa el ranking
// @Summary Incrementar ranking
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/upvote [patch]
func (h *Handler) Upvote(c *gin.Context) {
	id := c.Param("id")
	var series models.Series

	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	series.Ranking++
	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}

// Downvote decrementa el ranking
// @Summary Decrementar ranking
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/downvote [patch]
func (h *Handler) Downvote(c *gin.Context) {
	id := c.Param("id")
	var series models.Series

	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	if series.Ranking > 0 {
		series.Ranking--
		h.DB.Save(&series)
	}

	c.JSON(http.StatusOK, series)
}
