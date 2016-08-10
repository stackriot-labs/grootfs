package unpacker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/cloudfoundry/gunk/command_runner"

	"code.cloudfoundry.org/grootfs/cloner"
	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter . IDMapper
type IDMapper interface {
	MapUIDs(logger lager.Logger, pid int, mappings []groot.IDMappingSpec) error
	MapGIDs(logger lager.Logger, pid int, mappings []groot.IDMappingSpec) error
}

type NamespacedCmdUnpacker struct {
	commandRunner command_runner.CommandRunner
	idMapper      IDMapper
	unpackCmdName string
}

func NewNamespacedCmdUnpacker(commandRunner command_runner.CommandRunner, idMapper IDMapper, unpackCmdName string) *NamespacedCmdUnpacker {
	return &NamespacedCmdUnpacker{
		commandRunner: commandRunner,
		idMapper:      idMapper,
		unpackCmdName: unpackCmdName,
	}
}

func (u *NamespacedCmdUnpacker) Unpack(logger lager.Logger, spec cloner.UnpackSpec) error {
	logger = logger.Session("unpacked-with-namespaced-cmd", lager.Data{"spec": spec})
	logger.Debug("start")
	defer logger.Debug("end")

	ctrlPipeR, ctrlPipeW, err := os.Pipe()
	if err != nil {
		return fmt.Errorf("creating tar control pipe: %s", err)
	}

	unpackCmd := exec.Command(os.Args[0], u.unpackCmdName, spec.RootFSPath)
	unpackCmd.Stdin = spec.Stream
	if len(spec.UIDMappings) > 0 || len(spec.GIDMappings) > 0 {
		unpackCmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUSER,
		}
	}

	outBuffer := bytes.NewBuffer([]byte{})
	unpackCmd.Stdout = outBuffer
	unpackCmd.Stderr = outBuffer
	unpackCmd.ExtraFiles = []*os.File{ctrlPipeR}

	if err := u.commandRunner.Start(unpackCmd); err != nil {
		return fmt.Errorf("starting unpack command: %s", err)
	}
	logger.Debug("command-is-started")

	if err := u.setIDMappings(logger, spec, unpackCmd.Process.Pid); err != nil {
		ctrlPipeW.Close()
		return err
	}

	if _, err := ctrlPipeW.Write([]byte{0}); err != nil {
		return fmt.Errorf("writing to tar control pipe: %s", err)
	}
	logger.Debug("command-is-signaled-to-continue")

	logger.Debug("waiting-for-command")
	if err := u.commandRunner.Wait(unpackCmd); err != nil {
		return fmt.Errorf(outBuffer.String())
	}

	return nil
}

func (u *NamespacedCmdUnpacker) setIDMappings(logger lager.Logger, spec cloner.UnpackSpec, untarPid int) error {
	if len(spec.UIDMappings) > 0 {
		if err := u.idMapper.MapUIDs(logger, untarPid, spec.UIDMappings); err != nil {
			return fmt.Errorf("setting uid mapping: %s", err)
		}
		logger.Debug("uid-mappings-are-set")
	}

	if len(spec.GIDMappings) > 0 {
		if err := u.idMapper.MapGIDs(logger, untarPid, spec.GIDMappings); err != nil {
			return fmt.Errorf("setting gid mapping: %s", err)
		}
		logger.Debug("gid-mappings-are-set")
	}

	return nil
}
