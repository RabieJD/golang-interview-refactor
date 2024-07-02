package utils

const SourceHtmlTestFile = "../testdata/utils/valid_template.html"

type TTRenderTemplate struct {
	Name        string
	PageData    interface{}
	Source      string
	Expected    string
	ExpectError bool
}

func GenerateTTRenderTemplate() []TTRenderTemplate {
	return []TTRenderTemplate{
		{
			Name:        "Valid template",
			PageData:    map[string]string{"Title": "Test"},
			Source:      SourceHtmlTestFile,
			Expected:    "<h1>Test</h1>",
			ExpectError: false,
		},
		{
			Name:        "Invalid template path",
			PageData:    map[string]string{"Title": "Test"},
			Source:      "/testdata/invalid_path.html",
			Expected:    "",
			ExpectError: true,
		},
		{
			Name:        "Template execution error",
			PageData:    make(chan int),
			Source:      "../testdata/utils/valid_template.html",
			Expected:    "",
			ExpectError: true,
		},
	}

}
