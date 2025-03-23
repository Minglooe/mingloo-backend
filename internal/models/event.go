package models

import (
	"time"

	"github.com/google/uuid"
)

// Event representa as informações de um evento
// swagger:response EventGetResponse
// swagger:model
type Event struct {
	ID          uuid.UUID `extensions:"x-order=0" json:"id"`          // ID do evento
	Name        string    `extensions:"x-order=1" json:"name"`        // Nome do evento
	Address     Address   `extensions:"x-order=2" json:"address"`     // Endereço completo do evento
	Date        Date      `extensions:"x-order=3" json:"date"`        // Data do evento
	Model       string    `extensions:"x-order=4" json:"model"`       // Modelo do evento (Presencial, Remoto, Híbrido)
	Photo       string    `extensions:"x-order=5" json:"photo"`       // Foto (Banner) do evento
	Description string    `extensions:"x-order=6" json:"description"` // Descrição do evento (markdown ou WYSIWYG)
	Capacity    int       `extensions:"x-order=7" json:"capacity"`    // Lotação máxima do evento
	Type        Type      `extensions:"x-order=8" json:"type"`        // Tipo de ingresso (Livre, Compra, Convite, Híbrido)
}

// Address representa o endereço completo do evento
type Address struct {
	State        string `extensions:"x-order=0" json:"state"`        // Estado
	Neighborhood string `extensions:"x-order=1" json:"neighborhood"` // Bairro
	City         string `extensions:"x-order=2" json:"city"`         // Cidade
	Street       string `extensions:"x-order=3" json:"street"`       // Rua
	Number       string `extensions:"x-order=4" json:"number"`       // Número
	Complement   string `extensions:"x-order=5" json:"complement"`   // Complemento
	ZipCode      string `extensions:"x-order=6" json:"zip_code"`     // CEP
}

// Date representa as datas de início e fim do evento
type Date struct {
	Start time.Time `extensions:"x-order=0" json:"start"` // Data e hora de início
	End   time.Time `extensions:"x-order=1" json:"end"`   // Data e hora de término
}

// Type representa o tipo de ingresso para o evento
type Type struct {
	Label       string `extensions:"x-order=0" json:"label"`       // Tipo de ingresso (Livre, Compra, Convite, Híbrido)
	Description string `extensions:"x-order=1" json:"description"` // Descrição do tipo de ingresso (o que ele implica)
}
