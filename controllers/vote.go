package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/BruzGus/edcomment/commons"
	"github.com/BruzGus/edcomment/configuration"
	"github.com/BruzGus/edcomment/models"
)

//VoteRegister ..., controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar:%s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID
	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? and user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)

	// Si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		err := updateCommentVotes(vote.CommentID, vote.Value)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
		}
		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto Actualizado"
		m.Code = http.StatusOK
		commons.DisplayMessage(w, m)
		return
	}
	m.Message = "Este voto ya esta registrado"
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}

//funcion privada
func updateCommentVotes(commentID uint, vote bool) (err error) {

	comment := models.Comment{}
	db := configuration.GetConnection()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected
	if rows > 0 {
		if vote {
			comment.Votes++
		} else {
			comment.Votes--
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontro un registro de comentario para asignarle el voto")
	}
	return
}
