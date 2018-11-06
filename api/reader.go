package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rhasan33/goplate/models"
	"github.com/rhasan33/goplate/repository/reader"
	"golang.org/x/crypto/bcrypt"

	"github.com/rhasan33/goplate/repository"
)

// ReaderAPI ..
type ReaderAPI struct {
	repo repository.UserRepo
}

// CreateUser ..
func (ra *ReaderAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	type userData struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		ReferralCode string `json:"referral_code"`
	}

	var user userData

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		respondWithError(w, http.StatusBadRequest, errHash.Error())
		return
	}

	readerUser := models.UserSettings{
		Username:     user.Username,
		Email:        user.Email,
		ReferralCode: user.ReferralCode,
		Password:     string(passHash),
	}

	payload, errCrUser := ra.repo.CreateUser(r.Context(), &readerUser)
	if errCrUser != nil {
		respondWithError(w, http.StatusConflict, errCrUser.Error())
		return
	}

	respondwithJSON(w, http.StatusCreated, payload)
}

// GetUser ..
func (ra *ReaderAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	getUser, errGet := ra.repo.GetUser(r.Context(), id)
	if errGet != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, getUser)
}

// NewReaderAPI ..
func NewReaderAPI() *ReaderAPI {
	return &ReaderAPI{
		repo: reader.NewReader(),
	}
}
