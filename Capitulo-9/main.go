// Video: https://www.youtube.com/watch?v=fNhIkLnEb44&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=9&ab_channel=RobertoMorais
// Minuto 7:12
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS usuarios (
			id_usuarios int not null auto_increment,
    		nombre varchar(50) default null,
    		apellido varchar(50) default null,
    		correo text not null,
			primary key (id_usuarios)
		);
	`

	_, err := db.Exec(query)

	return err
}

func insertUser(db *sql.DB, nombre, apellido, correo string) (int64, error) {
	query := `
		INSERT INTO Usuarios(nombre, apellido, correo)
		VALUES (?, ?, ?)
	`
	result, err := db.Exec(query, nombre, apellido, correo)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getUsersEmails(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT correo FROM Usuarios
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usersEmails []string

	for rows.Next() {
		var correo string
		if err := rows.Scan(&correo); err != nil {
			return nil, err
		}

		usersEmails = append(usersEmails, correo)
	}

	return usersEmails, nil
}

func updateUsuarioApellido(db *sql.DB, id int64, newApellido string) error {
	_, err := db.Exec(`
		UPDATE Usuarios SET apellido = ? WHERE id_usuarios = ?
	`, newApellido, id)

	if err != nil {
		return err
	}

	return nil
}

func deleteUser(db *sql.DB, id int64) error {
	_, err := db.Exec(`
		DELETE FROM Usuarios WHERE id_usuarios = ?
	`, id)

	if err != nil {
		return err
	}

	return nil
}

// func main() {
// 	// DNS (Data source Name) para conectarse a la base de datos MySQL

// 	dns := "root:3482@tcp(127.0.0.1:3306)/barberflow"

// 	db, err := sql.Open("mysql", dns)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	delete := deleteUser(db, 3)

// 	if delete != nil {
// 		fmt.Println("Error al querer borrar el Usuario de ID:", 2)
// 	}

// 	fmt.Println("Se borró éxitosamente el usuario solicitado..")

// newSurname := updateUsuarioApellido(db, 1, "Gonzalez")

// if newSurname != nil {
// 	fmt.Println("Hubo un error al actualizar el apellido.. ERROR:", newSurname)
// 	return
// }

// fmt.Println("La actualización del apellido ha sido éxitosa..")

// query, err := getUsersEmails(db)

// if err != nil {
// 	fmt.Println("Error al obtener los correos de los usuarios")
// 	return
// }

// fmt.Println("Los correos son:",query)

/*
	id, err := insertUser(db, "Melania", "Nardelli", "melanialis2000@gmail.com")

	if err != nil {
		fmt.Println("No se puedieron insertar los datos en la tabla:", err)
		return
	}

	fmt.Println("Los datos fueron insertados correctamente.. :) en el usuario ID:", id)
*/
// Crear una tabla en la base de datos.
// err = createTable(db)
// if err != nil {
// 	fmt.Println("Error al crear la tabla:", err)
// 	return
// }

// fmt.Println("La tabla ha sido creada con éxito.. =)")

// Verificamos la conexión
// err = db.Ping()
// if err != nil {
// 	panic(err)
// }

// fmt.Println("Conexión a la base de datos éxitosa.")
// }

// GORM

type Cliente struct {
	ID       uint
	Nombre   string
	Apellido string
	Email    string
}

func main() {
	dns := "root:3482@tcp(127.0.0.1:3306)/barberflow"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println("Error al establecer al conexión con la base de datos:", err)
		return
	}

	db.AutoMigrate(&Cliente{})
	fmt.Println("Migración para Clientes completada exitosamente..")

	// Crear un registro
	// db.Create(&Cliente{Nombre: "Yair", Apellido: "Escobar", Email: "ivanescobargg@gmail.com"})
	// Leer un registro
	var clientes string
	db.First(&clientes, 1) // Busca por ID
	fmt.Println(clientes)
	// Actualizar un registro
	// db.Model(&usuario).Update("Nombre", "Iván")
	// Eliminar un registro
	// db.Delete(&usuario, 1)

	fmt.Println("Conexión con base de datos mediante GORM éxitosa :)..")
}
