package yata

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const (
	GoogleDrive = iota
	Dropbox
)

const noServerMsgFmt = "I would love to %s your tasks %s a server, but I don't see any configured yet! If you want to configure a server, type `yata %s --help`"

// PushableFiles TODO docs
var PushableFiles = []struct {
	Name, Path string
}{
	{Name: "tasks.json", Path: NewDirectoryService().GetFullPath()},
	{Name: ".yataid", Path: NewDirectoryService().GetFullIDPath()},
}

// ServerManager TODO docs
type ServerManager interface {
	Push() error
	Fetch() error
}

// NoServerManager TODO docs
type NoServerManager struct{}

// GoogleDriveManager TODO docs
type GoogleDriveManager struct {
	cfgMgr *ConfigManager
	srv    *drive.Service
}

// NewServerManager TODO docs
func NewServerManager(serverType int) ServerManager {
	switch serverType {
	case GoogleDrive:
		m := &GoogleDriveManager{
			cfgMgr: NewConfigManager(),
		}
		m.setService()
		return m
	default:
		return &NoServerManager{}
	}
}

// Push TODO docs
func (m NoServerManager) Push() error {
	return fmt.Errorf(noServerMsgFmt, "push", "to", "push")
}

// Fetch TODO docs
func (m NoServerManager) Fetch() error {
	return fmt.Errorf(noServerMsgFmt, "fetch", "from", "fetch")
}

// Push TODO docs
func (m GoogleDriveManager) Push() error {
	fileList, err := m.appDataFiles()
	if err != nil {
		return err
	}

	for _, pf := range PushableFiles {
		var found bool

		fileMetadata := drive.File{
			Name: pf.Name,
		}

		file, err := os.Open(pf.Path)
		if err != nil {
			return err
		}
		defer func() {
			file.Close()
		}()

		for _, f := range fileList.Files {
			if f.Name == pf.Name {
				found = true

				_, err = m.srv.Files.Update(f.Id, &fileMetadata).Media(file).Do()
				if err != nil {
					return err
				}

				break
			}
		}

		if !found {
			fileMetadata.Parents = []string{"appDataFolder"}
			_, err = m.srv.Files.Create(&fileMetadata).Media(file).Do()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Fetch TODO docs
func (m GoogleDriveManager) Fetch() error {
	NewDirectoryService().ClearFetchFiles()

	fileList, err := m.appDataFiles()
	if err != nil {
		return err
	}

	for _, f := range fileList.Files {
		resp, err := m.srv.Files.Get(f.Id).Download()
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		path := filepath.Join(NewDirectoryService().GetFetchPath(), f.Name)
		for {
			n, err := resp.Body.Read(buf)
			if err != nil && err != io.EOF {
				return err
			}

			if n == 0 {
				break
			}

			fo, err := os.Create(path)
			if err != nil {
				return err
			}
			defer func() {
				fo.Close()
			}()

			if _, err := fo.Write(buf[:n]); err != nil {
				return err
			}
		}
	}

	return nil
}

// appDataFiles TODO docs
func (m GoogleDriveManager) appDataFiles() (*drive.FileList, error) {
	return m.srv.Files.List().Spaces("appDataFolder").Fields("nextPageToken, files(id, name)").PageSize(10).Do()
}

// getConfig TODO docs
func (m GoogleDriveManager) getConfig() (*oauth2.Config, error) {
	secPath := filepath.Join(NewDirectoryService().RootPath, m.cfgMgr.Config.GoogleDrive.SecretFile)

	dat, err := ioutil.ReadFile(secPath)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(dat, drive.DriveAppdataScope)
}

// setService TODO docs
func (m *GoogleDriveManager) setService() {
	cfg, _ := m.getConfig()
	client, _ := m.getClient(context.Background(), cfg)
	m.srv, _ = drive.New(client)
}

// getClient TODO docs
func (m GoogleDriveManager) getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	tok, err := m.tokenFromConfig()
	if err != nil {
		tok = m.tokenFromWeb(config)
		m.saveToken(tok)
	}
	return config.Client(ctx, tok), nil
}

// tokenFromConfig TODO docs
func (m GoogleDriveManager) tokenFromConfig() (*oauth2.Token, error) {
	tok := m.cfgMgr.Config.GoogleDrive.OAuthToken
	if tok != nil {
		return tok, nil
	}
	return nil, fmt.Errorf("Invalid token")
}

// tokenFromWeb TODO docs
func (m GoogleDriveManager) tokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	open.Run(authURL)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// saveToken TODO docs
func (m GoogleDriveManager) saveToken(tok *oauth2.Token) {
	m.cfgMgr.SetKey("googledrive.oauthtoken", tok)
}
