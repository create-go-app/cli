package tools

// Ansible ...
type Ansible struct {
	BecomeSudoUser bool   `koanf:"become_sudo_user"`
	Connection     string `koanf:"connection"`
	Host           *Host  `koanf:"host"`
}

// Host ...
type Host struct {
	IP                    string `koanf:"ip"`
	User                  string `koanf:"user"`
	Group                 string `koanf:"group"`
	PythonInterpreterPath string `koanf:"python_interpreter_path"`
	ProjectFolderPath     string `koanf:"project_folder_path"`
}
