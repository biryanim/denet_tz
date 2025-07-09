package main

import (
	"context"
	"github.com/biryanim/denet_tz/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

//TODO: запустить main
//TODO: напилить мидлвару для проверки токена
//TODO: написать все ручки
//TODO: придумать как сделать возврат ошибок в ответе на запрос
//TODO: накинуть транзакций
//TODO: неплохо будет написать логгер
//TODO: сделать closer & graceful shutdown
