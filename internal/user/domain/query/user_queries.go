package query

type GetUserByIDQuery struct {
    ID string
}

type GetUserByEmailQuery struct {
    Email string
}

type ListUsersQuery struct {
    // Можно добавить параметры для фильтрации и пагинации
    Limit  int
    Offset int
} 