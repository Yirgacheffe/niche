package mocking

type Restorer func()

func (r Restorer) Restore() {
	r()
}

func Patch(dest, value interface{}) Restorer {

}
