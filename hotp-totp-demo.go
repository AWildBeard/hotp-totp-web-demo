package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/hotp"
	"github.com/pquerna/otp/totp"
	"html/template"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
)

type Registry struct {
	users []Keys
	mux *HttpHandler
}

func NewRegistry(mux *HttpHandler) Registry {
	return Registry{
		make([]Keys, 0),
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
	owner string
	hotpKey otp.Key
	hotpCounter uint64
	totpKey otp.Key
	initialized bool
}

func (keys *Keys) increment() {
	keys.hotpCounter = keys.hotpCounter + 1
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
	mux.functions["/testHotp"] = registry.TestHotp
	mux.functions["/testTotp"] = registry.TestTotp
	mux.functions["/logout"] = registry.Logout

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

			registry.users = append(registry.users, Keys{
				owner: email,
			})

			r.Method = "GET"

			http.Redirect(w, r, "/otp-setup", http.StatusFound)
			return
		}
	}

	w.Write([]byte("BAD LOGIN ATTEMPT!"))
}

func (registry *Registry) TestHotp(w http.ResponseWriter, r *http.Request) {
	type Details struct {
		Email string `json:"email"`
		KeyToTest string `json:"key"`
	}
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-Type") == "application/json" {
			var deets Details

			if content, err := ioutil.ReadAll(r.Body); err == nil {
				if json.Valid(content) && json.Unmarshal(content, &deets) == nil {
					var keysToValidate *Keys
					found := false
					for index, key := range registry.users {
						if key.owner == deets.Email {
							keysToValidate = &registry.users[index]
							found = true
							break
						}
					}

					if found && hotp.Validate(deets.KeyToTest, keysToValidate.hotpCounter,
						keysToValidate.hotpKey.Secret()) {
						keysToValidate.hotpCounter++
						w.Write([]byte("{\"status\": \"success\"}"))
						return
					}
					keysToValidate.hotpCounter++
				}
			}
		}
	default:
		fmt.Println("Method:", r.Method)
	}
	w.Write([]byte("{\"status\": \"failure\"}"))
}

func (registry *Registry) TestTotp(w http.ResponseWriter, r *http.Request) {
	type Details struct {
		Email string `json:"email"`
		KeyToTest string `json:"key"`
	}
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-Type") == "application/json" {
			var deets Details

			if content, err := ioutil.ReadAll(r.Body); err == nil {
				if json.Valid(content) && json.Unmarshal(content, &deets) == nil {
					var keysToValidate *Keys
					found := false
					for index, key := range registry.users {
						if key.owner == deets.Email {
							keysToValidate = &registry.users[index]
							found = true
							break
						}
					}

					if found && totp.Validate(deets.KeyToTest, keysToValidate.totpKey.Secret()) {
						w.Write([]byte("{\"status\": \"success\"}"))
						return
					}
				}
			}
		}
	default:
		fmt.Println("Method:", r.Method)
	}
	w.Write([]byte("{\"status\": \"failure\"}"))

}

func (registry *Registry) OtpSetup(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("registered"); err == nil {
		found := false
		var keyToAdd *Keys
		for index, key := range registry.users {
			if cookie.Value == key.owner {
				keyToAdd = &registry.users[index]
				found = true
				break
			}
		}

		if found && ! keyToAdd.initialized {
			totpKey := otp.Key{}
			hotpKey := otp.Key{}

			totpUri := ""
			hotpUri := ""

			if newTotp, err := totp.Generate(totp.GenerateOpts{
				Issuer: "AWildBeard",
				AccountName: cookie.Value + "-TOTP",
			}) ; err == nil {
				totpKey = *newTotp
			} else {
				fmt.Println("Failed to create TOTP key!")
			}

			if newHotp, err := hotp.Generate(hotp.GenerateOpts{
				Issuer: "AWildBeard",
				AccountName: cookie.Value + "-HOTP",
			}); err == nil {
				hotpKey = *newHotp
			} else {
				fmt.Println("Failed to create HOTP key!")
			}

			keyToAdd.hotpKey = hotpKey
			keyToAdd.hotpCounter = 1
			keyToAdd.totpKey = totpKey
			keyToAdd.initialized = true

			if newImage, err := hotpKey.Image(200, 200); err == nil {
				hotpUri = "/" + cookie.Value + "-hotp"

				hotpImageHolder := ImageHolder{
					image: newImage,
					registry: registry,
					uri: hotpUri,
				}

				registry.mux.functions[hotpUri] = hotpImageHolder.ServeImage
			}

			if newImage, err := totpKey.Image(200, 200); err == nil {
				totpUri = "/" + cookie.Value + "-totp"

				totpImageHolder := ImageHolder{
					image: newImage,
					registry: registry,
					uri: totpUri,
				}

				registry.mux.functions[totpUri] = totpImageHolder.ServeImage
			}

			type returnType struct {
				Email string
				HotpImagePath string
				TotpImagePath string
			}

			htmlTemplate, err := template.New("Output").Parse(setupOTPHTMLPage)
			if err != nil {
				fmt.Println("Encountered error:", err)
			}

			if err := htmlTemplate.Execute(w, returnType{
				Email: cookie.Value,
				HotpImagePath: hotpUri,
				TotpImagePath: totpUri,
			}) ; err != nil {
				fmt.Println("Encountered error:", err)
			}
		} else {
			if found && keyToAdd.initialized {
				http.Redirect(w, r, "/otp", http.StatusFound)
			} else {
				http.Redirect(w, r, "/logout", http.StatusFound)
			}
		}
	} else {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
}

func (registry *Registry) OtpViewer(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/logout", http.StatusFound)

}

func (registry *Registry) Logout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("registered"); err == nil {
		http.SetCookie(w, &http.Cookie{
			Name: "registered",
			Value: cookie.Value,
			MaxAge: -1,
		})

		for index, key := range registry.users {
			if cookie.Value == key.owner {
				registry.users = append(registry.users[:index], registry.users[index + 1:]...)
				break
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (registry *Registry) RootHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("registered"); err == nil {
		for _, key := range registry.users {
			if cookie.Value == key.owner {
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