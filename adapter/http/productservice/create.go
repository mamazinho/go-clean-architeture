package productservice

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/mamazinho/go-clean-architeture/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		reader := strings.NewReader(err.Error())
		io.Copy(response, reader)
		return
	}

	product, err := service.usecase.Create(productRequest)

	if err != nil {
		response.WriteHeader(500)
		reader := strings.NewReader(err.Error())
		io.Copy(response, reader)
		return
	}

	json.NewEncoder(response).Encode(product)
}
