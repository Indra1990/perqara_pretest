package vendingmachine

import (
	"net/http"
	"perqara_testing/api/vendingmachine/usecase"
	"perqara_testing/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VendingMachineHttpRouterRegistry struct {
	serviceVendingMachine usecase.VendingMachineService
}

func NewTransactionHttpRouterRegistry(serviceVendingMachine usecase.VendingMachineService) *VendingMachineHttpRouterRegistry {
	return &VendingMachineHttpRouterRegistry{
		serviceVendingMachine: serviceVendingMachine,
	}
}

func (p *VendingMachineHttpRouterRegistry) Test(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "testing doang kok ...")
}

// List Item godoc
//
// @Summary		Get List Item
// @Description	get data item list
// @Tags 		Test-Perqara
// @Produce		json
// @Accept		json
// @Success		200		{object}	utils.ResponseBody{data=[]usecase.VendingMachine}
// @Failure		400		{object}	utils.ResponseBody
// @Router			/items [get]
func (p *VendingMachineHttpRouterRegistry) ListItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vendingMachines, err := p.serviceVendingMachine.ItemAll(ctx)
	if err != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, err.Error(), nil)
		return
	}
	utils.RenderJSONWithData(w, r, http.StatusOK, vendingMachines, nil, nil)
}

// Create Item godoc
//
// @Summary		Create Item
// @Description	add data item
// @Tags 		Test-Perqara
// @Produce		json
// @Accept		json
// @param 		Item 	 body   usecase.VendingMachineRequestCommand true "Request"
// @Success		200		{object}	utils.ResponseBody
// @Failure		400		{object}	utils.ResponseBody
// @Router		/item/create [post]
func (p *VendingMachineHttpRouterRegistry) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cmd := usecase.VendingMachineRequestCommand{}

	if decodeErr := render.DecodeJSON(r.Body, &cmd); decodeErr != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, decodeErr.Error(), nil)
		return
	}

	if err := p.serviceVendingMachine.CreateItem(ctx, &cmd); err != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, err.Error(), nil)
		return
	}

	utils.RenderJSONWithData(w, r, http.StatusCreated, "success created", nil, nil)
}

// Detail Item godoc
//
// @Summary		Detail Item
// @Description	Detail data item
// @Tags 		Test-Perqara
// @Produce		json
// @Accept		json
// @Param		itemId		path		string		true 	"itemId"
// @Success		200		{object}	utils.ResponseBody{data=usecase.VendingMachine}
// @Failure		400		{object}	utils.ResponseBody
// @Router			/item/{itemId} [get]
func (p *VendingMachineHttpRouterRegistry) DetailItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	itemId := chi.URLParam(r, "itemId")

	detailItem, detailItemErr := p.serviceVendingMachine.DetailItem(ctx, itemId)
	if detailItemErr != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, detailItemErr.Error(), nil)
		return
	}

	utils.RenderJSONWithData(w, r, http.StatusOK, detailItem, nil, nil)
}

// Update Item godoc
//
// @Summary		Update Item
// @Description	Update data item
// @Tags 		Test-Perqara
// @Produce		json
// @Accept		json
// @Param		itemId		path		string		true 	"itemId"
// @param 		Item 	 body   usecase.VendingMachineUpdateCommand true "Request"
// @Success		200		{object}	utils.ResponseBody
// @Failure		400		{object}	utils.ResponseBody
// @Router			/item/{itemId} [put]
func (p *VendingMachineHttpRouterRegistry) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cmd := usecase.VendingMachineUpdateCommand{}
	itemId := chi.URLParam(r, "itemId")

	if decodeErr := render.DecodeJSON(r.Body, &cmd); decodeErr != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, decodeErr.Error(), nil)
		return
	}

	cmd.ID = itemId

	if err := p.serviceVendingMachine.UpdateItem(ctx, &cmd); err != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, err.Error(), nil)
		return
	}

	utils.RenderJSONWithData(w, r, http.StatusOK, "success updated", nil, nil)

}

// Delete Item godoc
//
// @Summary		Delete Item
// @Description	Delete data item
// @Tags 		Test-Perqara
// @Produce		json
// @Accept		json
// @Param		itemId		path		string		true 	"itemId"
// @Success		200		{object}	utils.ResponseBody
// @Failure		400		{object}	utils.ResponseBody
// @Router			/item/{itemId} [delete]
func (p *VendingMachineHttpRouterRegistry) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	itemId := chi.URLParam(r, "itemId")

	if err := p.serviceVendingMachine.DeleteItem(ctx, itemId); err != nil {
		utils.RenderJSONWithData(w, r, http.StatusBadRequest, nil, err.Error(), nil)
		return
	}

	utils.RenderJSONWithData(w, r, http.StatusOK, "success deleted", nil, nil)

}
