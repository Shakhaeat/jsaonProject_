-- name: find-users-by-degree-year-board-roll
SELECT * FROM students WHERE roll_id = ? And degree = ? AND board = ? AND year = ?;
