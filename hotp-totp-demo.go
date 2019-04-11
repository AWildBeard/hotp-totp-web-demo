package main

import (
	"bytes"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/hotp"
	"github.com/pquerna/otp/totp"
	"image"
	"image/png"
	"io"
	"net/http"
)

type Registry struct {
	users map[string]Keys
	mux *HttpHandler
}

func NewRegistry(mux *HttpHandler) Registry {
	return Registry{
		map[string]Keys {},
		mux,
	}
}

type HttpHandler struct {
	functions map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{
		map[string] func(w http.ResponseWriter, r *http.Request) {},
	}
}

type Keys struct {
	hotpKey otp.Key
	totpKey otp.Key
}

type ImageHolder struct {
	image image.Image
	registry *Registry
	uri string
}

func main() {
	mux := NewHttpHandler()
	registry := NewRegistry(&mux)
	server := http.Server{
		Addr: ":80",
		Handler: &mux,
	}

	registry.mux = &mux

	mux.functions["/"] = registry.RootHandler
	mux.functions["/register"] = registry.Register
	mux.functions["/otp-setup"] = registry.OtpSetup
	mux.functions["/otp"] = registry.OtpViewer

	panic(server.ListenAndServe())
}

func (registry *Registry) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.Copy(w, bytes.NewBufferString(registerHTMLPage))
		return
	case "POST":
		if r.ParseForm() == nil {
			email := r.FormValue("email")

			if email == "" { w.Write([]byte("BAD LOGIN ATTEMPT!")) }

			http.SetCookie(w, &http.Cookie{
				Name: "registered",
				Value: email,
			})

			registry.users[email] = Keys{}

			r.Method = "GET"

			http.Redirect(w, r, "/otp-setup", http.StatusFound)
			return
		}
	}

	w.Write([]byte("BAD LOGIN ATTEMPT!"))
}

func (registry *Registry) OtpSetup(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("registered"); err == nil {
		found := false
		for key := range registry.users {
			if cookie.Value == key {
				found = true
				break
			}
		}

		if found {
			totpKey := otp.Key{}
			hotpKey := otp.Key{}

			if newTotp, err := totp.Generate(totp.GenerateOpts{
				Issuer: "AWildBeard",
				AccountName: cookie.Value,
			}) ; err == nil {
				totpKey = *newTotp
			} else {
				fmt.Println("Failed to create TOTP key!")
			}

			if newHotp, err := hotp.Generate(hotp.GenerateOpts{
				Issuer: "AWildBeard",
				AccountName: cookie.Value,
			}); err == nil {
				hotpKey = *newHotp
			} else {
				fmt.Println("Failed to create HOTP key!")
			}

			registry.users[cookie.Value] = Keys{
				hotpKey: hotpKey,
				totpKey: totpKey,
			}

			if newImage, err := hotpKey.Image(200, 200); err == nil {
				uri := cookie.Value + "-hotp"

				hotpImageHolder := ImageHolder{
					image: newImage,
					registry: registry,
					uri: uri,
				}

				registry.mux.functions[uri] = hotpImageHolder.ServeImage
			}

			if newImage, err := totpKey.Image(200, 200); err == nil {
				uri := cookie.Value + "-totp"

				totpImageHolder := ImageHolder{
					image: newImage,
					registry: registry,
					uri: uri,
				}

				registry.mux.functions[uri] = totpImageHolder.ServeImage
			}
		}
	}
}

func (registry *Registry) OtpViewer(w http.ResponseWriter, r *http.Request) {

}

func (registry *Registry) RootHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("registered"); err == nil {
		for key := range registry.users {
			if cookie.Value == key {
				http.Redirect(w, r, "/otp", http.StatusFound)
				break
			}
		}

	} else {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}

func (imgHolder *ImageHolder) ServeImage(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}

	png.Encode(&buf, imgHolder.image)

	delete(imgHolder.registry.mux.functions, imgHolder.uri)

	io.Copy(w, &buf)
}

func (handler *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler.functions[r.RequestURI] != nil {
		handler.functions[r.RequestURI](w, r)
	} else {
		w.Write([]byte("HEH! THAT TICKLES!"))
	}
}