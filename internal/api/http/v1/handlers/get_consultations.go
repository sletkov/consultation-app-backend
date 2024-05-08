package handlers

import (
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetConsultations(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//fmt.Println(vars)
	//page := r.FormValue("page")
	//fmt.Println(page)
	//page := vars["page"]
	//pageInt, err := strconv.Atoi(page)
	//if err != nil {
	//	utils.ErrRespond(w, r, http.StatusInternalServerError, err)
	//	return
	//}

	//limit := vars["limit"]
	//limit := r.FormValue("limit")
	//fmt.Println(limit)
	//limitInt, err := strconv.Atoi(limit)
	//if err != nil {
	//	utils.ErrRespond(w, r, http.StatusInternalServerError, err)
	//	return
	//}

	consultations, err := c.consultationService.GetConsultations(r.Context())
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get consultations", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	//start := 1*pageInt - 1
	//end := 1*limitInt - 1
	//
	//if end > len(consultations) {
	//	end = len(consultations) - 1
	//}

	utils.Respond(w, r, http.StatusOK, consultations)
}
