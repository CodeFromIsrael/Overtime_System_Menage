package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func converterParameterId(r *http.Request, key string) (uint64, error) {

	parameter := r.PathValue(key)

	id, erro := strconv.ParseUint(parameter, 10, 64)

	if erro != nil {
		return 0, errors.New("erro a o converter o id da requisição para uint 64")

	}

	return id, nil
}

func readerOfRequestBody(r *http.Request) ([]byte, error) {

	return io.ReadAll(r.Body)
}

func ConverterJsonToStruct[T any](bodyRequest []byte) (T, error) {

	var model T

	if erro := json.Unmarshal(bodyRequest, &model); erro != nil {
		return model, erro
	}

	return model, nil
}

func readParameter(r *http.Request, srt string) (string, error) {

	dayAppointment := strings.ToLower(r.URL.Query().Get(srt))

	if dayAppointment == "" {

		return "", errors.New("problema na leitura do valor que esta no parametro ")
	}

	return dayAppointment, nil
}

func retriveUserInToken(r *http.Request) uint64 {

	userInToken := r.Context().Value("userid").(uint64)

	return userInToken
}

func converStringToFloat(str string) (float64, error) {

	float, erro := strconv.ParseFloat(str, 64)

	if erro != nil {
		return 0, erro
	}

	return float, nil
}
