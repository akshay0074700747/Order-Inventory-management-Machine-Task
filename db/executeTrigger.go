package db

import (
	"fmt"
	"io/ioutil"

	"gorm.io/gorm"
)

//setting up triggers to automate pricing and inventory management on order
func ExecTrigger(db *gorm.DB) error {
	// Read trigger creation statements from file
	triggerSQL, err := ioutil.ReadFile("triggers.sql")
	if err != nil {
		return err
	}

	// Execute the trigger creation statements
	if err = db.Exec(string(triggerSQL)).Error; err != nil {
		return err
	}
	fmt.Println("Triggers Created Successfully")

	return nil
}
