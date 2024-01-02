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


func (db *PostgresRepo) AddFact(fact string) (*models.BCFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	total := db.getFactCount()

	stmt := `INSERT INTO facts (fact_id, fact_text) VALUES ($1, $2) RETURNING fact_id;`

	var factID int
	err := db.DB.QueryRowContext(ctx, stmt, total+1, fact).Scan(&factID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	retFact := &models.BCFact{
		ID:   factID,
		Fact: fact,
	}

	return retFact, nil
}

func(db  *PostgresRepo) DeleteFact(factId int) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `DELETE from facts where fact_id = $1;`

	var factID int
	log.Println(factID)
	_, err := db.DB.ExecContext(ctx, stmt, factID)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}