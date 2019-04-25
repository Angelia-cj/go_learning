package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
	"time"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	router := gin.Default()

	router.Use()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/redict/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	v1 := router.Group("/v1")

	v1.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "v1 login")
	})

	v2 := router.Group("/v2")

	v2.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "v2 login")
	})

	router.Use(MiddleWare())

	router.GET("/before", MiddleWare(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})

	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for     c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/User/:name/*action",func (c *gin.Context){
		name:= c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK,message)
	})

	router.GET("/welcome2", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for     c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/testing", startPage)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	router.GET("/bookable", getBookable)

	router.Run(":8001")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("before middleware")
	}
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}