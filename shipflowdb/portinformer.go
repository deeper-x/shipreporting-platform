package db

import "log"

// SelectPortinformerName returns portinformer name
func (r Repository) SelectPortinformerName(portinformerID interface{}) string {
	var portinformerName = ""
	var result = ""

	rows, err := r.Conn.Query("SELECT portinformer_code FROM portinformers WHERE id_portinformer = $1", portinformerID)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&portinformerName); err != nil {
			log.Println(err)
		}

		result = portinformerName
	}

	return result
}
