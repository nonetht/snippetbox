package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet 代表单个 Snippet 的数据信息
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel Define a SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	// Write the sql statement we want to execute. I've split it over 2 lines
	// for readability (which is surrounded with backquotes instead of normal double ...)
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use Exec() method on .. to execute the statement
	// DB.Exec() 返回 sql.Result 类型，提供了两种方法：
	// 1. LastInsertId() 返回数据库响应命令生成的整数（一个int64）
	// 2. RowsAffected() 返回受语句影响的行数
	result, err := m.DB.Exec(stmt, title, content, expires) // Exec() 语句幕后发生了什么？1. 传入预处理语句 2. 传入参数 3.关闭/释放预处理语句
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result to get the ID of your newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, convert it to an int
	return int(id), nil
}

// Get return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Write the SQL statement we want to execute. Again, I've split it over 2 lines for readability
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use the QueryRow() method on the connection pool to execute our SQL statement, passing in ...
	// Returns a pointer to a sql.Row object which holds the result from the database
	row := m.DB.QueryRow(stmt, id)

	// Initialize a new zeroed Snippet struct
	var s Snippet

	// Copy the values from each field in sql.Row to the corresponding field in the Snippet struct.
	// row.Scan 的参数都是**指针**类型，指向你想要填充的参数的位置。
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a sql.ErrNoRows error.
		// We use the errors.Is() function check for that error specifically, and return our
		// own ErrNoRecord error
		if errors.Is(err, sql.ErrNoRows) { // 要使用 error.Is() 函数更为安全，不要使用 "=="
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	// If everything went OK, return the filled Snippet struct
	return s, nil
}

// Latest return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// Write the SQL statement we want to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Use Query() method on the connection pool to execute our SQL statement. This return ...
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is ...
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs
	var snippets []Snippet

	// User rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets
		snippets = append(snippets, s)
	}

	// When the rows.Next() loop has finished we call rows.Err() to ree..
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
