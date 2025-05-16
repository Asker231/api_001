package auth


//request
type (
	RegisterRequst struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequst struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
)
//response
type(
		RegisterResponse struct{
			
		}
)