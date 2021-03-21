package main

import (
    "io"
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.GET("/", hello)
    e.GET("/hello/:id", helloPathParam)
    e.GET("/hello", helloQueryParam)
    e.POST("/form", form)
    e.POST("/form-data", formData)
    e.POST("/bind", bind)

    // read file for path "/static/image.png"
    e.Static("/static", "image.png")

    e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

func helloPathParam(c echo.Context) error {
    id := c.Param("id")

    return c.String(http.StatusOK, "Hello, World!, id: "+id)
}

func helloQueryParam(c echo.Context) error {
    name := c.QueryParam("name")

    return c.String(http.StatusOK, "hello, Mr."+name)
}

// application/x-www-form-urlencoded
func form(c echo.Context) error {
    name := c.FormValue("name")

    return c.String(http.StatusOK, "this is form, name = "+name)
}

// multipart/form-data
func formData(c echo.Context) error {
    avatar, err := c.FormFile("avatar")
    if err != nil {
        return err
    }

    src, err := avatar.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    // Destination
    dst, err := os.Create(avatar.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()

    // Copy
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    return c.HTML(http.StatusOK, "<b>Tank you!</b>")
}

type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// json, xml, form, query bind to Go struct from Content-Type
func bind(c echo.Context) error {
    user := new(User)
    if err := c.Bind(user); err != nil {
        return err
    }

    return c.JSON(http.StatusOK, user)
}