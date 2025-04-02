package models

import (
	"gorm.io/gorm"
)

// Series representa una serie de TV dentro del sistema de seguimiento de series.
type Series struct {
	gorm.Model                // Incluye automáticamente los campos created_at, updated_at y deleted_at.
	Id                 int    `gorm:"primaryKey" json:"id"`  // Identificador único de la serie.
	Title              string `gorm:"not null" json:"title"` // Nombre de la serie (obligatorio).
	Status             string `json:"status"`                // Estado actual de la serie.
	LastEpisodeWatched int    `json:"lastEpisodeWatched"`    // Último episodio visto por el usuario.
	TotalEpisodes      int    `json:"totalEpisodes"`         // Número total de episodios de la serie.
	Ranking            int    `json:"ranking"`               // Puntuación de la serie según la valoración del usuario.
}
