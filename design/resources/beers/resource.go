package beers

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BeerResource
var _ = Resource("beer", func() {
	BasePath("/beers")
	DefaultMedia(BeerMedia)
	Action("list", func() {
		Description("Get all beers")
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(BeerMedia))
		})
		Response(NotFound)
	})
	Action("show", func() {
		Description("Get beer by id")
		Routing(GET("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {
		Description("create new beer")
		Routing(POST(""))
		Payload(BeerPayload)
		Response(Created, func() {
			Media(BeerMedia)
		})
	})
	Action("change", func() {
		Description("change exist beer")
		Routing(PATCH("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Payload(BeerPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("remove", func() {
		Description("remove exist beer")
		Routing(DELETE("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Response(InternalServerError)
})

// TapResource
var _ = Resource("tap", func() {
	BasePath("/taps")
	DefaultMedia(TapMedia)
	Action("list", func() {
		Description("Get all taps")
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(TapMedia))
		})
		Response(NotFound)
	})
	Action("show", func() {
		Description("Get beerInfo by tap id")
		Routing(GET("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {
		Description("create bew tap")
		Routing(POST(""))
		Payload(TapPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("change", func() {
		Description("change beer ID related tap")
		Routing(PATCH("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Payload(TapPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("remove", func() {
		Description("remove exist tap")
		Routing(DELETE("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Response(InternalServerError)
})

var BeerMedia = MediaType("application/vnd.goa.study.beer+json", func() {
	Description("A bottle of beer")
	Attributes(func() {
		Attribute("id", Integer, "beer ID")
		Attribute("title", String, "beer name")
		Attribute("type", String, "beer type")
		Attribute("created_at", String, "create time stamp")
		Attribute("updated_at", String, "update time stamp")
		Required("id", "title", "type")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("type")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var TapMedia = MediaType("application/vnd.goa.study.tap+json", func() {
	Description("A beer info of tap")
	Attributes(func() {
		Attribute("id", Integer, "beer tap ID")
		Attribute("beerID", Integer, "beer name")
		Attribute("beerTitle", String, "beer name")
		Attribute("created_at", String, "create time stamp")
		Attribute("updated_at", String, "update time stamp")
		Required("id", "beerTitle")
	})
	View("default", func() {
		Attribute("id")
		Attribute("beerID")
		Attribute("beerTitle")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var BeerPayload = Type("BeerPayload", func() {
	Member("title", String, "Beer name")
	Member("type", String, "beer type")
	Required("title", "type")
})

var TapPayload = Type("TypePayload", func() {
	Member("beerID", Integer, "beer ID related tap")
	Required("beerID")
})
