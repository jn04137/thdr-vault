package repositories

import (
	"log"
	"database/sql"
)

type SecretRepo struct {
	dbConn *sql.DB
}

type Secret struct {
	Key string
	Secret string
}

func CreateSecretRepo(dbConn *sql.DB) *SecretRepo {
	return &SecretRepo{
		dbConn: dbConn,
	}
}

func (secretRepo *SecretRepo) GetSecrets() ([]Secret, error) {
	dbConn := secretRepo.dbConn
	getSecrets := `select key, value from secret;`
	rows, err := dbConn.Query(getSecrets)
	if err != nil {
		log.Printf("failed at query: %v", err.Error())
		return nil, err
	}

	var secrets []Secret
	for rows.Next() {
		var secret Secret
		if rowErr := rows.Scan(&secret.Key, &secret.Secret); err != nil {
			log.Printf("Failed to convert sql result to list: %v", rowErr.Error())
			return nil, rowErr
		}
		secrets = append(secrets, secret)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error in rows: %v", err.Error())
		return nil, err
	}
	return secrets, nil
}

func (secretRepo *SecretRepo) PutSecret(key string, value string) error {
	dbConn := secretRepo.dbConn

	putSecretQuery := `INSERT INTO secret(key, value) VALUES(?,?)`
	_, err := dbConn.Exec(putSecretQuery, key, value)
	return err
}
