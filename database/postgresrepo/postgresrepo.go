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

	stmt := `INSERT INTO facts (fact_text) VALUES ($1) RETURNING fact_id;`

	var factID int
	err := db.DB.QueryRowContext(ctx, stmt, fact).Scan(&factID)
	if err != nil {
		
		return nil, err
	}

	retFact := &models.BCFact{
		ID:   factID,
		Fact: fact,
	}

	return retFact, nil
}

func(db  *PostgresRepo) DeleteFact(factId int) (error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE from facts where fact_id = $1;`

	
	
	sqlRes, err := db.DB.ExecContext(ctx, stmt, factId)

	if err != nil {
		
		return err
	}
	rowsAffected, _ := sqlRes.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("No records deleted")
	}
	

	return nil

}

func(db *PostgresRepo) UpdateFact(text string, id int) (*models.BCFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update facts set fact_text = $1 where fact_id = $2 RETURNING fact_id, fact_text;`
	
	result, err := db.DB.ExecContext(ctx, stmt, text, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return nil, errors.New("incorrect amount of rows affected")
	}

	fact, err := db.OneFact(id)
	if err != nil {
		return nil, err
	}
	
		
	return fact, nil 
}
