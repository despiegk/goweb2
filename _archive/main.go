package main

import (
	"log"
	"net/http"

	"github.com/despiegk/goweb2/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server implements the ServerInterface generated by oapi-codegen
type Server struct{}

// Implement all the generated interface methods
func (s *Server) ListIssues(ctx echo.Context) error {
	// TODO: Implement database integration
	return ctx.JSON(200, []api.Issue{})
}

func (s *Server) CreateIssue(ctx echo.Context) error {
	var issue api.Issue
	if err := ctx.Bind(&issue); err != nil {
		return ctx.JSON(400, map[string]string{"error": err.Error()})
	}
	// TODO: Implement database integration
	return ctx.JSON(201, issue)
}

func (s *Server) GetIssue(ctx echo.Context, id string) error {
	// TODO: Implement database integration
	return ctx.JSON(200, api.Issue{})
}

func (s *Server) UpdateIssue(ctx echo.Context, id string) error {
	var issue api.Issue
	if err := ctx.Bind(&issue); err != nil {
		return ctx.JSON(400, map[string]string{"error": err.Error()})
	}
	// TODO: Implement database integration
	return ctx.JSON(200, issue)
}

func (s *Server) DeleteIssue(ctx echo.Context, id string) error {
	// TODO: Implement database integration
	return ctx.NoContent(204)
}

func (s *Server) AddTimeEntry(ctx echo.Context, id string) error {
	var timeEntry api.TimeEntry
	if err := ctx.Bind(&timeEntry); err != nil {
		return ctx.JSON(400, map[string]string{"error": err.Error()})
	}
	// TODO: Implement database integration
	return ctx.JSON(201, timeEntry)
}

const swaggerUI = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>GoWeb2 API Documentation</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js"></script>
    <script>
        window.onload = () => {
            window.ui = SwaggerUIBundle({
                url: '/openapi.yaml',
                dom_id: '#swagger-ui',
                deepLinking: true,
            });
        };
    </script>
</body>
</html>
`

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Serve OpenAPI spec
	e.File("/openapi.yaml", "api/openapi.yaml")

	// Serve Swagger UI
	e.GET("/docs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swaggerUI)
	})

	// Create an instance of our server
	server := &Server{}

	// Register our server implementation with Echo
	api.RegisterHandlers(e, server)

	// Start server
	if err := e.Start(":8081"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
