package db

import "log"

// SelectUserID todo doc
func (r Repository) SelectUserID(token string) string {
	var key = ""
	var result string

	rows, err := r.Conn.Query("SELECT user_id FROM authtoken_token WHERE key = $1 LIMIT 1", token)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&key); err != nil {
			log.Fatal(err)
		}
		result = key
	}

	return result
}

// SelectUserPortinformer todo doc
func (r Repository) SelectUserPortinformer(userID string) string {
	var fkPortinformer = ""
	var result = ""

	rows, err := r.Conn.Query("SELECT fk_portinformer FROM portinformer_managers WHERE fk_user = $1 LIMIT 1", userID)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&fkPortinformer); err != nil {
			log.Println(err)
		}
		result = fkPortinformer
	}
	return result
}
