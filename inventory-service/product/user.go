package product

import (
	_"encoding/json"
	"fmt"
	"net/http"
)


type UserHandler struct {
	Message string
}

func (uc *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(uc.Message))
	
	
}


func SetUP(apiBasePath string){
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "users"),&UserHandler{Message: "hiii"})

	usr := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("gfgfg"))
	}

	  http.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, "users/"),usr)
	 
}