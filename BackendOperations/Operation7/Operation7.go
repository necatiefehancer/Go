package operation7

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myModules/BackendOperations/Operation7/models"
	"net/http"
	"os"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Name:        "Kullanıcılar",
		ID:          3,
		Description: "Kullanıcı Listesi",
		URL:         "/users",
	}
	users := loadUsers()
	interests := loadInterests()
	interestMappings := loadInterestMappings()

	var newUsers []models.User

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserId {
				for _, interest := range interests {
					if interest.ID == interestMapping.InterestId {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	var viewModel = models.UserViewModel{
		Page:  page,
		Users: newUsers,
	}

	template, err := template.ParseFiles("./BackendOperations/Operation7/template/page.html")
	if err != nil {
		log.Fatal("Open template errr")
	}
	template.Execute(w, viewModel)

}

func Demo() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func loadFile(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadUsers() []models.User {
	usersJson, err := os.Open("./BackendOperations/Operation7/json/users.json")
	if err != nil {
		log.Fatal("Open Json File Error ", err.Error())
	}
	jsonDecoder := json.NewDecoder(usersJson)
	var models []models.User
	err = jsonDecoder.Decode(&models)
	if err != nil {
		log.Fatal("Decode error ", err.Error())
	}
	return models
}

func loadInterests() []models.Interest {
	InterestJson, err := os.Open("./BackendOperations/Operation7/json/interests.json")
	if err != nil {
		log.Fatal("Open Json File err ", err.Error())
	}
	jsonDecoder := json.NewDecoder(InterestJson)
	var interests []models.Interest
	err = jsonDecoder.Decode(&interests)
	if err != nil {
		log.Fatal("Decode Error ", err.Error())
	}
	return interests
}

func loadInterestMappings() []models.InterestMapping {
	IntMappingJson, err := os.Open("./BackendOperations/Operation7/json/userİnterestMappings.json")
	if err != nil {
		log.Fatal("Open Json file error ", err.Error())
	}
	jsonDecoder := json.NewDecoder(IntMappingJson)
	var mappings []models.InterestMapping
	err = jsonDecoder.Decode(&mappings)
	if err != nil {
		log.Fatal("Decode error ", err.Error())
	}
	return mappings
}
