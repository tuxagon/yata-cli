package yata

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
}

// NewServerManager TODO docs
func NewServerManager(serverType int) ServerManager {
	switch serverType {
	case GoogleDrive:
		return &GoogleDriveManager{
			cfgMgr: NewConfigManager(),
		}
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
	cfg, err := m.getConfig()
	if err != nil {
		return err
	}
	ctx := context.Background()
	client, err := m.getClient(ctx, cfg)
	if err != nil {
		return err
	}

	srv, err := drive.New(client)
	if err != nil {
		return err
	}

	fileMetadata := &drive.File{
		Name:    "tasks.json",
		Parents: []string{"appDataFolder"},
	}

	_, err = srv.Files.Create(fileMetadata).Do()
	if err != nil {
		return err
	}

	return nil
}

// Fetch TODO docs
func (m GoogleDriveManager) Fetch() error {
	return nil
}

func (m GoogleDriveManager) getConfig() (*oauth2.Config, error) {
	dirSrv := NewDirectoryService()
	secPath := filepath.Join(dirSrv.RootPath, m.cfgMgr.Config.GoogleDrive.SecretFile)

	dat, err := ioutil.ReadFile(secPath)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(dat, drive.DriveAppdataScope)
}

func (m GoogleDriveManager) getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	tok, err := m.tokenFromConfig()
	if err != nil {
		tok = m.tokenFromWeb(config)
		m.saveToken(tok)
	}
	return config.Client(ctx, tok), nil
}

func (m GoogleDriveManager) tokenFromConfig() (*oauth2.Token, error) {
	tok := m.cfgMgr.Config.GoogleDrive.OAuthToken
	if tok.Valid() {
		return tok, nil
	}
	return nil, fmt.Errorf("Invalid token")
}

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

func (m GoogleDriveManager) saveToken(tok *oauth2.Token) {
	m.cfgMgr.SetKey("googledrive.oauthtoken", tok)
}
