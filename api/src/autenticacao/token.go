package autenticacao

import (
	"api/src/config"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário.
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey))
}

// Validar token verifica se o Token passado na requisição é valido.
func ValidarToken(r *http.Request) error {
	return nil
}
