package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

func (u *productHandler) Create(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	// 1. 参数解析
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 参数校验
	product.Username = c.MustGet(gin.AuthUserKey).(string)
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s already exist", product.Name)})
		return
	}
	product.CreatedAt = time.Now()

	// 3. 逻辑处理
	u.products[product.Name] = product
	log.Printf("Register product %s success", product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}

func (u *productHandler) Get(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	product, ok := u.products[c.Param("name")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("can not found product %s", c.Param("name"))})
		return
	}

	// 参数校验
	if c.MustGet(gin.AuthUserKey).(string) != product.Username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// RequestID
	router.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return "test"
		},
	}))

	router.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))

	// 认证
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			// 路由匹配
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get)
		}
	}

	return router
}

func main() {
	var eg errgroup.Group

	// 一进程多端口
	insecureServer := &http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	secureServer := &http.Server{
		Addr:         ":8443",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS("server.pem", "server.key")
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
