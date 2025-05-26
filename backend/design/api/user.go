package design

import (
	. "goa.design/goa/v3/dsl"
)

var ResultPayload = Type("ResultPayload", func() {
	Description("Payload of ResultType")

	Attribute("status", String, "Status result of scrape", func() {
		Example("AVAILABLE, NOT_AVAILABLE or TIMEOUT")
	})
})

var ResultInterface = Type("Result", func() {
	Description("Result type")
	Extend(ResultPayload)

	Attribute("id", String, "Unique result ID")
	Attribute("createdAt", String, "Created at date of scrape")
	Attribute("updateAt", String, "Updated at date of scrape")

	Required("id", "status", "createdAt", "updateAt")
})

var ResultList = ResultType("ResultList", func() {
	Description("List of results")
	Attribute("data", ArrayOf(ResultInterface))
	Attribute("page", Int, "Page number")
	Attribute("limit", Int, "Items per page")
	Attribute("total", Int, "Total number of results")

	Required("data", "page", "limit", "total")
})

var _ = Service("Result", func() {
	Description("Service to collect data about the payment slip")

	Method("listResults", func() {
		Payload(func() {
			Attribute("page", Int, "Page number", func() {
				Minimum(0)
				Default(1)
			})
			Attribute("limit", Int, "Items per page", func() {
				Minimum(0)
				Maximum(100)
				Default(10)
			})
		})

		Result(ResultList)

		HTTP(func() {
			GET("/results")

			Param("page", Int, "Page number", func() {
				Minimum(0)
			})
			Param("limit", Int, "Number of items per page", func() {
				Minimum(0)
				Maximum(100)
			})

			Response(StatusOK)
		})
	})

	Method("createResult", func() {
		Payload(ResultPayload)

		Result(ResultInterface)

		HTTP(func() {
			POST("/results")

			Response(StatusCreated)
		})
	})
})
