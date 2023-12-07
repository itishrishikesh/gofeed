package mainDb

// PublicSchema struct corresponds to the public schema of the DB
type PublicSchema struct {
	GooseDbVersionDao *PublicGooseDbVersionDao // Dao for goose_db_version
	UserDao           *PublicUserDao           // Dao for users
}

var Public PublicSchema

func init() {
	Public = PublicSchema{
		GooseDbVersionDao: NewPublicGooseDbVersionDao(), // Dao for goose_db_version
		UserDao:           NewPublicUserDao(),           // Dao for users
	}
}

// MAGIC COMMENT (DO NOT EDIT): Please write any custom code only below this line.

// Make sure code below is valid before running code generator else the generator will fail
