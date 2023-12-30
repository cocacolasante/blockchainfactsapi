package postgresrepo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/cocacolasante/blockchainfacts/models"
)

const dbTimeout = time.Second * 3

type PostgresRepo struct {
	DB *sql.DB
}

func (db *PostgresRepo) Connection() *sql.DB {
	return db.DB
}

func (db *PostgresRepo) OneFact(id int) (*models.BCFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	count := db.getFactCount()
	if count == 0 {
		return nil, errors.New("No Facts In Database")
	}
	
	factNum := rand.Intn(count)
	log.Println("fact num:", factNum)

	fact := &models.BCFact{}

	stmt := `SELECT fact_id, fact_text FROM facts WHERE fact_id = $1;`

	row := db.DB.QueryRowContext(ctx, stmt, id)
	
	err := row.Scan(&fact.ID, &fact.Fact)
	if err != nil {
		log.Println("repo 1", err)
		log.Fatal(err)
	}

	return fact, nil
	
	

}
func (db *PostgresRepo) OneFactRandom() (*models.BCFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	count := db.getFactCount()
	if count == 0 {
		return nil, errors.New("No Facts In Database")
	}
	
	factNum := rand.Intn(count)
	log.Println("fact num:", factNum)

	fact := &models.BCFact{}

	stmt := `SELECT fact_id, fact_text FROM facts WHERE fact_id = $1;`

	row := db.DB.QueryRowContext(ctx, stmt, factNum)
	
	err := row.Scan(&fact.ID, &fact.Fact)
	if err != nil {
		log.Println("repo 1", err)
		log.Fatal(err)
	}

	return fact, nil
	
	

}


func (db *PostgresRepo) getFactCount() int {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT COUNT(*) AS total_facts FROM facts`

	var totalCount int
	err := db.DB.QueryRowContext(ctx, stmt).Scan(&totalCount)
	if err != nil {
		log.Println("repo 2", err)
		log.Fatal("Cannot receive count:", err)
		
	}

	return totalCount

}

func (db *PostgresRepo) AllFacts()([]*models.BCFact, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var facts []*models.BCFact

	stmt := `SELECT fact_id, fact_text FROM facts;`

	row, err := db.DB.QueryContext(ctx, stmt)
	if err != nil {
		
		return nil, err
	}

	defer row.Close()
	for row.Next(){
		var fact models.BCFact
		err := row.Scan(&fact.ID, &fact.Fact)
		if err !=nil {
			return nil, err
		}
		facts = append(facts, &fact)

	}

	return facts, nil
}
