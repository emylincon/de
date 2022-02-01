package templates

// PythonTemplate returns layout
func PythonTemplate() string {
	layout := `'''
Copyright © 2022 {{ .Name }} <{{ .Email }}>
'''


def main():
	print("Hello World")


if __name__ == "__main__":
	main()
	`
	return layout
}
