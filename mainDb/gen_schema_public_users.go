package mainDb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/techrail/ground/logger"
	"time"
)

// PublicUser struct corresponds to the users table in public schema of the DB
// Table Comment:
type PublicUser struct {
	Id        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
}

func (user *PublicUser) baseValidation() error {
	return nil
}

func (user *PublicUser) commonValidation() error {
	err := user.baseValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (user *PublicUser) validateForInsert() error {
	err := user.commonValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (user *PublicUser) validateForUpdate() error {
	err := user.commonValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (user *PublicUser) Insert() error {
	var err error
	err = user.validateForInsert()
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO public.users (
			"id", 
			"name"
		) VALUES (
			$1, 
			$2
		) RETURNING "id", "created_at", "updated_at"`

	resultRow := MainDb.QueryRowx(insertQuery,
		user.Id,
		user.Name,
	)

	if resultRow.Err() != nil {
		errMsg := fmt.Sprintf("E#00YQ38 - Could not insert into database: %v", resultRow.Err())
		logger.Println(errMsg)
		return errors.New(errMsg)
	}

	var insertedId string
	var insertedCreatedAt time.Time
	var insertedUpdatedAt time.Time

	err = resultRow.Scan(&insertedId, &insertedCreatedAt, &insertedUpdatedAt)
	if err != nil {
		return fmt.Errorf("E#00YQ39 - Scan failed. Error: %v", err)
	}

	user.Id = insertedId
	user.CreatedAt = insertedCreatedAt
	user.UpdatedAt = insertedUpdatedAt
	return nil
}

func (user *PublicUser) Update() error {
	err := user.validateForUpdate()
	if err != nil {
		return err
	}
	updateQuery := `UPDATE public.users SET
			"name" = $1
		WHERE 
			"id" = $2`

	_, err = MainDb.Exec(updateQuery,
		user.Name,
		user.Id)
	if err != nil {
		return fmt.Errorf("E#00YQ3A - Could not update PublicUser in database: %v", err)
	}

	return nil
}

func (user *PublicUser) Delete() error {
	_, err := MainDb.Exec(`DELETE FROM public.users WHERE id = $1`, user.Id)
	if err != nil {
		return fmt.Errorf("E#00YQ3B - Could not delete User from database: %v", err)
	}

	return nil
}

// =============================================
// Table methods end here. Dao functions below
// =============================================

type PublicUserDao struct{}

func NewPublicUserDao() *PublicUserDao {
	return &PublicUserDao{}
}
func (userDao *PublicUserDao) GetFromDbById(id string, getFromMainDb ...bool) (PublicUser, error) {
	var err error
	query := `SELECT * FROM public.users WHERE "id" = $1`
	user := PublicUser{}

	if len(getFromMainDb) > 0 && getFromMainDb[0] == true {
		err = MainDb.Get(&user, query, id)
	} else {
		err = MainDbReader.Get(&user, query, id)
	}

	if err == sql.ErrNoRows {
		return user, err
	}
	if err != nil {
		errMsg := fmt.Sprintf("E#00YQ3C - Could not load Users by Id Error: %v", err)
		logger.Println(errMsg)
		return user, errors.New(errMsg)
	}
	return user, nil
}

// MAGIC COMMENT (DO NOT EDIT): Please write any custom code only below this line.

// Make sure code below is valid before running code generator else the generator will fail

