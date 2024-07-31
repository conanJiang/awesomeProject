package main

func main() {
	//服务启动
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	//字符转码
	//r := gin.Default()
	//
	//r.GET("/someJSON", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"lang": "GO语言",
	//		"tag":  "<br>",
	//	}
	//
	//	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	//	c.AsciiJSON(http.StatusOK, data)
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8080")

	//HTML渲染
	//router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	////router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	//router.Run(":8080")

	//使用不同目录下名称相同的模板
	//router := gin.Default()
	//router.LoadHTMLGlob("templates/**/*")
	//router.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	//		"title": "Posts",
	//	})
	//})
	//router.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	//		"title": "Users",
	//	})
	//})
	//router.Run(":8080")

	//pure JSON
	//r := gin.Default()
	//// 提供 unicode 实体
	//r.GET("/json", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"html": "<b>Hello, world!</b>",
	//	})
	//})
	//
	//// 提供字面字符
	//r.GET("/purejson", func(c *gin.Context) {
	//	c.PureJSON(200, gin.H{
	//		"html": "<b>Hello, world!</b>",
	//	})
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8080")

	//Query 和 post form
	//POST /post?id=1234&page=1 HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//
	//name=manu&message=this_is_great
	//router := gin.Default()
	//
	//router.POST("/post", func(c *gin.Context) {
	//
	//	id := c.Query("id")
	//	page := c.DefaultQuery("page", "0")
	//	name := c.PostForm("name")
	//	message := c.PostForm("message")
	//
	//	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	//})
	//router.Run(":8080")

	//SecureJSON
	//r := gin.Default()
	//
	//// 你也可以使用自己的 SecureJSON 前缀
	//// r.SecureJsonPrefix(")]}',\n")
	//
	//r.GET("/someJSON", func(c *gin.Context) {
	//	names := []string{"lena", "austin", "foo"}
	//
	//	// 将输出：while(1);["lena","austin","foo"]
	//	c.SecureJSON(http.StatusOK, names)
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8080")

	//XML/JSON/YAML/ProtoBuf 渲染
	//r := gin.Default()
	//
	//// gin.H 是 map[string]interface{} 的一种快捷方式
	//r.GET("/someJSON", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//r.GET("/moreJSON", func(c *gin.Context) {
	//	// 你也可以使用一个结构体
	//	var msg struct {
	//		Name    string `json:"user"`
	//		Message string
	//		Number  int
	//	}
	//	msg.Name = "Lena"
	//	msg.Message = "hey"
	//	msg.Number = 123
	//	// 注意 msg.Name 在 JSON 中变成了 "user"
	//	// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
	//	c.JSON(http.StatusOK, msg)
	//})
	//
	//r.GET("/someXML", func(c *gin.Context) {
	//	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//r.GET("/someYAML", func(c *gin.Context) {
	//	c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//})
	//
	//r.GET("/someProtoBuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//	// protobuf 的具体定义写在 testdata/protoexample 文件中。
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	// 请注意，数据在响应中变为二进制数据
	//	// 将输出被 protoexample.Test protobuf 序列化了的数据
	//	c.ProtoBuf(http.StatusOK, data)
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8080")

	//安全页眉
	//r := gin.Default()
	//
	//expectedHost := "localhost:8080"
	//
	//// Setup Security Headers
	//r.Use(func(c *gin.Context) {
	//	if c.Request.Host != expectedHost {
	//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
	//		return
	//	}
	//	c.Header("X-Frame-Options", "DENY")
	//	c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
	//	c.Header("X-XSS-Protection", "1; mode=block")
	//	c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	//	c.Header("Referrer-Policy", "strict-origin")
	//	c.Header("X-Content-Type-Options", "nosniff")
	//	c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	//	c.Next()
	//})
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.Run() // listen and serve on 0.0.0.0:8080

	//绑定 HTML 复选框
	//r := gin.Default()
	//
	//r.LoadHTMLGlob("views/*")
	//r.GET("/", indexHandler)
	//r.POST("/", formHandler)
	//
	//r.Run(":8080")

	//绑定 Uri
	//route := gin.Default()
	//route.GET("/:name/:id", func(c *gin.Context) {
	//	var person Person
	//	if err := c.ShouldBindUri(&person); err != nil {
	//		c.JSON(400, gin.H{"msg": err.Error()})
	//		return
	//	}
	//	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	//})
	//route.Run(":8080")

	//绑定表单数据至自定义结构体
	//r := gin.Default()
	//r.GET("/getb", GetDataB)
	//r.GET("/getc", GetDataC)
	//r.GET("/getd", GetDataD)
	//
	//r.Run(":8080")

	//绑定查询字符串或表单数据
	//route := gin.Default()
	//route.GET("/testing", startPage)
	//route.Run(":8080")

	//查询字符串参数
	//router := gin.Default()
	//
	//// 使用现有的基础请求对象解析查询字符串参数。
	//// 示例 URL： /welcome?firstname=Jane&lastname=Doe
	//router.GET("/welcome", func(c *gin.Context) {
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式
	//
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})
	//router.Run(":8080")

	//从 reader 读取数据
	//router := gin.Default()
	//router.GET("/someDataFromReader", func(c *gin.Context) {
	//	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	//	if err != nil || response.StatusCode != http.StatusOK {
	//		c.Status(http.StatusServiceUnavailable)
	//		return
	//	}
	//
	//	reader := response.Body
	//	contentLength := response.ContentLength
	//	contentType := response.Header.Get("Content-Type")
	//
	//	extraHeaders := map[string]string{
	//		"Content-Disposition": `attachment; filename="gopher.png"`,
	//	}
	//
	//	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//})
	//router.Run(":8080")

	//定义路由日志的格式
	//r := gin.Default()
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	//
	//r.POST("/foo", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, "foo")
	//})
	//
	//r.GET("/bar", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, "bar")
	//})
	//
	//r.GET("/status", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, "ok")
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run()

	//
}

//type Person struct {
//	Name     string    `form:"name"`
//	Address  string    `form:"address"`
//	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
//}
//
//func startPage(c *gin.Context) {
//	var person Person
//	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
//	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
//	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
//	if c.ShouldBind(&person) == nil {
//		log.Println(person.Name)
//		log.Println(person.Address)
//		log.Println(person.Birthday)
//	}
//
//	c.String(200, "Success")
//}

//type StructA struct {
//	FieldA string `form:"field_a"`
//}
//
//type StructB struct {
//	NestedStruct StructA
//	FieldB       string `form:"field_b"`
//}
//
//type StructC struct {
//	NestedStructPointer *StructA
//	FieldC              string `form:"field_c"`
//}
//
//type StructD struct {
//	NestedAnonyStruct struct {
//		FieldX string `form:"field_x"`
//	}
//	FieldD string `form:"field_d"`
//}
//
//func GetDataB(c *gin.Context) {
//	var b StructB
//	c.Bind(&b)
//	c.JSON(200, gin.H{
//		"a": b.NestedStruct,
//		"b": b.FieldB,
//	})
//}
//
//func GetDataC(c *gin.Context) {
//	var b StructC
//	c.Bind(&b)
//	c.JSON(200, gin.H{
//		"a": b.NestedStructPointer,
//		"c": b.FieldC,
//	})
//}
//
//func GetDataD(c *gin.Context) {
//	var b StructD
//	c.Bind(&b)
//	c.JSON(200, gin.H{
//		"x": b.NestedAnonyStruct,
//		"d": b.FieldD,
//	})
//}
//
//type Person struct {
//	ID   string `uri:"id" binding:"required,uuid"`
//	Name string `uri:"name" binding:"required"`
//}
//
//func indexHandler(c *gin.Context) {
//	c.HTML(200, "form.html", nil)
//}
//
//func formHandler(c *gin.Context) {
//	var fakeForm myForm
//	c.Bind(&fakeForm)
//	c.JSON(200, gin.H{"color": fakeForm.Colors})
//}
//
//type myForm struct {
//	Colors []string `form:"colors[]"`
//}
