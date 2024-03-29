package main

import (
	"context"

	mymodule "github.com/zaporter/viam-timelapse"
	"go.viam.com/utils"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
)

var (
	Version     = "development"
	GitRevision = ""
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("timelapse"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	var versionFields []interface{}
	if Version != "" {
		versionFields = append(versionFields, "version", Version)
	}
	if GitRevision != "" {
		versionFields = append(versionFields, "git_rev", GitRevision)
	}
	if len(versionFields) != 0 {
		logger.Infow("timelapse", versionFields...)
	} else {
		logger.Info("timelapse" + " built from source; version unknown")
	}
	mod, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}
	if err := mod.AddModelFromRegistry(ctx, camera.API, mymodule.Model); err != nil {
		return err
	}
	if err := mod.AddModelFromRegistry(ctx, camera.API, mymodule.ModelFake); err != nil {
		return err
	}

	if err := mod.Start(ctx); err != nil {
		return err
	}
	defer mod.Close(ctx)
	<-ctx.Done()
	return nil
}
