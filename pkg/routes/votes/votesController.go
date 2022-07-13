package votes

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/dto"
	"io/ioutil"
	"log"
	"net/http"
)

func Create(writer http.ResponseWriter, req *http.Request) {
	var voteDTO dto.Vote

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE CREATE!", err)
		return
	}

	err = json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return
	}

	voteDAO, err := dao.CreateVote(voteDTO)
	if err != nil {
		log.Println("FAILED to create new vote", err)
		return
	}

	voteDTO = dto.Vote{
		Success: true,
		Message: "",
		Data:    dto.VoteData{UUID: voteDAO.UUID, UserId: voteDAO.UserId, CountryId: voteDAO.CountryId, Costume: voteDAO.Costume, Song: voteDAO.Song, Performance: voteDAO.Performance, Props: voteDAO.Props},
	}

	json.NewEncoder(writer).Encode(voteDTO)
}
