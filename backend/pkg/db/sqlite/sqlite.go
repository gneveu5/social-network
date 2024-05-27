package socialnetwork

import (
    "database/sql"
    "fmt"
    "path/filepath"
    "os"
	//"log"
    "sort"
    //models "socialnetwork/pkg/models"
    _ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
)

var Db *sql.DB

func MigrateUp() {
    migration := &migrate.FileMigrationSource{
        Dir: "pkg/db/migrations/sqlite",
    }
    n, err := migrate.Exec(Db, "sqlite3", migration, migrate.Up)
    if err != nil {
        fmt.Println("Migration Up error:", err)
        return
    }
    fmt.Printf("Migration Up done. %d migrations have been done.\n", n)
}

func MigrateDown() {

    migration := &migrate.FileMigrationSource{
        Dir: "pkg/db/migrations/sqlite",
    }
    n, err := migrate.ExecMax(Db, "sqlite3", migration, migrate.Down, 2)
    if err != nil {
        fmt.Println("Migration Down error:", err)
        return
    }

    //delete last two files
    files, err := filepath.Glob("pkg/db/migrations/sqlite/*.sql")
	if err != nil {
		fmt.Println("File retrieving error:", err)
		return
	}
	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	for i, file := range files {
		if err := os.Remove(file); err != nil {
			fmt.Println("Error while trying to delete file:", err)
			return
		}
		fmt.Printf("Migration file successfully deleted %s\n", file)
		if i == 1 {
			break
		}
	}

    fmt.Printf("Migration Down done. %d migrations have been done.\n", n)
}
