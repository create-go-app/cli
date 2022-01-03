package cgapp

import (
	"io/fs"
	"os"
	"testing"

	"github.com/create-go-app/cli/v3/pkg/registry"
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
		err := os.RemoveAll("Makefile")
		if err != nil {
			return
		}
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

	_ = os.Mkdir("cgapp-project", 0o750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFolder(tt.args.folderName); (err != nil) != tt.wantErr {
				t.Errorf("MakeFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		err := os.RemoveAll(tt.args.folderName)
		if err != nil {
			return
		}
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

	_ = os.MkdirAll("../../tmp/folder-1", 0o750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveFolders(tt.args.rootFolder, tt.args.foldersToRemove)
		})
	}
}

func TestCopyFromEmbeddedFS(t *testing.T) {
	type args struct {
		efs *EmbeddedFileSystem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully copy from embedded fs",
			args{
				efs: &EmbeddedFileSystem{
					Name:       registry.EmbedTemplates,
					RootFolder: "templates",
					SkipDir:    false,
				},
			},
			false,
		},
		{
			"successfully copy from embedded fs with skip dirs",
			args{
				efs: &EmbeddedFileSystem{
					Name:       registry.EmbedTemplates,
					RootFolder: "templates",
					SkipDir:    true,
				},
			},
			false,
		},
		{
			"fail to copy from embedded fs",
			args{
				efs: &EmbeddedFileSystem{
					Name:       registry.EmbedTemplates,
					RootFolder: "does-not-exist",
					SkipDir:    false,
				},
			},
			true,
		},
		{
			"fail (no args)",
			args{
				efs: &EmbeddedFileSystem{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFromEmbeddedFS(tt.args.efs); (err != nil) != tt.wantErr {
				t.Errorf("CopyFromEmbeddedFS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		_ = os.Remove("hosts.ini.tmpl")
		_ = os.Remove("playbook.yml.tmpl")
		_ = os.RemoveAll(tt.args.efs.RootFolder)
	}
}

func TestGenerateFileFromTemplate(t *testing.T) {
	type args struct {
		fileName  string
		variables map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully generate file",
			args{
				fileName:  "../../tmp/test.txt",
				variables: map[string]interface{}{},
			},
			false,
		},
		{
			"failed to generate file",
			args{},
			true,
		},
	}

	_ = os.Mkdir("../../tmp", 0o750)
	_, _ = os.Create("../../tmp/test.txt")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateFileFromTemplate(tt.args.fileName, tt.args.variables); (err != nil) != tt.wantErr {
				t.Errorf("GenerateFileFromTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		_ = os.RemoveAll("../../tmp")
	}
}
