package plague

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/rs/cors"
)

const (

	//AllowedHeaders for cors headers access
	AllowedHeaders = "Accept,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Access-Control-Expose-Headers,X-Requested-With,Content-Type,Authorization,Auth"
	//DefaultServerPort : the default service endpoint entrance
	DefaultServerPort string = ":9528"
	//Root : url root path
	Root = "/plague"

	/*****************************socks and rests*****************************/
	//InsRestSetURL : install settings for restful access
	URLCheckingFrontdoor  = "/c/d"
	URLLogin              = "/l"
	URLInputsEvery        = "/inputs/e"
	URLFetchEveryOneBasic = "/fb/{page}"
	URLFetchByArea        = "/fa"
	URLDel                = "/d/{index}"
	URLGet                = "/cnt"
	/*****************************socks and rests*****************************/

	/**************************AUTHSECRETWORD****************************/
	//AuthenticationName : the authentication url path
	AuthenticationName = "Auth"
	AuthToken          = "FF315A2408701FEA"
	//AuthNotApproved : the authentication failure
	AuthNotApproved = "not-approved-auth"
	/**************************AUTHSECRETWORD****************************/

)

//WebForDistributions config
func WebForDistributionsPlagues() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("WebForDistributionsPlagues-Error: ", err)
		}
	}()
	fmt.Println("DefaultServer == " + DefaultServerPort)
	ro := mux.NewRouter().StrictSlash(true)
	ro.HandleFunc(strings.Join([]string{Root, URLCheckingFrontdoor}, ""), GetCompanies).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLLogin}, ""), Login).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLInputsEvery}, ""), EveryoneInfoInput).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLFetchEveryOneBasic}, ""), FetchEveryoneBasic).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLFetchByArea}, ""), FetchByCurrentArea).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLDel}, ""), DeleteOnePlagueMan).Methods("POST")
	ro.HandleFunc(strings.Join([]string{Root, URLGet}, ""), GetCountMain).Methods("POST")
	http.ListenAndServe(DefaultServerPort,
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
			handlers.AllowedHeaders(strings.Split(AllowedHeaders, ",")),
			handlers.AllowedOrigins([]string{"*"}))(ro))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Auth")
}

func DealWithAuthToken(w http.ResponseWriter, r *http.Request) string {
	authToken := r.Header.Get(AuthenticationName)
	fmt.Println("DealWithAuthToken::" + authToken)
	if len(authToken) == 0 || strings.Compare(AuthToken, authToken) != 0 {
		SendBackCommonHead("{\"TraceBackServer\":\"not authorized\"}", http.StatusForbidden, w)
		return AuthNotApproved
	}
	return authToken
}

func GetCountMain(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("DeleteOnePlagueMan-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	setupResponse(&w, r)
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	total := GetCountTotal()
	ReturnABunchOfStuffDirect(w, total, http.StatusOK)
}

func DeleteOnePlagueMan(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("DeleteOnePlagueMan-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	setupResponse(&w, r)
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	vars := mux.Vars(r)
	index := vars["index"]
	DeleteOneMan(index)
	ReturnABunchOfStuffDirect(w, "s", http.StatusOK)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GetCompanies-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	setupResponse(&w, r)
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	comps := GetComps()
	fmt.Println("comp::" + comps)
	ReturnABunchOfStuffDirect(w, comps, http.StatusOK)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Login-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	setupResponse(&w, r)
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	bodyData, errBody := ioutil.ReadAll(r.Body)
	if checkCommError(w, errBody) != 0 {
		return
	}
	bodyDataContext := string(bodyData[:])
	fmt.Println(bodyDataContext)
	if len(bodyDataContext) > 0 {
		lgn := AdminInfo{}
		errJson := json.Unmarshal(bodyData, &lgn)
		if errJson != nil {
			fmt.Println("admin info is not parsed")
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		} else {
			adm := AdminLogin(lgn)
			ReturnABunchOfStuffDirect(w, adm, http.StatusOK)
		}
	} else {
		ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
		return
	}
}

//Inputs info of everyone
func EveryoneInfoInput(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("EveryoneInfoInput-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	setupResponse(&w, r)
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	bodyData, errBody := ioutil.ReadAll(r.Body)
	if checkCommError(w, errBody) != 0 {
		return
	}
	bodyDataContext := string(bodyData[:])
	if len(bodyDataContext) > 0 {
		fmt.Printf("body (EveryoneInfo)==%s\n", bodyDataContext)
		everyone := EveryoneInfo{}
		err := json.Unmarshal(bodyData, &everyone)
		if err != nil {
			fmt.Printf("EveryoneInfoInputErr==%T\n", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail(err.Error()), http.StatusNoContent)
		} else {
			Recording(everyone)
			ReturnABunchOfStuffDirect(w, "s", http.StatusOK)
			return
		}
	} else {
		ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
	}
}

func FetchEveryoneBasic(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("FetchEveryoneBasic-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	bodyData, errBody := ioutil.ReadAll(r.Body)
	if checkCommError(w, errBody) != 0 {
		return
	}
	vars := mux.Vars(r)
	page := vars["page"]
	bodyDataContext := string(bodyData[:])
	if len(bodyDataContext) > 0 {
		fmt.Printf("body (EveryoneInfoBasic)==%s\n", bodyDataContext)
		filter := FetchFilter{}
		err := json.Unmarshal(bodyData, &filter)
		if err != nil {
			fmt.Printf("FetchEveryoneBasic==%T\n", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail(err.Error()), http.StatusNoContent)
			return
		} else {
			p, _ := strconv.Atoi(page)
			data := GetAllPagedFetches(filter, p)
			dataBytes, errJson := json.Marshal(&data)
			if errJson != nil {
				fmt.Printf("FetchEveryoneBasic==%T\n", errJson)
				ReturnABunchOfStuff(w, FeedBackReturnFail(errJson.Error()), http.StatusNoContent)
				return
			} else {
				dataStr := string(dataBytes[:])
				fmt.Println("#FetchEveryoneBasic::\n" + dataStr)
				ReturnABunchOfStuffDirect(w, dataStr, http.StatusOK)
				return
			}
		}
	} else {
		ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
		return
	}
}

func FetchByCurrentArea(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("FetchByCurrentArea-Error: ", err)
			ReturnABunchOfStuff(w, FeedBackReturnFail("rest-api set internal error"), http.StatusNoContent)
			return
		}
	}()
	auth := DealWithAuthToken(w, r)
	if strings.Compare(auth, AuthNotApproved) == 0 {
		return
	}
	data := FindPassedZones()
	dataBytes, errJson := json.Marshal(&data)
	if errJson != nil {
		fmt.Printf("FetchByCurrentArea==%T\n", errJson)
		ReturnABunchOfStuff(w, FeedBackReturnFail(errJson.Error()), http.StatusNoContent)
		return
	} else {
		dataStr := string(dataBytes[:])
		fmt.Println("#FetchByCurrentArea::\n" + dataStr)
		ReturnABunchOfStuffDirect(w, dataStr, http.StatusOK)
		return
	}

}

//Redirection method for http requests redirections with given destination
func Redirection(w http.ResponseWriter, r *http.Request, urlDest string) {
	setupResponse(&w, r)
	fmt.Print("REDIRECT_TO====%s\n" + urlDest)
	http.Redirect(w, r, urlDest, http.StatusMovedPermanently)
}

//DoReqsRHB request helps getting feedback from a url. it's like a gate keeper for me.
func DoReqsRHB(URL string, method string, reqBody string, reqHeaders map[string]string) ([]byte, *http.Header) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("DoReqsRHB-Error: ", err)
			return
		}
	}()
	cli := http.Client{}
	fmt.Printf("req==%s--%s--%s\n", URL, method, reqBody)
	req, err := http.NewRequest(method, URL, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("DoReqsRHB req-error::", err)
	}
	if len(reqHeaders) > 0 {
		for k, v := range reqHeaders {
			fmt.Println("HEAD==" + k + "::" + v)
			req.Header.Add(k, v)
		}
	}
	response, _ := cli.Do(req)
	resp, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("resp==%s\n", string(resp))
	return resp[:], &response.Header
}

func addHeadersToAccess(addAccess map[string]string, accessHeaders map[string]string) {
	if len(accessHeaders) > 0 {
		for k, v := range accessHeaders {
			addAccess[k] = v
		}
	}
}

func checkCommError(w http.ResponseWriter, err error) int {
	if err != nil {
		fmt.Printf("body-fetch-err==%s\n", err)
		ReturnABunchOfStuff(w, FeedBackReturnFail(err.Error()), http.StatusNoContent)
		return 1
	}
	return 0
}

//DeferFunc : for everything try-catch
func deferFunc(descriptor string) {
	err := recover()
	if err != nil {
		fmt.Println("FUNC:"+descriptor+"-Err: ", err)
		return
	}
}
