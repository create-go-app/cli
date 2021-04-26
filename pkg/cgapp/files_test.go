package cgapp

import (
	"io/fs"
	"os"
	"testing"

	"github.com/create-go-app/cli/pkg/registry"
)

func TestMakeFile(t *testing.T) {

	fileData, err := fs.ReadFile(registry.EmbedMiscFiles, "misc/Makefile")
	if err != nil {
		t.Error()
	}

	type args struct {
		rootFolder string
		file       string
		data       []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully created files",
			args{
				rootFolder: "Makefile",
				file:       "Makefile",
				data:       fileData,
			},
			false,
		},
		{
			"failed created files",
			args{
				rootFolder: "",
				file:       "",
				data:       fileData,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFile(tt.args.rootFolder, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("MakeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("Makefile")
	}
}

func TestMakeFolder(t *testing.T) {
	type args struct {
		folderName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully created folder",
			args{
				folderName: "../../tmp",
			},
			false,
		},
		{
			"failed, folder is exists",
			args{
				folderName: "",
			},
			true,
		},
		{
			"failed, folder is exists",
			args{
				folderName: "cgapp-project",
			},
			true,
		},
	}

	_ = os.Mkdir("cgapp-project", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFolder(tt.args.folderName); (err != nil) != tt.wantErr {
				t.Errorf("MakeFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll(tt.args.folderName)
	}
}

func TestRemoveFolders(t *testing.T) {
	type args struct {
		rootFolder      string
		foldersToRemove []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"successfully removed",
			args{
				rootFolder:      "../../tmp",
				foldersToRemove: []string{"folder-1"},
			},
		},
	}

	_ = os.MkdirAll("../../tmp/folder-1", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveFolders(tt.args.rootFolder, tt.args.foldersToRemove)
		})
	}
}
