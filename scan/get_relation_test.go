package scan_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/og/gofree/scan"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type User struct {
	ID int64 `db:"id"`
	Name []byte `db:"name"`
	IsSuper int64 `db:"is_super"`
}
func (self User) TableName() string { return "user"}
type Book struct {
	ID int64 `db:"id"`
	Name []byte `db:"name"`
	Price []uint8 `db:"price"`
	UserID int64 `db:"user_id"`
	LastReadTime time.Time `db:"last_read_time"`
}
func (self Book) TableName() string { return "book"}
type UserAndBook struct {
	User User
	BookList []Book
}

func TestGetRelation(t *testing.T) {
	{
		userAndBook := UserAndBook{}
		assert.Equal(t, scan.Relation{
			Single:[]scan.RelationItem{
				{
					FieldIndex:0,
					TableName: "user",
					DBTag:map[int]string{
						0:"id",
						1:"name",
						2:"is_super",
					},
				},
			},
			Many:[]scan.RelationItem{
				{
					FieldIndex:1,
					TableName: "book",
					DBTag:map[int]string{
						0:"id",
						1:"name",
						2:"price",
						3:"user_id",
						4:"last_read_time",
					},
				},
			},
		}, scan.GetRelation(userAndBook))
	}
	{
		userAndBookList := []UserAndBook{}
		assert.Equal(t, scan.Relation{
			Single:[]scan.RelationItem{
				{
					FieldIndex:0,
					TableName: "user",
					DBTag:map[int]string{
						0:"id",
						1:"name",
						2:"is_super",
					},
				},
			},
			Many:[]scan.RelationItem{
				{
					FieldIndex:1,
					TableName: "book",
					DBTag:map[int]string{
						0:"id",
						1:"name",
						2:"price",
						3:"user_id",
						4:"last_read_time",
					},
				},
			},
		}, scan.GetSliceRelation(userAndBookList))
	}
}
