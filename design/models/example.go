package models

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
	"github.com/KokiKono/go_tutorial/design/media"
)

var _ = StorageGroup("Model", func() {
	Description("MySQL model")
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQL")
		Model("Message", func() {
			Description("Message model")
			RendersTo(media.MessageType)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("message", gorma.String)
		})
	})
})
