// package repository_test

// import (
// 	"log"
// 	"os"
// 	"testing"
// 	"todo-list/entities"
// 	"todo-list/repository"

// 	"github.com/stretchr/testify/require"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var testRepo repository.TaskRepository
// var db *gorm.DB

// func TestMain(m *testing.M) {
// 	dsn := "host=localhost user=postgres password=Parthu732 sslmode=disable port=5432 dbname=db_test"
// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn))
// 	if err != nil {
// 		log.Fatalf("failed to connect database %v", err)
// 	}
// 	if err := db.AutoMigrate(&entities.Task{}); err != nil {
// 		log.Fatalf("failed to migrate tables %v", err)
// 	}
// 	testRepo = repository.NewTaskRepository(db)
// 	code := m.Run()
// 	os.Exit(code)
// }

// func TestCreateRepo(t *testing.T) {
// 	task := entities.Task{
// 		Title:       "Learn Go",
// 		Description: "Write unit Test Cases",
// 		Completed:   false,
// 	}
// 	created, err := testRepo.CreateRepo(task)
// 	require.NoError(t, err)
// 	require.NotZero(t, created.ID, "ID should be set by the database")
// 	require.Equal(t, "Learn Go", created.Title)
// }

// func TestGetAllRepo(t *testing.T) {
// 	tasks, err := testRepo.GetAllRepo()
// 	require.NoError(t, err)
// 	require.GreaterOrEqual(t, len(tasks), 1, "should have atleast 1 task")
// }

// func TestGetByIdRepo(t *testing.T) {
// 	task := entities.Task{
// 		Title:       "Temp",
// 		Description: "for getbyid test",
// 	}
// 	created, _ := testRepo.CreateRepo(task)
// 	g, err := testRepo.GetByIdRepo(created.ID)
// 	require.NoError(t, err)
// 	require.Equal(t, created.ID, g.ID)
// 	require.Equal(t, "Temp", g.Title)
// }

// func TestUpdateRepo(t *testing.T) {
// 	task := entities.Task{
// 		Title:       "Old",
// 		Description: "Old Desc",
// 	}
// 	created, _ := testRepo.CreateRepo(task)
// 	created.Title = "New"
// 	updated, err := testRepo.UpdateRepo(created)
// 	require.NoError(t, err)
// 	require.Equal(t, "New", updated.Title)
// }

// func TestDeleteRepo(t *testing.T) {
// 	task := entities.Task{
// 		Title:       "Old",
// 		Description: "Old Desc",
// 	}
// 	created, _ := testRepo.CreateRepo(task)
// 	err := testRepo.DeleteRepo(created.ID)
// 	require.NoError(t, err)
// 	_, err = testRepo.GetByIdRepo(created.ID)
// 	require.Error(t, err, "should return an error because record is not found")
// }













package repository_test

import (
	"log"
	"os"
	"testing"
	"todo-list/entities"
	"todo-list/repository"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	testRepo repository.TaskRepository
	db    *gorm.DB
)

func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=Parthu732 sslmode=disable port=5432 dbname=db_test"
	var err error
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	if err := db.AutoMigrate(&entities.Task{}); err != nil {
		log.Fatalf("failed to migrate tables %v", err)
	}
	testRepo = repository.NewTaskRepository(db)
	code := m.Run()
	os.Exit(code)
}
func TestTaskRepository(t *testing.T) {
	type args struct {
		task entities.Task
	}
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "Create",
			run: func(t *testing.T) {
				task := entities.Task{
					Title:       "Learn Go",
					Description: "Write unit test cases",
					Completed:   false,
				}
				created, err := testRepo.CreateRepo(task)
				require.NoError(t, err)
				require.NotZero(t, created.ID, "ID should be set by database")
				require.Equal(t, task.Title, created.Title)
			},
		},
		{
			name: "GetAll",
			run: func(t *testing.T) {
				created, err := testRepo.GetAllRepo()
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(created), 1, "length should be greater than or equal to 1")
			},
		},
		{
			name: "GetById",
			run: func(t *testing.T) {
				task := entities.Task{
					Title:       "Learn Go",
					Description: "Write unit test cases",
				}
				created, _ := testRepo.CreateRepo(task)
				g, err := testRepo.GetByIdRepo(created.ID)
				require.NoError(t, err)
				require.Equal(t, g.ID, created.ID)
				require.Equal(t, g.Title, created.Title)
			},
		},
		{
			name: "update",
			run: func(t *testing.T) {
				task := entities.Task{
					Title:       "Learn Go",
					Description: "Write unit test cases",
				}
				created, _ := testRepo.CreateRepo(task)
				created.Title = "Learn Golang"
				g, err := testRepo.UpdateRepo(created)
				require.NoError(t, err)
				require.Equal(t, created.Title, g.Title)
			},
		},
		{
			name: "Delete",
			run: func(t *testing.T) {
				task := entities.Task{
					Title:       "Learn Go",
					Description: "Write unit test cases",
				}
				created, _ := testRepo.CreateRepo(task)
				err := testRepo.DeleteRepo(created.ID)
				require.NoError(t, err)
				_, err = testRepo.GetByIdRepo(created.ID)
				require.Error(t, err, "there is a problem in deleting")
			},
		},
	}
	for _, a := range tests {
		t.Run(a.name, a.run)
	}
}
