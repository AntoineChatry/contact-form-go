package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func forms() {
    tmpl := template.Must(template.ParseFiles("forms.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // Do smth not decided yet with the details ...
		_ = details
        success := true

        if success {
            successMessage := "Your form has been submitted successfully!"
			tmpl.Execute(w, struct {
                Success        bool
                SuccessMessage string
            }{true, successMessage})
        }

        // Redirect to the homepage
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })
}

func main() {
	forms()
	http.ListenAndServe(":80", nil)


}
