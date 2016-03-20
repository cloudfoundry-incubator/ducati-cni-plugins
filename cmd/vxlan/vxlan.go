package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/appc/cni/pkg/skel"
	"github.com/appc/cni/pkg/types"
	"github.com/cloudfoundry-incubator/ducati-daemon/client"
	"github.com/cloudfoundry-incubator/ducati-daemon/models"
)

type NetConf struct {
	types.NetConf
	NetworkID     string `json:"network_id"`
	DaemonBaseURL string `json:"daemon_base_url"`
}

func loadConf(bytes []byte) (*NetConf, error) {
	n := &NetConf{}

	if err := json.Unmarshal(bytes, n); err != nil {
		return nil, fmt.Errorf("failed to load netconf: %v", err)
	}

	if n.DaemonBaseURL == "" {
		return nil, fmt.Errorf(`"daemon_base_url" field required.`)
	}

	return n, nil
}

func cmdAdd(args *skel.CmdArgs) error {
	if args.ContainerID == "" {
		return errors.New("CNI_CONTAINERID is required")
	}

	netConf, err := loadConf(args.StdinData)
	if err != nil {
		return fmt.Errorf("loading config: %s", err)
	}

	if netConf.NetworkID == "" {
		return fmt.Errorf(`loading config: "network_id" field required.`)
	}

	daemonClient := client.New(netConf.DaemonBaseURL, http.DefaultClient)

	ipamResult, err := daemonClient.ContainerUp(models.CNIAddPayload{
		Args:               args.Args,
		ContainerNamespace: args.Netns,
		InterfaceName:      args.IfName,
		NetworkID:          netConf.NetworkID,
		ContainerID:        args.ContainerID,
	})
	if err != nil {
		return err
	}

	return ipamResult.Print()
}

func cmdDel(args *skel.CmdArgs) error {
	netConf, err := loadConf(args.StdinData)
	if err != nil {
		return fmt.Errorf("loading config: %s", err)
	}

	daemonClient := client.New(netConf.DaemonBaseURL, http.DefaultClient)

	err = daemonClient.ContainerDown(models.CNIDelPayload{
		ContainerNamespace: args.Netns,
		InterfaceName:      args.IfName,
		ContainerID:        args.ContainerID,
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdDel)
}
