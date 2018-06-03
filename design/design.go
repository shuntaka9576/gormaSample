package design // The convention consists of naming the design
// package "design"
import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("celler", func() {
	Description("celler Model")
	// Mysqlを使う
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQLのリレーションナルデータベース")
		// accountsテーブルのModelなら、Account
		Model("Account", func() {
			// MediaTypeで作成したAccountにマッピングする
			RendersTo(Account)
			Description("celler account")
			// PrimaryKeyの設定
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("email", gorma.String)
			// timestamp系の定義
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
	})
})
