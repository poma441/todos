package repository

type AuthRepo struct {
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (r *AuthRepo) PlugFunc() {

}
