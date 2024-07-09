package handler

import (
	"github.com/egorus1442/Report-Generation-Microservice/pkg/service"
	"net/http"
	"strconv"

	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input rgm.SalesList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.SalesList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []rgm.SalesList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.SalesList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getReport(c *gin.Context) {
	AllId, err := h.services.SalesList.GetAllUserId()
	if err != nil {
		panic(err)
	}
	var AllSales [][]rgm.SalesPdf
	for _, id := range AllId {
		sales, err := h.services.SalesList.GetAll(id)
		if err != nil {
			panic(err)
		}
		name, err := h.services.SalesList.GetUserNameById(id)
		if err != nil {
			panic(err)
		}
		var s []rgm.SalesPdf
		for _, p := range sales {
			var sale = rgm.SalesPdf{
				Saller: name,
				Title:  p.Title,
				Amount: p.Amount,
				Price:  p.Price,
			}
			s = append(s, sale)
		}

		AllSales = append(AllSales, s)
	}
	BiggestSale, err := h.services.SalesList.GetBiggerSale()
	if err != nil {
		panic(err)
	}
	var allBiggestSales []rgm.SalesPdf
	for _, p := range BiggestSale {
		var sale = rgm.SalesPdf{
			Saller: strconv.Itoa(p.Id),
			Title:  p.Title,
			Amount: p.Amount,
			Price:  p.Price,
		}
		allBiggestSales = append(allBiggestSales, sale)
	}

	LowerSale, err := h.services.SalesList.GetLowerSale()
	if err != nil {
		panic(err)
	}
	var allLowerSales []rgm.SalesPdf
	for _, p := range LowerSale {
		var sale = rgm.SalesPdf{
			Saller: strconv.Itoa(p.Id),
			Title:  p.Title,
			Amount: p.Amount,
			Price:  p.Price,
		}
		allLowerSales = append(allLowerSales, sale)
	}

	report, err := service.PdfMaker(AllSales, allBiggestSales, allLowerSales)
	if err != nil {
		panic(err)
	}

	c.Header("Content-Disposition", "attachment; filename=example.pdf")
	c.Data(http.StatusOK, "application/pdf", report)

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
