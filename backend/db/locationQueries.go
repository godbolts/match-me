package db

import (
	"fmt"
	"log"
)

func SetBrowser(userID, longitude, latitude string) error {
	userQuery := `
	UPDATE users 
	SET 
		browser_location = ST_SetSRID(ST_MakePoint($1, $2), 4326) 
	WHERE uuid = $3
`
	_, err := DB.Exec(userQuery, longitude, latitude, userID)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update city: %w", err)
	}

	return nil
}
