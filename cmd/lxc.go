package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thirdwavellc/go-proxmox/proxmox"
)

func init() {
	rootCmd.AddCommand(lxcCmd)
}

var lxcCmd = &cobra.Command{
	Use:   "lxc",
	Short: "manage lxc containers",
	Long:  "manage lxc containers",
	Run: func(cmd *cobra.Command, args []string) {
		request := &proxmox.ContainerRequest{
			// TODO: use flag here
			Node: "pve01",
		}
		containers, err := client.GetContainers(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(containers)
	},
}

var lxcListCmd = &cobra.Command{
	Use:   "lxc",
	Short: "list lxc containers",
	Long:  "list lxc containers",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
