package cli

// notest
import (
	_ "embed"
	"fmt"
	"github.com/NethermindEth/juno/internal/services"
	"github.com/NethermindEth/juno/pkg/db"
	"github.com/NethermindEth/juno/pkg/starknet"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/NethermindEth/juno/internal/config"
	"github.com/NethermindEth/juno/internal/errpkg"
	"github.com/NethermindEth/juno/internal/log"
	"github.com/NethermindEth/juno/internal/process"
	"github.com/NethermindEth/juno/pkg/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cobra configuration.
var (
	// cfgFile is the path of the juno configuration file.
	cfgFile string
	// longMsg is the long message shown in the "juno --help" output.
	//go:embed long.txt
	longMsg string

	processHandler *process.Handler
	// rootCmd is the root command of the application.
	rootCmd = &cobra.Command{
		Use:   "juno [options]",
		Short: "Starknet client implementation in Go.",
		Long:  longMsg,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(longMsg)

			processHandler = process.NewHandler()

			// Handle signal interrupts and exits.
			sig := make(chan os.Signal)
			signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-sig
				log.Default.Info("Trying to close...")
				cleanup()
				os.Exit(0)
			}()

			// Breaking initial cases
			if config.Runtime.Ethereum.Node == "" || config.Runtime.Starknet.FeederGateway == "" {
				log.Default.Panic("Ethereum client needed")
			}

			// Subscribe the RPC client to the main loop if it is enabled in
			// the config.
			if config.Runtime.RPC.Enabled {
				s := rpc.NewServer(":" + strconv.Itoa(config.Runtime.RPC.Port))
				processHandler.Add("RPC", s.ListenAndServe, s.Close)
			}

			// Initialize ABI Service
			abiService := services.NewABIService()
			processHandler.Add("ABI Service", abiService.Run, abiService.Close)

			// Initialize State storage service
			stateService := services.NewStateService()
			processHandler.Add("State Storage Service", stateService.Run, stateService.Close)

			// Initialize Contract Hash storage service
			contractHashService := services.NewContractHashService()
			processHandler.Add("Contract Hash Storage Service", contractHashService.Run, stateService.Close)

			// Subscribe the Starknet Synchronizer to the main loop if it is enabled in
			// the config.
			if config.Runtime.Starknet.Enabled {
				// Layer 1 synchronizer for Ethereum State
				stateSynchronizer := starknet.NewSynchronizer(db.NewKeyValueDb(config.Runtime.DbPath, 0))
				processHandler.Add("Starknet Synchronizer", stateSynchronizer.UpdateState,
					stateSynchronizer.Close)
			}

			// Endless running process
			log.Default.Info("Starting all processes...")
			processHandler.Run()
			cleanup()
		},
	}
)

func cleanup() {
	processHandler.Close()
	log.Default.Info("App closing...Bye!!!")
}

// init defines flags and handles configuration.
func init() {
	// Set the functions to be run when rootCmd.Execute() is called.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf(
		"config file (default is %s)", filepath.Join(config.Dir, "juno.yaml")))
}

// initConfig reads in Config file or environment variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use Config file specified by the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Use the default path for user configuration.
		viper.AddConfigPath(config.Dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("juno1")
	}

	// Check whether the environment variables match any of the existing
	// keys and loads them if they are found.
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Default.Info("Config file not found.")
		config.New()
		err = viper.ReadInConfig()
		errpkg.CheckFatal(err, "Failed to read in Config after generation.")
	}

	// Unmarshal and log runtime config instance.
	err = viper.Unmarshal(&config.Runtime)
	errpkg.CheckFatal(err, "Unable to unmarshal runtime config instance.")
}

// Execute handle flags for Cobra execution.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Default.With("Error", err).Error("Failed to execute CLI.")
	}
}
