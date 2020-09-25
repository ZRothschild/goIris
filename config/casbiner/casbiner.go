package casbiner

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// casbin 初始化
func NewCasbin(db *gorm.DB, conf, prefix, tableName string) (e *casbin.Enforcer, err error) {
	var (
		a *gormadapter.Adapter
	)

	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	// a, _ := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.

	// a, err := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	a, err = gormadapter.NewAdapterByDBUseTableName(db, prefix, tableName)
	if err != nil {
		return e, err
	}

	e, err = casbin.NewEnforcer(conf, a)

	// // Load the policy from DB.
	// err = e.LoadPolicy()
	//
	// // Check the permission.
	// b, err := e.Enforce("alice", "data1", "read")
	//
	// // Modify the policy.
	// // e.AddPolicy(...)
	// // e.RemovePolicy(...)
	//
	// // Save the policy back to DB.
	// e.SavePolicy()

	return e, err
}
