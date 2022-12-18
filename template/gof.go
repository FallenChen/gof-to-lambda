package main

type Resource struct {
	run func()
}

func (r *Resource) Init() {
	println("Resource created")
}

func (r *Resource) Dispose() {
	println("Resource disposed")
}

func (r *Resource) Execute() {
	r.Init()
	r.run()
	r.Dispose()
}

type ResourceUser struct {
	Resource
}

func NewResourceUser() *ResourceUser {
	resourceUser := ResourceUser{}
	resourceUser.run = resourceUser.UseResource
	return &resourceUser
}

func (user *ResourceUser) UseResource() {
	println("Resource used")
}

type ResourceEmployer struct {
	Resource
}

func NewResourceEmployer() *ResourceEmployer {
	resourceEmployer := ResourceEmployer{}
	resourceEmployer.run = resourceEmployer.EmployResource
	return &resourceEmployer
}

func (user *ResourceEmployer) EmployResource() {
	println("Resource employed")
}

func main() {
	user := NewResourceUser()
	user.Execute()

	employer := NewResourceEmployer()
	employer.Execute()
}
