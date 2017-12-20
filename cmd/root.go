package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thirdwavellc/go-proxmox/proxmox"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-proxmox",
	Short: "go-proxomx is a cli for managing Proxmox clusters",
	Long:  "go-proxmox is a cli for managing Proxmox clusters",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var (
	client      proxmox.ProxmoxClient
	config      string
	host        string
	user        string
	password    string
	defaultNode string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "host for proxmox cluster (e.g. https://proxmox.example.com)")
	rootCmd.PersistentFlags().StringVar(&user, "user", "", "proxmox user for authentication")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "proxmox user password for authentication")
	rootCmd.PersistentFlags().StringVar(&defaultNode, "default-node", "", "default proxmox node to use if necessary")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("defaultNode", rootCmd.PersistentFlags().Lookup("default-node"))

	// TODO: these don't seem to be set...
	client.Host = host
	client.User = user
	client.Password = password
	ticketReq := &proxmox.TicketRequest{
		Username: client.User,
		Password: client.Password,
	}
	var err error
	client.Auth, err = client.GetAuth(ticketReq)
	if err != nil {
		proxmox.PrintError(err)
	}
}

func initConfig() {
	if config != "" {
		viper.SetConfigFile(config)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if c, err := rootCmd.ExecuteC(); err != nil {
		c.Println(err)
		os.Exit(1)
	}
}
