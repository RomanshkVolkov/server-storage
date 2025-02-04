package service

import (
	"github.com/RomanshkVolkov/server-storage/internal/adapters/repository"
	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	schema "github.com/RomanshkVolkov/server-storage/internal/core/domain/schemas"
)

func (server Server) GetAllUsers() domain.APIResponse[[]domain.UserTableCRUD] {
	repo := repository.GetDBConnection(server.Host)
	users, err := repo.GetAllUsers()
	if err != nil {
		return domain.APIResponse[[]domain.UserTableCRUD]{
			Success: false,
			Message: domain.Message{
				En: "Error on get users",
				Es: "Error al obtener usuarios",
			},
			Error: err,
		}
	}

	return domain.APIResponse[[]domain.UserTableCRUD]{
		Success: true,
		Message: domain.Message{
			En: "Users list",
			Es: "Lista de usuarios",
		},
		Data: users,
	}
}

func (server Server) GetUserByID(id uint) domain.APIResponse[domain.EditableUser] {
	repo := repository.GetDBConnection(server.Host)
	user, err := repo.GetUserByID(id)
	if err != nil {
		return domain.APIResponse[domain.EditableUser]{
			Success: false,
			Message: domain.Message{
				En: "Error on get user",
				Es: "Error al obtener usuario",
			},
			Error: err,
		}
	}

	if user.ID == 0 {
		return repository.RecordNotFound[domain.EditableUser]()
	}

	return domain.APIResponse[domain.EditableUser]{
		Success: true,
		Message: domain.Message{
			En: "User data",
			Es: "Datos de usuario",
		},
		Data: user,
	}
}

func (server Server) CreateUser(request *domain.CreateUserRequest) domain.APIResponse[domain.User] {
	fields := schema.GenericForm[domain.CreateUserRequest]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return SchemaFieldsError[domain.User](failValidatedFields)
	}

	repo := repository.GetDBConnection(server.Host)
	user, err := repo.CreateUser(fields.Data)
	if err != nil {
		return domain.APIResponse[domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error on create user",
				Es: "Error al crear usuario",
			},
			Error: err,
		}
	}

	return domain.APIResponse[domain.User]{
		Success: true,
		Message: domain.Message{
			En: "User created",
			Es: "Usuario creado",
		},
		Data: user,
	}
}

func (server Server) UpdateUser(request *domain.EditableUser) domain.APIResponse[domain.User] {
	fields := schema.GenericForm[domain.EditableUser]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)

	if len(failValidatedFields) > 0 {
		return SchemaFieldsError[domain.User](failValidatedFields)
	}

	repo := repository.GetDBConnection(server.Host)
	user, err := repo.UpdateUser(fields.Data)
	if err != nil {
		return domain.APIResponse[domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error on update user",
				Es: "Error al actualizar usuario",
			},
			Error: err,
		}
	}

	return domain.APIResponse[domain.User]{
		Success: true,
		Message: domain.Message{
			En: "User updated",
			Es: "Usuario actualizado",
		},
		Data: user,
	}
}
