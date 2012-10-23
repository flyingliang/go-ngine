package samplescenes

import (
	"fmt"
	"runtime"

	util "github.com/go3d/go-util"

	ngine "github.com/go3d/go-ngine/core"
)

var (
	MaxKeyHint = len(keyHints) - 1

	ctl *ngine.TController
	curKeyHint = 0
	keyHints = []string {
		"[F2]  --  Toggle Render Technique",
		"[F3]  --  Toggle Backface Culling",
		"[F4]  --  Toggle Texture Filtering",
		"[F5]  --  Increase Filtering Anisotropy",
		"[W][S]  --  Camera rise / sink",
		"[A][D]  --  Camera strafe left / right",
		"[<][>]  --  Camera turn left / right",
		"[^][v]  --  Camera move forward / backward",
		"[PgUp][PgDown]  --  Camera turn up / down",
		"[Alt][LShift][RShift]  --  Camera move-speed x 0.1 / 10 / 100",
	}
	sec = 0
)

func AssetRootDirPath () string {
	return util.BaseCodePath("go-ngine", "_sampleprogs", "_sharedassets")
}

func NewMaterialFromLocalTextureImageFile (assetRootRelativeFilePath string) *ngine.TMaterial {
	ngine.Core.Textures[assetRootRelativeFilePath] = ngine.NewTextureLoadLocalFile(assetRootRelativeFilePath, false)
	return ngine.NewMaterial(assetRootRelativeFilePath)
}

func NewMaterialFromRemoteTextureImageFile (fileUrl string) *ngine.TMaterial {
	ngine.Core.Textures[fileUrl] = ngine.NewTextureLoadRemoteFile(fileUrl, true)
	return ngine.NewMaterial(fileUrl)
}

func PrintPostLoopSummary () {
	fmt.Printf("Avg. FPS: %v\n", ngine.Stats.FpsOverallAverage())
	if ngine.Stats.TrackGC {
		fmt.Printf("GC: avg=%v max=%v\n", ngine.Stats.GcOverallAverage(), ngine.Stats.GcMaxTime)
	}
}

func SamplesMainFunc (loader func ()) {
	runtime.LockOSThread()
	var err error
	defer ngine.Dispose()

	if err = ngine.Init(ngine.NewOptions(AssetRootDirPath(), 1280, 720, 0, false), "Loading Sample...", SamplesOnSec); err != nil {
		fmt.Printf("ABORT: %v\n", err)
	} else {
		ngine.Stats.TrackGC = true
		loader()
		ngine.Loop.Loop()
		PrintPostLoopSummary()
	}
}

func SamplesOnSec () {
	if sec++; sec == 4 {
		sec = 0
		if curKeyHint++; (curKeyHint > MaxKeyHint) || (curKeyHint >= (len(keyHints))) { curKeyHint = 0 }
	}
	ngine.Windowing.SetTitle(WindowTitle())
}

func WindowTitle () string {
	ctl = ngine.Core.CurCamera.Controller
	return fmt.Sprintf("%v FPS @ %vx%v   |   %s   |   Cam: P=%v D=%v", ngine.Stats.FpsLastSec, ngine.Core.Options.WinWidth(), ngine.Core.Options.WinHeight(), keyHints[curKeyHint], ctl.Pos, ctl.Dir)
}
