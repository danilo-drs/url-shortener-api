package model

import (
	"database/sql"
	"fmt"
	"meli-api/repository"
	"meli-api/utils"
	"os"
)

type ShortUrl struct {
	Key         string `json:"shortUrlKey"`
	OriginalURL string `json:"url"`
	ShortUrl    string `json:"shortUrl"`
	CreateAt    string `json:"createAt"`
	CreatedBy   string `json:"createdBy"`
	UpdateAt    string `json:"updateAt"`
	UpdatedBy   string `json:"updatedBy"`
}

func (s *ShortUrl) saveShortUrl() error {
	//TODO: garantir que todos os campos estão preenchidos
	const query = `INSERT INTO short_url (key, url, short_url, create_at, created_by, update_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := repository.DB.Exec(query, s.Key, s.OriginalURL, s.ShortUrl, s.CreateAt, s.CreatedBy, s.UpdateAt, s.UpdatedBy)
	if err != nil {
		fmt.Println("Error saving short URL: ", err)
		return err
	}
	return nil
}

// CheckKeyExist checks if the key exists in the database
func checkKeyExist(key string) (bool, error) {
	// get the short URL from the database
	query := "SELECT true FROM short_url WHERE key = $1"
	var exists bool
	err := repository.DB.QueryRow(query, key).Scan(&exists)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, err
	}
	return exists, nil
}

// GenerateShortUrl generates a unique short URL
func (s *ShortUrl) GenerateShortUrl() (string, error) {
	var exists bool = true
	var err error

	// keep generating a short URL until it is unique
	for exists {
		s.Key = utils.GenerateShortKey()
		// check if the key exists
		exists, err = checkKeyExist(s.Key)
		if err != nil {
			fmt.Println("Error checking if the key exists: ", err)
			return "", err
		}
		if !exists {
			// get the short URL domain from the environment variables
			shortUrlDomain := os.Getenv("SHORT_URL_DOMAIN")

			// set the short URL
			s.ShortUrl = shortUrlDomain + s.Key
			err = s.saveShortUrl()
			if err != nil {
				return "", err
			}
		}
	}
	return s.ShortUrl, nil
}

// FillFromKey fills the short URL from the key
func (s *ShortUrl) FillFromKey(key string) (bool, error) {
	query := "SELECT key, url, short_url, create_at, created_by, update_at, updated_by FROM short_url WHERE key = $1"
	err := repository.DB.QueryRow(query, key).Scan(&s.Key, &s.OriginalURL, &s.ShortUrl, &s.CreateAt, &s.CreatedBy, &s.UpdateAt, &s.UpdatedBy)
	var found bool = false
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return found, nil
		}
		fmt.Println("Error getting short URL with key: ", err)
		return found, err
	}
	return true, nil
}

// GetAllShortUrls gets all the short URLs
func GetAllShortUrls() ([]ShortUrl, error) {
	query := "SELECT key, url, short_url, create_at, created_by, update_at, updated_by FROM short_url"
	rows, err := repository.DB.Query(query)
	if err != nil {
		fmt.Println("Error getting all short URLs: ", err)
		return nil, err
	}
	defer rows.Close()
	return rowsToShortUrls(rows), nil
}

// rowsToShortUrls converts the rows to short URLs
func rowsToShortUrls(rows *sql.Rows) []ShortUrl {
	var shortUrls []ShortUrl
	for rows.Next() {
		var s ShortUrl
		err := rows.Scan(&s.Key, &s.OriginalURL, &s.ShortUrl, &s.CreateAt, &s.CreatedBy, &s.UpdateAt, &s.UpdatedBy)
		if err != nil {
			fmt.Println("Error scanning short URL: ", err)
			return nil
		}
		shortUrls = append(shortUrls, s)
	}
	return shortUrls
}

// Update updates the short URL
func (s *ShortUrl) Update() error {
	query := "UPDATE short_url SET url = $1, update_at = $2, updated_by = $3 WHERE key = $4"
	_, err := repository.DB.Exec(query, s.OriginalURL, s.UpdateAt, s.UpdatedBy, s.Key)
	if err != nil {
		fmt.Println("Error updating short URL: ", err)
		return err
	}
	return nil
}

// Delete deletes the short URL
func (s *ShortUrl) Delete() error {
	query := "DELETE FROM short_url WHERE key = $1"
	_, err := repository.DB.Exec(query, s.Key)
	if err != nil {
		fmt.Println("Error deleting short URL: ", err)
		return err
	}
	return nil
}

// String return the struct as a json string
func (s *ShortUrl) String() string {
	return fmt.Sprintf(`{"key": "%s", "url": "%s", "shortUrl": "%s", "createAt": "%s", "createdBy": "%s", "updateAt": "%s", "updatedBy": "%s"}`, s.Key, s.OriginalURL, s.ShortUrl, s.CreateAt, s.CreatedBy, s.UpdateAt, s.UpdatedBy)
}
