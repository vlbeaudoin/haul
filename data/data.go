package data

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Assembly struct {
	gorm.Model
	ID         uint64     `json:"id"`
	Label      string     `json:"label"`
	State      string     `json:"state"`
	Final      bool       `json:"final"`
	AssemblyID uint64     `json:"assembly_id"`
	Assets     []Asset    `json:"assets"`
	Assemblies []Assembly `json:"assemblies"`
}

type Asset struct {
	gorm.Model
	ID         uint64 `json:"id"`
	AssemblyID uint64 `json:"assembly_id"`
	Label      string `json:"label"`
	State      string `json:"state"`
	Final      bool   `json:"final"`
}

func OpenDatabase() error {
	var err error

	var dialector gorm.Dialector

	switch t := viper.GetString("db.type"); t {
	case "sqlite":
		log.Println("Using driver gorm.io/driver/sqlite")

		db_path := viper.GetString("db.path")

		if db_path == "" {
			log.Fatal("No valid database file found in `--db-path` or `db.path`.")
		}

		log.Println("Using database file:", db_path)

		dialector = sqlite.Open(db_path)
	default:
		log.Fatalf("Unrecognized database driver requested (%s).\n", t)
	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func MigrateDatabase() {
	err := db.AutoMigrate(&Assembly{}, &Asset{})
	if err != nil {
		log.Fatal(err)
	}
}

// CRUD

// Create
func InsertAsset(asset Asset) {
	db.Create(&asset)
}

func InsertAssembly(assembly Assembly) {
	db.Create(&assembly)
}

/* REFERENCE
func InsertTasks(tasks []*Task) error {
	if len(tasks) == 0 {
		return errors.New("Cannot insert empty batch of tasks.")
	}

	for _, task := range tasks {
		task.ID = 0
	}

	db.CreateInBatches(&tasks, 500)

	return nil
}
*/

// TODO func InsertAssets

// TODO func InsertAssemblies

// Read

/* REFERENCE
func FindTaskByID(taskID uint64) (task Task, err error) {
	result := db.First(&task, taskID)
	return task, result.Error
}

func FindStepsByTaskID(taskID uint64) (steps []Step, err error) {
	result := db.Where("task_id = ?", taskID).Find(&steps)
	return steps, result.Error
}
*/

// TODO func FindAssetByID

// TODO func FindAssemblyByID

/* REFERENCE
func ListTasks() ([]Task, error) {
	var tasks []Task

	result := db.Model(&Task{}).Find(&tasks)

	return tasks, result.Error
}
*/

// TODO func ListAssets

// TODO func ListAssemblies

// Update

// Delete

/* REFERENCE for TODO InsertAssemblyWithAssetsAndAssemblies (name pending)
func InsertTaskWithSteps(label string, steps []string) {
	if len(steps) > 0 {
		// Populate steps
		structSteps := []Step{}

		for _, step := range steps {
			structSteps = append(structSteps, Step{Description: step})
		}

		// Insert task and steps
		db.Create(&Task{
			Label: label,
			Steps: structSteps,
		})
	} else {
		InsertTask(label)
	}
}
*/
