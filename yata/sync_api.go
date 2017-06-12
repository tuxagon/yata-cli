package yata

import (
	"fmt"
)

const (
	GoogleDrive = iota
	Dropbox
)

// SynchronizableFiles represents each file that will be synchronized
var SynchronizableFiles = []struct {
	Name, Path string
}{
	{Name: "tasks.json", Path: GetDirectory().TasksFilePath()},
	{Name: "id", Path: GetDirectory().IDPath()},
}

// SyncAPI provides the contract necessary for any service that
// synchronizes yata data
type SyncAPI interface {
	Push()
	Fetch()
}

// NoneAPI represents the lack of a configured synchronization
// service, so that the nil case does not need handled
// separately
type NoneAPI struct{}

// Push will display a helpful message about configuring a
// synchronization service to the user
func (m NoneAPI) Push() {
	HandleError(fmt.Errorf("No server has been configured to push yet! "+
		"If you want to configure a server, please consult the README"), true)
}

// Fetch will display a helpful message about configuring a
// synchronization service to the user
func (m NoneAPI) Fetch() {
	HandleError(fmt.Errorf("No server has been configured to fetch yet!"+
		"If you want to configure a server, please consult the README"), true)
}

// NewSyncAPI creates a new instance of a synchronization service
func NewSyncAPI(serverType int) SyncAPI {
	switch serverType {
	case GoogleDrive:
		GetLogger().Verbose("Using Google Drive")
		m := &DriveAPI{
			cfgMgr: NewConfigManager(),
		}
		m.setService()
		return m
	default:
		return &NoneAPI{}
	}
}

// appendBytes takes 2 byte slices and appends the second
// to the first, similar to how append works with a
// single item in a slice
func appendBytes(buf, newBytes []byte) []byte {
	for _, b := range newBytes {
		buf = append(buf, b)
	}
	return buf
}
