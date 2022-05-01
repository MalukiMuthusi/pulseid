package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/store/mysql"
	"github.com/MalukiMuthusi/pulseid/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Create a deadline to wait for.
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*5, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	store, err := mysql.New()
	if err != nil {
		logger.Log.Fatal("failed to initialize database ", err)
	}

	r := handlers.SetUpRouter(store, handlers.DebugPrintRoute)

	port := viper.GetString(utils.Port)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		if err := srv.ListenAndServe(); err != nil {
			logger.Log.WithFields(logrus.Fields{"serverListenError": err}).Info("server listen error")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	c := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	logger.Log.WithField("shuttingDown", "shutdown").Info("shutdown server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	os.Exit(0)
}

func init() {
	viper.AutomaticEnv()

	viper.SetEnvPrefix(utils.AppName)

	BindEnvs()

	viper.SetDefault(utils.Port, "8080")
	viper.SetDefault(utils.DbHost, "127.0.0.1")
	viper.SetDefault(utils.DbHostedOnCloud, false)

	CheckMustBeSetEnvs()
}

func BindEnvs() {
	viper.BindEnv(utils.Port, "PORT")
	viper.BindEnv(utils.DbUser)
	viper.BindEnv(utils.DbPwd)
	viper.BindEnv(utils.DbPort)
	viper.BindEnv(utils.DbName)
	viper.BindEnv(utils.DbHost)
	viper.BindEnv(utils.DbHostedOnCloud)
	viper.BindEnv(utils.DbConnectionName)
	viper.BindEnv(utils.DbTimeZone)
}

func EnvMustBeSet(key string) {
	if !viper.IsSet(key) {
		logger.Log.WithField(key, false).Fatal("not set")
	}
}

func CheckMustBeSetEnvs() {
	EnvMustBeSet(utils.DbUser)
	EnvMustBeSet(utils.DbPwd)
	EnvMustBeSet(utils.DbPort)
	EnvMustBeSet(utils.DbName)
	EnvMustBeSet(utils.DbHostedOnCloud)
	EnvMustBeSet(utils.DbConnectionName)
	EnvMustBeSet(utils.DbTimeZone)
}
