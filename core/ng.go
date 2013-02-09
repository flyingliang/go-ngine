package core

import (
	"fmt"

	ugl "github.com/go3d/go-opengl/util"
)

//	Call this to "un-init" go:ngine and to release any and all GPU or RAM resources still allocated.
func Dispose() {
	Core.dispose()
	glDispose()
	UserIO.dispose()
}

//	Initializes go:ngine; this first attempts to initialize OpenGL and then open a window to your supplied specifications with a GL 3.3-or-higher profile.
func Init(options *EngineOptions) (err error) {
	var (
		glVerIndex = len(ugl.KnownVersions) - 1
		badVer     string
		glVer      float64
	)
	Core.Options = *options
tryInit:
	if Core.Options.Initialization.GlContext.CoreProfile.ForceFirst {
		for i, v := range ugl.KnownVersions {
			if v == Core.Options.Initialization.GlContext.CoreProfile.VersionHint {
				glVerIndex = i
				break
			}
		}
		glVer = ugl.KnownVersions[glVerIndex]
	}
	if err = UserIO.init(glVer); err == nil {
		if err, badVer = glInit(); err == nil && len(badVer) == 0 {
			Stats.reset()
			Loop.init()
			Core.init()
			ugl.LogLastError("INIT")
		} else if len(badVer) > 0 && !Core.Options.Initialization.GlContext.CoreProfile.ForceFirst {
			Core.Options.Initialization.GlContext.CoreProfile.ForceFirst = true
			UserIO.isGlfwInit, UserIO.Window.isCreated = false, false
			goto tryInit
		}
	} else if Core.Options.Initialization.GlContext.CoreProfile.ForceFirst && (glVerIndex > 0) {
		glVerIndex--
		UserIO.isGlfwInit, UserIO.Window.isCreated = false, false
		goto tryInit
	} else {
		badVer = glc.lastBadVer
	}
	if len(badVer) > 0 {
		err = fmtErr(glVersionErrorMessage(glMinVerStr, badVer))
	}
	return
}

func fmtErr(format string, fmtArgs ...interface{}) error {
	return fmt.Errorf(format, fmtArgs...)
}

func fmtStr(format string, fmtArgs ...interface{}) string {
	return fmt.Sprintf(format, fmtArgs...)
}
