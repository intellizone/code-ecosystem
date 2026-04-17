package teacherportal

import "html/template"

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"pkg/teacherportal/templates/students.go.html",
		"pkg/teacherportal/templates/student.go.html",
	)
	if err != nil {
		return err
	}
	return nil
}
