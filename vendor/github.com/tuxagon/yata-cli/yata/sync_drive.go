package yata

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// DriveAPI wraps the Google Drive API within
// an implementation of the SyncAPI interface
type DriveAPI struct {
	cfgMgr *ConfigManager
	srv    *drive.Service
}

// Push uploads Yata files to Google Drive under the
// application data folder
func (m DriveAPI) Push() {
	GetLogger().Verbose("Initiated push to Google Drive")

	fileList := m.appDataFiles()
	for _, sf := range SynchronizableFiles {
		var found bool

		GetLogger().Info("Pushing: " + sf.Name)

		fileMetadata := drive.File{
			Name: sf.Name,
		}

		file, err := os.Open(sf.Path)
		HandleError(err, true)
		defer func() { file.Close() }()

		for _, f := range fileList.Files {
			if f.Name == sf.Name {
				found = true
				_, err = m.srv.Files.Update(f.Id, &fileMetadata).Media(file).Do()
				HandleError(err, true)
				break
			}
		}

		if !found {
			fileMetadata.Parents = []string{"appDataFolder"}
			_, err = m.srv.Files.Create(&fileMetadata).Media(file).Do()
			HandleError(err, true)
		}
	}
}

// Fetch downloads Yata application files from Google Drive to the
// specified fetch directory
func (m DriveAPI) Fetch() {
	GetLogger().Verbose("Initiated fetch from Google Drive")

	GetDirectory().EmptyFetchDir()

	fileList := m.appDataFiles()
	for _, f := range fileList.Files {
		GetLogger().Info("Fetching: " + f.Name)

		resp, err := m.srv.Files.Get(f.Id).Download()
		HandleError(err, true)

		contents := make([]byte, 0)
		buf := make([]byte, 1024)
		for {
			n, err := resp.Body.Read(buf)
			HandleErrorWithFunc(err, false, func(err error) bool { return err != io.EOF })
			if n == 0 || err == io.EOF {
				break
			}

			contents = appendBytes(contents, buf)
		}
		GetLogger().Verbose("\n" + string(contents) + "\n")

		path := filepath.Join(GetDirectory().FetchDir(), f.Name)
		fo, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
		HandleError(err, true)
		_, err = fo.Write(contents)
		HandleError(err, true)
		defer func() { fo.Close() }()
	}
}

// appDataFiles will retrieve a list of all the application data files found
// in Google Drive under the approved Yata app
func (m DriveAPI) appDataFiles() *drive.FileList {
	fileList, err := m.srv.Files.
		List().
		Spaces("appDataFolder").
		Fields("nextPageToken, files(id, name)").
		PageSize(10).
		Do()
	HandleError(err, true)
	return fileList
}

// getConfig gets the Google Drive API oauth config
func (m DriveAPI) getConfig() (*oauth2.Config, error) {
	secPath := filepath.Join(GetDirectory().Root, m.cfgMgr.Config.GoogleDrive.SecretFile)
	dat, err := ioutil.ReadFile(secPath)
	HandleError(err, true)
	return google.ConfigFromJSON(dat, drive.DriveAppdataScope)
}

// setService initializes the Google Drive service
func (m *DriveAPI) setService() {
	cfg, _ := m.getConfig()
	client, _ := m.getClient(context.Background(), cfg)
	m.srv, _ = drive.New(client)
}

// getClient gets the HTTP client needed by the Google Drive
// API service
func (m DriveAPI) getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	tok, err := m.tokenFromConfig()
	if err != nil {
		tok = m.tokenFromWeb(config)
		m.saveToken(tok)
	}
	return config.Client(ctx, tok), nil
}

// tokenFromConfig attempts to read the Google Drive oauth token
// from the Yata config
func (m DriveAPI) tokenFromConfig() (*oauth2.Token, error) {
	tok := m.cfgMgr.Config.GoogleDrive.OAuthToken
	if tok != nil {
		return tok, nil
	}
	return nil, fmt.Errorf("Invalid token")
}

// tokenFromWeb will open a web page where Google requests you give Yata
// permission to save files to Google Drive
func (m DriveAPI) tokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	open.Run(authURL)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	_, err := fmt.Scan(&code)
	HandleError(err, true)

	tok, err := config.Exchange(oauth2.NoContext, code)
	HandleError(err, true)
	return tok
}

// saveToken writes the Google Drive oauth token to the Yata config
func (m DriveAPI) saveToken(tok *oauth2.Token) {
	m.cfgMgr.SetKey("googledrive.oauthtoken", tok)
}
