package utils

import (
	"os"
	"testing"
)

func TestMakeFiles(t *testing.T) {
	type args struct {
		rootFolder  string
		filesToMake map[string][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully created files",
			args{
				rootFolder: "../../tmp",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			false,
		},
		{
			"failed created files",
			args{
				rootFolder: "./does/not-exists",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			true,
		},
	}

	_ = os.Mkdir("../../tmp", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFiles(tt.args.rootFolder, tt.args.filesToMake); (err != nil) != tt.wantErr {
				t.Errorf("MakeFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("../../tmp")
	}
}
func TestRemoveFolders(t *testing.T) {
	type args struct {
		rootFolder      string
		foldersToRemove []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully removed",
			args{
				rootFolder:      "../../tmp",
				foldersToRemove: []string{"folder-1"},
			},
			false,
		},
	}

	_ = os.MkdirAll("../../tmp/folder-1", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFolders(tt.args.rootFolder, tt.args.foldersToRemove); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFolders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
