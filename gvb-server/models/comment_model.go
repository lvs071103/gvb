package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  // 子评论
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父级评论id
	Content            string          `gorm:"size:256" json:"content"`                         // 评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`              // 点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`           // 评论数
	Article            ArticleModel    `gorm:"foreignKey: ArticleID" json:"-"`                  // 关联文章
	ArticleID          uint            `json:"article_id"`                                      // 文章id
	User               UserModel       `json:"user"`                                            // 关联的用户
	UserID             uint            `json:"user_id"`                                         // 评论用户
}
