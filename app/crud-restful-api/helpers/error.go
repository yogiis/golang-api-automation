package helpers

import "database/sql"

func LogPanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback()
		LogPanicError(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		LogPanicError(errorCommit)
	}
}
