// Now I'm wondering if this does anything?

package databases

import "gorm.io/gorm"

type IDatabase interface {
	ConnectDatabase() (*gorm.DB, error)
}
