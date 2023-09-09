package main

// tipe data username dan password yang tidak sesuai yaitu menggunakan int, seharusnya menggunakan string
type user struct {
	id int
	username int
	password int
}

// penamaan variabel "t" dari struct userservice yang tidak mudah dipahami dan terlalu singkat sehingga tidak menjelaskan konteks
// penamaan struct "userservice" yang tidak sesuai dengan naming convention seperti camel case dan snake case
type userservice struct {
	t []user
}

// penamaan parameter "u" yang tidak mudah dipahami dan terlalu singkat sehingga tidak menjelaskan konteks
// penamaan method "getallusers" yang tidak sesuai dengan naming convention seperti camel case dan snake case
func (u userservice) getallusers() []user {
	return u.t
}

// penamaan parameter "u" dan value dari looping "r" yang tidak mudah dipahami dan terlalu singkat sehingga tidak menjelaskan konteks
// penamaan method "getuserbyid" yang tidak sesuai dengan naming convention seperti camel case dan snake case
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}