package mainDb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/techrail/ground/logger"
)

// PublicGooseDbVersion struct corresponds to the goose_db_version table in public schema of the DB
// Table Comment:
type PublicGooseDbVersion struct {
	Id        int32        `db:"id"`
	VersionId int64        `db:"version_id"`
	IsApplied bool         `db:"is_applied"`
	Tstamp    sql.NullTime `db:"tstamp"`
}

func (gooseDbVersion *PublicGooseDbVersion) baseValidation() error {
	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) commonValidation() error {
	err := gooseDbVersion.baseValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) validateForInsert() error {
	err := gooseDbVersion.commonValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) validateForUpdate() error {
	err := gooseDbVersion.commonValidation()
	if err != nil {
		return err
	}
	// More code to be written here for validation
	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) Insert() error {
	var err error
	err = gooseDbVersion.validateForInsert()
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO public.goose_db_version (
			"version_id", 
			"is_applied", 
			"tstamp"
		) VALUES (
			$1, 
			$2, 
			$3
		) RETURNING "id"`

	resultRow := MainDb.QueryRowx(insertQuery,
		gooseDbVersion.VersionId,
		gooseDbVersion.IsApplied,
		gooseDbVersion.Tstamp,
	)

	if resultRow.Err() != nil {
		errMsg := fmt.Sprintf("E#00YQ33 - Could not insert into database: %v", resultRow.Err())
		logger.Println(errMsg)
		return errors.New(errMsg)
	}

	var insertedId int32

	err = resultRow.Scan(&insertedId)
	if err != nil {
		return fmt.Errorf("E#00YQ34 - Scan failed. Error: %v", err)
	}

	gooseDbVersion.Id = insertedId
	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) Update() error {
	err := gooseDbVersion.validateForUpdate()
	if err != nil {
		return err
	}
	updateQuery := `UPDATE public.goose_db_version SET
			"version_id" = $1,
			"is_applied" = $2,
			"tstamp" = $3
		WHERE 
			"id" = $4`

	_, err = MainDb.Exec(updateQuery,
		gooseDbVersion.VersionId,
		gooseDbVersion.IsApplied,
		gooseDbVersion.Tstamp,
		gooseDbVersion.Id)
	if err != nil {
		return fmt.Errorf("E#00YQ35 - Could not update PublicGooseDbVersion in database: %v", err)
	}

	return nil
}

func (gooseDbVersion *PublicGooseDbVersion) Delete() error {
	_, err := MainDb.Exec(`DELETE FROM public.goose_db_version WHERE id = $1`, gooseDbVersion.Id)
	if err != nil {
		return fmt.Errorf("E#00YQ36 - Could not delete GooseDbVersion from database: %v", err)
	}

	return nil
}

// =============================================
// Table methods end here. Dao functions below
// =============================================

type PublicGooseDbVersionDao struct{}

func NewPublicGooseDbVersionDao() *PublicGooseDbVersionDao {
	return &PublicGooseDbVersionDao{}
}
func (gooseDbVersionDao *PublicGooseDbVersionDao) GetFromDbById(id int32, getFromMainDb ...bool) (PublicGooseDbVersion, error) {
	var err error
	query := `SELECT * FROM public.goose_db_version WHERE "id" = $1`
	gooseDbVersion := PublicGooseDbVersion{}

	if len(getFromMainDb) > 0 && getFromMainDb[0] == true {
		err = MainDb.Get(&gooseDbVersion, query, id)
	} else {
		err = MainDbReader.Get(&gooseDbVersion, query, id)
	}

	if err == sql.ErrNoRows {
		return gooseDbVersion, err
	}
	if err != nil {
		errMsg := fmt.Sprintf("E#00YQ37 - Could not load GooseDbVersion by Id Error: %v", err)
		logger.Println(errMsg)
		return gooseDbVersion, errors.New(errMsg)
	}
	return gooseDbVersion, nil
}

// MAGIC COMMENT (DO NOT EDIT): Please write any custom code only below this line.

// Make sure code below is valid before running code generator else the generator will fail

