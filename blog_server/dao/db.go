package dao

// Blog 将定义的模型放在这个目录中 其次，对于模型所有的原子操作，即增删改查等操作也放在此处
type Blog struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"Content"`
}

func (t *Blog) TableName() string {
	return "blog_tab"
}

//func InitModel() {
//	DB.AutoMigrate(&Blog{})
//}

func CreateBlog(blog *Blog) (err error) {
	err = DB.Create(&blog).Error
	return err
}

func GetBlogs() (todoList []Blog, err error) { //
	err = DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	} else {
		return
	}
}

//func GetATodoById(id string) (todo *Todo, err error) {
//	todo = new(Todo) //此处要加这句话，否则无法更新。？？
//	err = DB.Where("id=?", id).First(todo).Error
//	if err != nil {
//		return nil, err
//	} else {
//		return
//	}
//}
//
//func UpdateATodo(todo *Todo) (err error) {
//	err = DB.Save(&todo).Error
//	return nil
//}
//
//func DeleteATodo(id string) (err error) {
//	err = DB.Where("id = ?", id).Delete(&Todo{}).Error
//	return
//}
