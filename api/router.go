package api

import (
	"github.com/gin-gonic/gin"
	cartApi "go-shop/api/cart"
	categoryApi "go-shop/api/category"
	orderApi "go-shop/api/order"
	productApi "go-shop/api/product"
	userApi "go-shop/api/user"
	"go-shop/config"
	"go-shop/domain/cart"
	"go-shop/domain/category"
	"go-shop/domain/order"
	"go-shop/domain/product"
	"go-shop/domain/user"
	"go-shop/utils/database_handler"
	"go-shop/utils/middleware"
	"log"
)

type Databases struct {
	categoryRepository    *category.Repository
	userRepository        *user.Repository
	productRepository     *product.Repository
	cartRepository        *cart.Repository
	cartItemRepository    *cart.ItemRepository
	orderRepository       *order.Repository
	orderedItemRepository *order.OrderItemRepository
}

var AppConfig = &config.Configuration{}

func CreateDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		return nil
	}
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}
	db := database_handler.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)

	return &Databases{
		categoryRepository:    category.NewCategoryRepository(db),
		cartRepository:        cart.NewCartRepository(db),
		userRepository:        user.NewUserRepository(db),
		productRepository:     product.NewProductRepository(db),
		cartItemRepository:    cart.NewCartItemRepository(db),
		orderRepository:       order.NewOrderRepository(db),
		orderedItemRepository: order.NewOrderItemRepository(db),
	}

}
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	userGroup.POST("", userController.Create)
	userGroup.POST("/login", userController.Login)
}
func RegisterCategoryHandlers(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")
	categoryGroup.POST("", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.Create)
	categoryGroup.GET("", categoryController.GetAll)
	categoryGroup.POST("upload", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.BulkCreate)

}

func RegisterCartHandlers(r *gin.Engine, dbs Databases) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/", cartController.GetCart)
}

func RegisterProductHandlers(r *gin.Engine, dbs Databases) {
	productService := product.NewService(*dbs.productRepository)
	productController := productApi.NewProductController(*productService)
	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Create)
	productGroup.DELETE(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Delete)
	productGroup.PATCH(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Update)

}

func RegisterOrderHandlers(r *gin.Engine, dbs Databases) {
	orderService := order.NewService(
		*dbs.orderRepository, *dbs.orderedItemRepository, *dbs.productRepository, *dbs.cartRepository,
		*dbs.cartItemRepository)
	productController := orderApi.NewOrderController(orderService)
	orderGroup := r.Group("/order", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	orderGroup.POST("", productController.Complete)
	orderGroup.DELETE("", productController.Cancel)
	orderGroup.GET("", productController.GetAll)
}
