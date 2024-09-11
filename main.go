package main

import (
	"GenericProject/internal/domain"
	"GenericProject/internal/handler"
	"GenericProject/internal/pkg/generic_injector"
	"GenericProject/internal/repository"
	"GenericProject/internal/service"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var (
	dbHost              = flag.String("dbHost", "localhost", "database host")
	dbUser              = flag.String("dbUser", "g_user", "database user")
	dbPass              = flag.String("dbPass", "12332145", "database password")
	dbName              = flag.String("dbName", "g_db", "database name")
	dbPort              = flag.String("dbPort", "5432", "database port")
	defaultAllowOrigins = flag.String("cors", "http://localhost:3030/", "defaultAllowOrigins value string")
	appPort             = flag.String("appPort", "3030", "application start port")

	Conn              *sqlx.DB
	logFormat         = "[${time}] ${status} - ${latency} ${method} ${path}\n"
	defaultCorsMaxAge = 3600
)

func connect() {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", *dbUser, *dbPass, *dbHost, *dbPort, *dbName))
	if err != nil {
		panic(err)
	}
	if db == nil {
		log.Fatalln("не удалось подключиться к базе данных!")
	}
	Conn = db
}

func closeConnection(connect *sqlx.DB) error {
	return connect.Close()
}

func main() {
	flag.Parse()
	connect()
	generic_injector.GI = generic_injector.NewInjector()
	generic_injector.GI.InjectModels(
		reflect.TypeOf(domain.Card{}),
	)
	//app_mapper.NewMapper()

	app := fiber.New(fiber.Config{Prefork: false, BodyLimit: 16 * 1024 * 1024})
	//middlewares
	app.Use(logger.New(logger.Config{
		Format: logFormat,
		Output: os.Stdout,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     *defaultAllowOrigins,
		AllowCredentials: true,
		MaxAge:           defaultCorsMaxAge,
	}))
	rep := repository.NewRepository(Conn)
	services := service.NewService(rep, Conn)
	service.Services = services
	h := handler.NewHandler(services)
	h.Init(app)

	//start server
	go func() {
		if err := app.Listen(":" + *appPort); err != nil {
			log.Panicf("не удалось запустить инстанс веб-сервера: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c
	fmt.Println("Остановка веб-сервера...")
	_ = app.Shutdown()
	fmt.Println("--> соединение с базой данных закрыто")
	err := Conn.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("--> сервис остановлен")
}
