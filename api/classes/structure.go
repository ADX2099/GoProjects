package classes

//CUANDO TIENE LA PRIMER LETRA EN MAYUSCULA SIGNIFICA QUE ES PUBLIC
//Acces login
type UserLogin struct {
	User     string `json:"user" validate:"required,email,min=3"`
	Password string `json:"pass"`
}

type RequestCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RequestLogin struct {
	Response RequestCode `json:"response"`
}

type Book struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}

type RequestBook struct {
	Response RequestCode `json:"response"`
	Books    []Book      `json:"books"`
}

type ResponseSecurity struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}
