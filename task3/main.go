package main

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）
， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
题目3：钩子函数,
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount uint   `gorm:"default:0"`
}
type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"not null"`
	UserID       uint      `gorm:"not null"`
	Comments     []Comment `gorm:"foreignKey:PostID"`
	CommentCount uint      `gorm:"default:0"`
	HasComment   bool      `gorm:"default:false"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.ID).Update("post_count", gorm.Expr("post_count + 1")).Error
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	//查询相关文章评论还有几个
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	fmt.Println("Comment  AfterCreate", count, c.PostID)
	return tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(map[string]interface{}{
		"comment_count": count,
		"has_comment":   count > 0,
	}).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {

	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}
	fmt.Println("Comment  AfterCreate", count, c.ID, c.Content)

	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{"comment_count": count, "has_comment": count > 0}).Error

}

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:as5015520@tcp(127.0.0.1:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         171,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatal("数据库连接失败")
	}

	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	return db
}

func insertData(db *gorm.DB) {
	var count int64
	db.Model(&Comment{}).Count(&count)
	if count > 0 {
		return
	}

	users := []User{
		{Username: "小杰",
			Posts: []Post{
				{
					Title: "Gorm 指南",
					Comments: []Comment{
						{
							Content: "写的很好,下次继续",
						},
					},
				},
			}},
		{
			Username: "小熊",
			Posts: []Post{
				{
					Title: "修炼手册",
					Comments: []Comment{
						{
							Content: "如何修炼",
						},
					},
				},
			},
		},
	}

	db.Create(&users)
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func AssicationQuery(db *gorm.DB) {
	var user User
	db.Preload("Posts.Comments").Where("Username = ?", "小熊").Find(&user)
	fmt.Printf("博客名称:%v", user.Username)
	for _, V := range user.Posts {
		fmt.Printf("拥有的文章：%v", V.Title)
		for _, R := range V.Comments {
			fmt.Printf("%v的评论:%v", V.Title, R.Content)
		}
	}

}

// 查询评论数量最多的文章
func queryMostCommentedPost(db *gorm.DB) {

	var post Post

	// 使用子查询和JOIN获取评论最多的文章
	err := db.Model(&Post{}).
		Select("t_posts.*, COUNT(t_comments.id) as comment_count").
		Joins("LEFT JOIN t_comments ON t_comments.post_id = t_posts.id").
		Group("t_posts.id").
		Order("comment_count DESC").
		First(&post).
		Error
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	// 加载评论详情
	db.Preload("Comments").First(&post, post.ID)

	fmt.Printf("\n评论最多的文章: %s (评论数: %d)\n", post.Title, len(post.Comments))
	for _, comment := range post.Comments {
		fmt.Printf("  评论: %s\n", comment.Content)
	}
}

func main() {
	db := DBInit()
	insertData(db)

	// newcom := Comment{
	// 	Content: "小杰洗澡视频",
	// 	PostID:  10,
	// }

	// db.Create(&newcom)

	AssicationQuery(db)
	queryMostCommentedPost(db)

	var comments []Comment
	db.Where("content = ?", "小杰洗澡视频").Find(&comments)
	for _, c := range comments {
		db.Delete(&c) // 此时会正常触发钩子
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

}
