package zlog

import (
	"github.com/sohaha/zlsgo/zfile"
)

var LogMaxDurationDate = 15

func openFile(filepa string, archive bool) (file *zfile.MemoryFile, fileName, fileDir string, err error) {
	_ = "STUB: not implemented"
	return nil, "", "", nil
}

// Delete the log file that is too old

// SetFile Setting log file output
func (log *Logger) SetFile(filepath string, archive ...bool) { _ = "STUB: not implemented"; return }

func (log *Logger) SetLevelFile(level int, filepath string, archive ...bool) {
	_ = "STUB: not implemented"
	return
}

func (log *Logger) SetLevelSaveFile(level int, filepath string, archive ...bool) {
	_ = "STUB: not implemented"
	return
}

func (log *Logger) setLogfile(filepath string, archive bool) { _ = "STUB: not implemented"; return }

func (log *Logger) setLevelFile(level int, filepath string, archive bool, andStdout bool) {
	_ = "STUB: not implemented"
	return
}

func (log *Logger) Discard() { _ = "STUB: not implemented"; return }

func (log *Logger) SetSaveFile(filepath string, archive ...bool) { _ = "STUB: not implemented"; return }

func (log *Logger) CloseLevelFiles() { _ = "STUB: not implemented"; return }

func (log *Logger) closeLevelFilesLocked() { _ = "STUB: not implemented"; return }

func (log *Logger) CloseFile() { _ = "STUB: not implemented"; return }
