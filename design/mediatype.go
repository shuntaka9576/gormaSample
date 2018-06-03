package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() { // API defines the microservice endpoint and
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("bottle", func() { // Resources group related API endpoints
	// bottleリソースパスを指定する
	BasePath("/bottles") // together. They map to REST resources for REST
	// Mediaを代入
	DefaultMedia(BottleMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		// 説明
		Description("Get bottle by id") // with its path, parameters (both path
		// localhost:8080/bottles/1とか？
		Routing(GET("/:bottleID")) // parameters and querystring values) and payload
		Params(func() {            // (shape of the request body).
			Param("bottleID", Integer, "Bottle ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var _ = Resource("account", func() { // Resources group related API endpoints
	// bottleリソースパスを指定する
	BasePath("/accounts") // together. They map to REST resources for REST
	// Mediaを代入
	DefaultMedia(Account) // services.

	Action("show", func() { // Actions define a single API endpoint together
		// 説明
		Description("id") // with its path, parameters (both path
		// localhost:8080/bottles/1とか？
		Routing(GET("/:id")) // parameters and querystring values) and payload
		Params(func() {      // (shape of the request body).
			Param("id", Integer, "User ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// レスポンスデータの定義
var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	// 説明
	Description("A bottle of wine")
	// どのような値があるのか
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		// レスポンスに必要な要素
		Required("id", "href", "name")
	})
	// 返すレスポンスのフォーマット
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

var Account = MediaType("application/vnd.account+json", func() {
	Description("celler account")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("name", String, "名前", func() {
			Example("山田　太郎")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("example@gmail.com")
		})
		Required("id", "name", "email")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
})

var _ = Resource("swaggerui", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/")
})