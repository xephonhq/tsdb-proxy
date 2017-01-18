package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xephonhq/tsdb-proxy/pkg/util"
	"os"
)

var Version = "0.0.1-dev"
var log = util.Logger.NewEntryWithPkg("t.cmd")

var (
	configFile        = ""
	defaultConfigFile = "tsdb-proxy.yml"
	debug             = false
)

var RootCmd = &cobra.Command{
	Use:   "tsdb-proxy",
	Short: "Time series database proxy",
	Long:  `TSDB proxy is a proxy for time series databases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TSDB Proxy:" + Version + " Use `tsdb-proxy -h` for more information")
	},
}

func Execute() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "config file")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
}

func initConfig() {
	if debug {
		util.UseVerboseLog()
	}
	viper.AutomaticEnv()
	// TODO: check file existence
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
	} else {
		log.Debugf("config file %s is loaded", configFile)
	}
}
