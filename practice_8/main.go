package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Product struct {
	ID             int     `json:"id"`
	SupplierId     int     `json:"supplierId"`
	ManufacturerId int     `json:"manufacturerId"`
	CategoryId     int     `json:"categoryId"`
	ImagesId       int     `json:"imagesId"`
	InternalCode   int     `json:"internalCode"`
	PartNumber     string  `json:"partNumber"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TechnicalInfo  string  `json:"technicalInfo"`
	Price          float64 `json:"price"`
	Status         string  `json:"status"`
}

var products = []Product{
	{ID: 1, SupplierId: 1, ManufacturerId: 1, CategoryId: 1, ImagesId: 1, InternalCode: 439686211, PartNumber: "GEK1337-09BK01", Title: "Гиродин GlobalElectronics 09", Description: "A control moment gyroscope (CMG) is an attitude control device generally used in spacecraft attitude control systems.", TechnicalInfo: "Гиродин(МАХОВИК электродвигателя для космических аппаратов), масса 5кг, момент инерции 3 Н*мс, R-r= 0.095м-0.094м=0,001м", Price: 218000, Status: "В наличии"},
	{ID: 2, SupplierId: 2, ManufacturerId: 2, CategoryId: 2, ImagesId: 2, InternalCode: 400345944, PartNumber: "SMR4558-07KL03", Title: "Сенсор MagnetoResist 4558", Description: "A high-precision magnetoresistive sensor designed for industrial applications.", TechnicalInfo: "Диапазон измерений: 0-100 мТл, Точность: ±0.01 мТл, Рабочая температура: -40°C до 85°C", Price: 32000, Status: "В наличии"},
	{ID: 3, SupplierId: 3, ManufacturerId: 3, CategoryId: 3, ImagesId: 3, InternalCode: 848239501, PartNumber: "OPT7563-03XY09", Title: "Оптический датчик Omnix 7563", Description: "An advanced optical sensor for use in automation and robotics.", TechnicalInfo: "Диапазон чувствительности: 400-700 нм, Разрешение: 12 бит, Интерфейс: I2C", Price: 45000, Status: "В наличии"},
	{ID: 4, SupplierId: 4, ManufacturerId: 4, CategoryId: 4, ImagesId: 4, InternalCode: 836275999, PartNumber: "PWR3021-05MN12", Title: "Блок питания PowerMax 3021", Description: "A reliable power supply unit for various electronic devices.", TechnicalInfo: "Входное напряжение: 100-240V AC, Выходное напряжение: 12V DC, Мощность: 50W", Price: 15000, Status: "В наличии"},
	{ID: 5, SupplierId: 5, ManufacturerId: 5, CategoryId: 5, ImagesId: 5, InternalCode: 500399101, PartNumber: "CAP2225-09ZQ18", Title: "Конденсатор CapX 2225", Description: "High-capacity capacitor for energy storage applications.", TechnicalInfo: "Емкость: 4700 мкФ, Напряжение: 25V, Тип: Электролитический", Price: 200, Status: "В наличии"},
	{ID: 6, SupplierId: 6, ManufacturerId: 6, CategoryId: 6, ImagesId: 6, InternalCode: 211000310, PartNumber: "RES5330-11GH20", Title: "Резистор ResiTech 5330", Description: "Precision resistor for accurate current limiting and voltage division.", TechnicalInfo: "Сопротивление: 10kΩ, Допуск: ±0.1%, Мощность: 0.25W", Price: 50, Status: "В наличии"},
	{ID: 7, SupplierId: 2, ManufacturerId: 2, CategoryId: 2, ImagesId: 7, InternalCode: 529684211, PartNumber: "ECS5000-12AC02", Title: "Электромагнитный компенсатор ECS5000", Description: "Электромагнитный компенсатор используется для подавления электромагнитных помех в сложных системах.", TechnicalInfo: "Масса 2.5кг, рабочая частота 50-60Гц, напряжение 220В, эффективность подавления 98%", Price: 150000, Status: "В наличии"},
	{ID: 8, SupplierId: 8, ManufacturerId: 8, CategoryId: 8, ImagesId: 8, InternalCode: 768954312, PartNumber: "IND-10MH", Title: "Индуктивность IND-10MH", Description: "10mH inductor for filtering and energy storage applications.", TechnicalInfo: "Индуктивность 10 мГн, ток 2 А, сопротивление 50 мОм", Price: 25, Status: "В наличии"},
	{ID: 9, SupplierId: 9, ManufacturerId: 9, CategoryId: 9, ImagesId: 9, InternalCode: 654987213, PartNumber: "MCU-ATMEGA328", Title: "Микроконтроллер MCU-ATMEGA328", Description: "Popular 8-bit microcontroller used in various DIY electronics and embedded systems.", TechnicalInfo: "Тактовая частота 16 МГц, память 32 КБ, число выводов 28", Price: 250, Status: "В наличии"},
	{ID: 10, SupplierId: 10, ManufacturerId: 10, CategoryId: 10, ImagesId: 10, InternalCode: 321654987, PartNumber: "NPTD-LED-1W", Title: "Светодиод LED-1W", Description: "High-brightness 1W LED for lighting and indicator applications.", TechnicalInfo: "Мощность 1 Вт, яркость 100 лм, напряжение 3 В", Price: 30, Status: "В наличии"},
	{ID: 11, SupplierId: 11, ManufacturerId: 11, CategoryId: 11, ImagesId: 11, InternalCode: 214365879, PartNumber: "TRANS-220V12V", Title: "Трансформатор TRANS-220V12V", Description: "Step-down transformer to convert 220V AC to 12V AC.", TechnicalInfo: "Мощность 50 ВА, входное напряжение 220 В, выходное напряжение 12 В", Price: 1200, Status: "В наличии"},
	{ID: 12, SupplierId: 12, ManufacturerId: 12, CategoryId: 12, ImagesId: 12, InternalCode: 987654123, PartNumber: "MOSFET-IRF540N", Title: "MOSFET транзистор IRF540N", Description: "N-channel power MOSFET for switching applications.", TechnicalInfo: "Макс. ток 33 А, макс. напряжение 100 В, сопротивление в открытом состоянии 0.044 Ом", Price: 50, Status: "В наличии"},
	{ID: 13, SupplierId: 13, ManufacturerId: 13, CategoryId: 13, ImagesId: 13, InternalCode: 753951456, PartNumber: "REG-7805", Title: "Стабилизатор напряжения REG-7805", Description: "5V linear voltage regulator for power supply circuits.", TechnicalInfo: "Выходное напряжение 5 В, ток 1 А, входное напряжение 7-35 В", Price: 20, Status: "В наличии"},
	{ID: 14, SupplierId: 14, ManufacturerId: 14, CategoryId: 14, ImagesId: 14, InternalCode: 852369741, PartNumber: "RELAY-12V10A", Title: "Реле RELAY-12V10A", Description: "12V relay capable of switching up to 10A for control applications.", TechnicalInfo: "Управляющее напряжение 12 В, ток переключения 10 А, количество контактов 1", Price: 45, Status: "В наличии"},
	{ID: 15, SupplierId: 15, ManufacturerId: 15, CategoryId: 15, ImagesId: 15, InternalCode: 951357486, PartNumber: "CRYSTAL-16MHZ", Title: "Кварцевый резонатор CRYSTAL-16MHZ", Description: "16MHz crystal oscillator for clock generation in microcontrollers and other devices.", TechnicalInfo: "Частота 16 МГц, точность ±50 ppm, рабочая температура -20°C до 70°C", Price: 8, Status: "В наличии"},
}

func main() {
	router := gin.Default()

	// Получение всех товаров
	router.GET("/products", getProducts)

	// Получение товара по ID
	router.GET("/products/:id", getProductByID)

	// Создание нового товара
	router.POST("/products", createProduct)

	// Обновление существующего товара
	router.PUT("/products/:id", updateProduct)

	// Удаление товара
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range products {
		if strconv.Itoa(item.ID) == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, item := range products {
		if strconv.Itoa(item.ID) == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, item := range products {
		if strconv.Itoa(item.ID) == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
