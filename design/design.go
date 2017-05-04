package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("souls", func() {
	Title("Backup Dark Souls saves")
	Description("Deamon to backup Dark Souls saves.")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("saves", func() {
	Action("save", func() {
		Routing(GET("save"))
		Description("Backups the current save file.")
		Response(OK, "text/plain")
	})
	Action("load", func() {
		Routing(GET("load"))
		Description("Loads the last saved file.")
		Response(OK, "text/plain")
	})
	Action("start", func() {
		Routing(GET("start"))
		Description("Starts automatically saving the game every few seconds.")
		Response(OK, "text/plain")
	})
	Action("stop", func() {
		Routing(GET("stop"))
		Description("Stops automatically saving the game.")
		Response(OK, "text/plain")
	})

})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/index.html")
})
