package handlers

type CreateDataRequest struct {
	Content string `json:"content"`
}

// func CreateData(c echo.Context) error {
// 	req := new(CreateDataRequest)
// 	if err := c.Bind(req); err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid request")
// 	}

// 	data, err := db.CreateExampleData(c.Request().Context(), req.Content)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Failed to create data")
// 	}

// 	return c.JSON(http.StatusOK, data)
// }

// func GetData(c echo.Context) error {
// 	id := c.Param("id")

// 	data, err := db.GetExampleData(c.Request().Context(), id)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "Data not found")
// 	}

// 	return c.JSON(http.StatusOK, data)
// }
